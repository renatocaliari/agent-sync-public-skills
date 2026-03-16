---
name: boilerplate-go
description: Scaffolds Go web applications with Datastar, Templ, NATS, TailwindCSS, and Fabric.js whiteboard. Use when starting a new Go web project, creating real-time features, building hypermedia-driven apps, or scaffolding Go backend with SSE/WebSocket support. Triggers for "new go project", "scaffold go app", "go web boilerplate", "datastar go", "templ project", "go realtime app", "hypermedia go", "go sse server", "fabric.js whiteboard", or when user wants Go backend with reactive frontend.
---

# Go Web Application Boilerplate

A production-ready boilerplate for building real-time web applications in Go with hypermedia-driven architecture.

## How to Use This Skill

This skill contains bundled resources (templates, examples, database helper) that you should READ and WRITE to the user's project.

**All paths below are relative to this skill's directory.** Use the Read tool with the full path.

### File Paths Reference

| Resource | Path | What to do |
|----------|------|------------|
| **Scaffold** | `assets/scaffold/` | Copy entire directory to new project |
| **Database helper** | `references/database/database.go` | Copy to `<project>/db/database.go` |
| **Database example** | `references/database/user_crud_example.templ` | Reference for CRUD patterns |
| **Examples** | `references/examples/*.templ` | Copy relevant examples |

### Scaffolding a New Project

**Step 1:** List and read all scaffold files:
```
assets/scaffold/
├── cmd/web/main.go
├── config/config.go
├── config/config_dev.go
├── config/config_prod.go
├── router/router.go
├── nats/nats.go
├── features/common/layouts/base.templ
├── features/common/components/navigation.templ
├── features/common/components/shared.templ
├── features/whiteboard/
│   ├── routes.go
│   ├── handlers.go
│   ├── pages/whiteboard.templ
│   └── services/
│       ├── fabric_service.go
│       └── cursor_service.go
├── go.mod
├── Dockerfile
├── Taskfile.yml
└── .gitignore
```

**Step 2:** Write each file to the target project

**Step 3:** Update module name in `go.mod`

**Step 4:** Replace `northstar` in all import paths with new module name

### Adding Database Support

1. Read `references/database/database.go` → Write to `<project>/db/database.go`
2. Read `references/database/README.md` for documentation
3. Reference `references/database/user_crud_example.templ` for CRUD patterns

### Adding Examples

Read files from `references/examples/` as needed:

| File | Pattern |
|------|---------|
| `click_to_edit.templ` | Inline editing |
| `active_search.templ` | Debounced search |
| `counter.templ` | Real-time counter |
| `infinite_scroll.templ` | Infinite scroll |
| `lazy_load.templ` | Lazy loading |
| `file_upload.templ` | File uploads |
| `todo_mvc.templ` | Session-based CRUD |
| `whiteboard/` | Fabric.js real-time whiteboard |

## Stack

- **Go** - Backend language
- **Datastar** - Reactive frontend framework (hypermedia over the wire)
- **Templ** - Type-safe HTML templating
- **Chi** - HTTP router
- **NATS JetStream** - Real-time messaging and state persistence
- **TailwindCSS + DaisyUI** - Styling

## Project Structure

```
.
├── cmd/web/main.go          # Application entrypoint
├── config/                   # Environment-based configuration
├── router/router.go         # Route registration
├── nats/nats.go             # Embedded NATS setup
├── features/                 # Feature modules (vertical slices)
│   ├── common/
│   │   ├── layouts/base.templ
│   │   └── components/
│   └── <feature>/
│       ├── routes.go        # Route definitions
│       ├── handlers.go      # HTTP handlers
│       ├── services/        # Business logic
│       └── pages/           # Templ templates
├── web/resources/static/     # Compiled assets
├── Taskfile.yml             # Build tasks
├── Dockerfile
└── go.mod
```

## Feature Module Pattern

Each feature is self-contained in `features/<name>/`:

