package whiteboard

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"yourproject/features/whiteboard/pages"
	"yourproject/features/whiteboard/services"

	"github.com/go-chi/chi/v5"
	"github.com/gorilla/sessions"
)

type Handlers struct {
	fabricService *services.FabricService
	cursorService *services.CursorService
	sessionStore  sessions.Store
}

func NewHandlers(sessionStore sessions.Store) *Handlers {
	return &Handlers{sessionStore: sessionStore}
}

func (h *Handlers) getUserID(r *http.Request) string {
	if tabId := r.URL.Query().Get("tabId"); tabId != "" {
		return tabId
	}
	return "anon_" + fmt.Sprintf("%d", time.Now().UnixNano())
}

func (h *Handlers) WhiteboardPage(w http.ResponseWriter, r *http.Request) {
	if err := pages.WhiteboardPage().Render(r.Context(), w); err != nil {
		http.Error(w, http.StatusText(500), 500)
	}
}

func (h *Handlers) StreamCanvas(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userID := h.getUserID(r)
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	flusher, _ := w.(http.Flusher)

	state, _ := h.fabricService.GetCanvasState(ctx)
	if state == nil { state = &services.FabricCanvasState{Version: "5.3.0", Objects: []*services.FabricObject{}} }
	data, _ := json.Marshal(state)
	fmt.Fprintf(w, "event: init\ndata: %s\n\n", string(data))
	flusher.Flush()

	ch, _ := h.fabricService.SubscribeDeltas(ctx)
	for {
		select {
		case <-ctx.Done(): return
		case delta := <-ch:
			if delta == nil || delta.UserID == userID { continue }
			data, _ := json.Marshal(delta)
			fmt.Fprintf(w, "event: delta\ndata: %s\n\n", string(data))
			flusher.Flush()
		}
	}
}

func (h *Handlers) SaveDrawing(w http.ResponseWriter, r *http.Request) {
	userID := h.getUserID(r)
	var req struct {
		ID     string                 `json:"id"`
		Tool   string                 `json:"tool"`
		Color  string                 `json:"color"`
		Delta  *services.FabricObject `json:"delta"`
		Live   bool                   `json:"_live"`
		UserID string                 `json:"userId"`
	}
	json.NewDecoder(r.Body).Decode(&req)
	obj := req.Delta
	if obj == nil { http.Error(w, "no delta", 400); return }
	if obj.ID == "" { obj.ID = req.ID }
	if obj.Type == "" { obj.Type = req.Tool }

	delta := &services.FabricDelta{ Type: "object", ObjectID: obj.ID, FabricData: obj, UserID: userID, Timestamp: time.Now().UnixMilli() }
	// Always save to KV (not for live drawing)
	if !req.Live { h.fabricService.SaveObject(r.Context(), obj) }
	h.fabricService.BroadcastDelta(r.Context(), delta)
}

func (h *Handlers) ClearCanvas(w http.ResponseWriter, r *http.Request) {
	h.fabricService.ClearCanvas(r.Context())
}

func (h *Handlers) DeleteObject(w http.ResponseWriter, r *http.Request) {
	userID := h.getUserID(r)
	objectID := chi.URLParam(r, "objectId")
	h.fabricService.DeleteObject(r.Context(), objectID)
	delta := &services.FabricDelta{ Type: "delete", ObjectID: objectID, UserID: userID, Timestamp: time.Now().UnixMilli() }
	h.fabricService.BroadcastDelta(r.Context(), delta)
}

func (h *Handlers) StreamCursors(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userID := h.getUserID(r)
	cursor := h.cursorService.GetOrCreateCursor(userID)
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	flusher, _ := w.(http.Flusher)

	all := h.cursorService.GetAllCursors()
	other := make([]*services.Cursor, 0)
	for _, c := range all { if c.UserID != userID { other = append(other, c) } }
	data, _ := json.Marshal(map[string]interface{}{ "myCursor": cursor, "cursors": other })
	fmt.Fprintf(w, "event: init\ndata: %s\n\n", string(data))
	flusher.Flush()

	ch, _ := h.cursorService.Subscribe(ctx)
	for {
		select {
		case <-ctx.Done(): return
		case c := <-ch:
			data, _ := json.Marshal(c)
			fmt.Fprintf(w, "event: cursor\ndata: %s\n\n", string(data))
			flusher.Flush()
		}
	}
}

func (h *Handlers) UpdateCursor(w http.ResponseWriter, r *http.Request) {
	userID := h.getUserID(r)
	var req struct { X float64 `json:"x"; Y float64 `json:"y"` }
	json.NewDecoder(r.Body).Decode(&req)
	h.cursorService.UpdateCursor(userID, req.X, req.Y)
}
