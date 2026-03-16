package whiteboard

import (
	"github.com/go-chi/chi/v5"
	"github.com/gorilla/sessions"
)

func SetupRoutes(router chi.Router, sessionStore sessions.Store) error {
	handlers := NewHandlers(sessionStore)

	router.Get("/whiteboard", handlers.WhiteboardPage)
	router.Get("/whiteboard/stream", handlers.StreamCanvas)
	router.Post("/whiteboard/drawing", handlers.SaveDrawing)
	router.Delete("/whiteboard/drawings", handlers.ClearCanvas)
	router.Delete("/whiteboard/object/{objectId}", handlers.DeleteObject)
	router.Get("/whiteboard/cursors", handlers.StreamCursors)
	router.Post("/whiteboard/cursor", handlers.UpdateCursor)

	return nil
}