### routes.go
```go
package counter

import (
    "github.com/go-chi/chi/v5"
    "github.com/gorilla/sessions"
)

func SetupRoutes(router chi.Router, sessionStore sessions.Store) error {
    handlers := NewHandlers(sessionStore)
    
    router.Get("/counter", handlers.CounterPage)
    router.Get("/counter/data", handlers.CounterData)
    router.Post("/counter/increment", handlers.Increment)
    
    return nil
}
```

### handlers.go
```go
package counter

import (
    "net/http"
    "yourproject/features/counter/pages"
    "github.com/starfederation/datastar-go/datastar"
    "github.com/Jeffail/gabs/v2"
)

type Handlers struct {
    sessionStore sessions.Store
}

func NewHandlers(sessionStore sessions.Store) *Handlers {
    return &Handlers{sessionStore: sessionStore}
}

func (h *Handlers) CounterPage(w http.ResponseWriter, r *http.Request) {
    if err := pages.CounterPage().Render(r.Context(), w); err != nil {
        http.Error(w, http.StatusText(500), 500)
    }
}

func (h *Handlers) CounterData(w http.ResponseWriter, r *http.Request) {
    sse := datastar.NewSSE(w, r)
    signals := pages.CounterSignals{Count: 0}
    if err := sse.PatchElementTempl(pages.Counter(signals)); err != nil {
        http.Error(w, http.StatusText(500), 500)
    }
}

func (h *Handlers) Increment(w http.ResponseWriter, r *http.Request) {
    update := gabs.New()
    update.Set(1, "count")
    
    sse := datastar.NewSSE(w, r)
    if err := sse.MarshalAndPatchSignals(update); err != nil {
        http.Error(w, http.StatusText(500), 500)
    }
}
```

### pages/counter.templ
```templ
package pages

import (
    "github.com/starfederation/datastar-go/datastar"
    "yourproject/features/common/layouts"
)

type CounterSignals struct {
    Count int `json:"count"`
}

templ Counter(signals CounterSignals) {
    <div id="container" data-signals={ templ.JSONString(signals) }>
        <span data-text="$count"></span>
        <button
            class="btn btn-primary"
            data-on:click={ datastar.PostSSE("/counter/increment") }
        >
            Increment
        </button>
    </div>
}

templ CounterPage() {
    @layouts.Base("Counter") {
        <div id="container" data-init={ datastar.GetSSE("/counter/data") }></div>
    }
}
```

## Datastar Patterns

### Server-Sent Events (SSE)

```go
sse := datastar.NewSSE(w, r)

sse.PatchElementTempl(templComponent)

sse.MarshalAndPatchSignals(gabsContainer)

sse.ExecuteScript("window.location.reload()")

sse.ConsoleError(err)
```

### Signal Operations

```go
type Signals struct {
    Input string `json:"input"`
}

datastar.ReadSignals(r, &signals)

update := gabs.New()
update.Set(value, "signalName")
sse.MarshalAndPatchSignals(update)
```

### Templ Datastar Attributes

```templ
<div data-signals={ templ.JSONString(signals) }>

<div data-text="$signalName">

<input data-bind:signal-name>

<button data-on:click={ datastar.PostSSE("/path") }>

<div data-on:keydown__debounce.200ms={ datastar.GetSSE("/search") }>

<div data-init={ datastar.GetSSE("/data") }>

<div data-indicator="fetching" data-attr:disabled="$fetching">
```

### Datastar HTTP Helpers

```go
datastar.GetSSE("/path")
datastar.PostSSE("/path")
datastar.PutSSE("/path")
datastar.DeleteSSE("/path")
datastar.GetSSE("/path/%d", arg1, arg2)
```

## NATS Patterns

### Service with Key-Value Store

```go
type TodoService struct {
    kv jetstream.KeyValue
}

func NewTodoService(ns *embeddednats.Server) (*TodoService, error) {
    nc, _ := ns.Client()
    js, _ := jetstream.New(nc)
    
    kv, _ := js.CreateOrUpdateKeyValue(ctx, jetstream.KeyValueConfig{
        Bucket:      "todos",
        Description: "Todo storage",
        Compression: true,
        TTL:         time.Hour,
    })
    
    return &TodoService{kv: kv}, nil
}

func (s *TodoService) Save(ctx context.Context, key string, data []byte) error {
    _, err := s.kv.Put(ctx, key, data)
    return err
}

func (s *TodoService) Watch(ctx context.Context, key string) (jetstream.KeyWatcher, error) {
    return s.kv.Watch(ctx, key)
}
```

### Real-time Updates

```go
func (h *Handlers) StreamTodos(w http.ResponseWriter, r *http.Request) {
    sse := datastar.NewSSE(w, r)
    
    watcher, _ := h.service.Watch(r.Context(), sessionID)
    defer watcher.Stop()
    
    for {
        select {
        case <-r.Context().Done():
            return
        case entry := <-watcher.Updates():
            if entry == nil {
                continue
            }
            var data TodoMVC
            json.Unmarshal(entry.Value(), &data)
            sse.PatchElementTempl(TodosView(data))
        }
    }
}
```

## SQLite Database Patterns

See `references/database/` for the complete database helper.

### Setup

```go
import "yourproject/db"

database, _ := db.NewDatabase(ctx,
    db.DatabaseWithFilename("data/app.sqlite"),
    db.DatabaseWithMigrations([]string{
        `CREATE TABLE IF NOT EXISTS users (
            id TEXT PRIMARY KEY,
            name TEXT NOT NULL,
            email TEXT NOT NULL,
            created_at INTEGER
        )`,
    }),
)
defer database.Close()
```

### Read Transaction

```go
err := database.ReadTX(ctx, func(tx *sqlite.Conn) error {
    stmt := tx.Prep("SELECT id, name FROM users WHERE email = $email")
    defer stmt.Reset()
    
    stmt.SetText("$email", email)
    
    hasRow, err := stmt.Step()
    if err != nil || !hasRow {
        return err
    }
    
    id := stmt.GetText("id")
    name := stmt.GetText("name")
    // ...
    return nil
})
```

### Write Transaction

```go
err := database.WriteTX(ctx, func(tx *sqlite.Conn) error {
    stmt := tx.Prep(`INSERT INTO users (id, name, email, created_at)
                     VALUES ($id, $name, $email, $createdAt)`)
    defer stmt.Reset()
    
    stmt.SetText("$id", generateID())
    stmt.SetText("$name", name)
    stmt.SetText("$email", email)
    stmt.SetInt64("$createdAt", time.Now().Unix())
    
    _, err := stmt.Step()
    return err
})
```

## Configuration

### config/config.go
```go
package config

type Environment string

const (
    Dev  Environment = "dev"
    Prod Environment = "prod"
)

type Config struct {
    Environment   Environment
    Host          string
    Port          string
    LogLevel      slog.Level
    SessionSecret string
}

var Global *Config

func getEnv(key, fallback string) string {
    if val, ok := os.LookupEnv(key); ok {
        return val
    }
    return fallback
}
```

## Development Commands

```bash
go tool task live      # Start with hot reload
go tool task run       # Build and run
go tool task build     # Production build
go tool task debug     # Run with debugger
```

## Common Imports

```go
import (
    "github.com/go-chi/chi/v5"
    "github.com/gorilla/sessions"
    "github.com/starfederation/datastar-go/datastar"
    "github.com/a-h/templ"
    "github.com/Jeffail/gabs/v2"
    "github.com/nats-io/nats.go/jetstream"
    "github.com/delaneyj/toolbelt/embeddednats"
)
```

## Bundled Resources

See **"How to Use This Skill"** section above for:
- Complete file paths for scaffold, database, and examples
- Step-by-step instructions for scaffolding a new project
- How to add database support
- Available examples with patterns

## Whiteboard Feature (Fabric.js)

The scaffold includes a real-time collaborative whiteboard using Fabric.js with NATS sync.

### Structure

```
features/whiteboard/
├── routes.go              # Route definitions
├── handlers.go            # HTTP handlers + SSE streams
├── pages/
│   └── whiteboard.templ   # Fabric.js frontend
└── services/
    ├── fabric_service.go  # NATS KV + Pub/Sub
    └── cursor_service.go  # Multi-user cursors
```

### Routes

| Route | Method | Description |
|-------|--------|-------------|
| `/whiteboard` | GET | Whiteboard page |
| `/whiteboard/stream` | GET | SSE for object deltas |
| `/whiteboard/drawing` | POST | Save/update object |
| `/whiteboard/drawings` | DELETE | Clear canvas |
| `/whiteboard/object/{id}` | DELETE | Delete specific object |
| `/whiteboard/cursors` | GET | SSE for cursor positions |
| `/whiteboard/cursor` | POST | Update cursor position |

### Frontend Features

- **Tools**: Select, Pencil, Line, Rectangle, Circle, Text
- **Colors**: Black, Red, Blue, Green
- **Brush Size**: Slider 1-50
- **Image Upload**: Button or drag-and-drop
- **Clipboard Paste**: Ctrl+V for images/text
- **Delete**: Delete key or trash button
- **Live Sync**: Objects sync while moving/resizing (50ms debounce)
- **Multi-user Cursors**: Funny names + colored cursors

### NATS Patterns

```go
// Key-Value Storage
kv, _ := js.CreateOrUpdateKeyValue(ctx, jetstream.KeyValueConfig{
    Bucket:      "whiteboard",
    Compression: true,
    TTL:         time.Hour * 24 * 7,
    MaxBytes:    128 * 1024 * 1024,
})

// Broadcast deltas
delta := &FabricDelta{
    Type:       "object",
    ObjectID:   obj.ID,
    FabricData: obj,
    UserID:     userID,
    Timestamp:  time.Now().UnixMilli(),
}
s.nc.Publish("whiteboard.delta", deltaData)

// Subscribe to updates
ch, _ := s.SubscribeDeltas(ctx)
for delta := range ch {
    // Broadcast to SSE clients
}
```

### Delta Sync Optimization

```javascript
// Debounce utility
function debounce(fn, delay) {
    let timer = null;
    return function(...args) {
        clearTimeout(timer);
        timer = setTimeout(() => fn.apply(this, args), delay);
    };
}

// Send only changed properties
const sendDelta = debounce(function(delta) {
    fetch('/whiteboard/drawing', {
        method: 'POST',
        body: JSON.stringify({ delta })
    });
}, 150);

// Live sync while moving (50ms)
canvas.on('object:moving', (e) => {
    sendDelta({ id: obj.id, left: obj.left, top: obj.top });
});
```

### SSE Stream Handler

```go
func (h *Handlers) StreamCanvas(w http.ResponseWriter, r *http.Request) {
    flusher, _ := w.(http.Flusher)
    
    // Send initial state
    state, _ := h.fabricService.GetCanvasState(ctx)
    data, _ := json.Marshal(state)
    fmt.Fprintf(w, "event: init\ndata: %s\n\n", string(data))
    flusher.Flush()
    
    // Stream deltas
    ch, _ := h.fabricService.SubscribeDeltas(ctx)
    for delta := range ch {
        if delta.UserID != userID {
            data, _ := json.Marshal(delta)
            fmt.Fprintf(w, "event: delta\ndata: %s\n\n", string(data))
            flusher.Flush()
        }
    }
}
```

## Best Practices

1. **One feature per directory** - Keep routes, handlers, services, and templates together
2. **Templ generates Go code** - Run `go tool templ generate` after editing `.templ` files
3. **Use `data-init` for SSE streams** - Auto-connects on page load
4. **Signal names use kebab-case** - `data-bind:first-name` becomes `$firstName` in JS
5. **Debounce real-time events** - 50-150ms prevents flooding
6. **NATS KV for state** - Embedded NATS with JetStream for persistence
7. **Delta sync** - Send only changed properties, not full objects
8. **Session-based user identity** - Use gorilla/sessions for user tracking
5. **Use gabs for dynamic signal updates** - Build JSON patches flexibly
6. **NATS KV for state** - Embedded NATS with JetStream for persistence
7. **Session-based user identity** - Use gorilla/sessions for user tracking
8. **Hot reload in dev** - Air for backend, templ watch, tailwind watch