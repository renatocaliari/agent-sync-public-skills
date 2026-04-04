---
name: boilerplate-go
description: Scaffolds Go web applications with Datastar, Templ, NATS, TailwindCSS, and Fabric.js whiteboard. Use when starting a new Go web project, creating real-time features, building hypermedia-driven apps, or scaffolding Go backend with SSE/WebSocket support. Triggers for "new go project", "scaffold go app", "go web boilerplate", "datastar go", "templ project", "go realtime app", "hypermedia go", "go sse server", "fabric.js whiteboard", or when user wants Go backend with reactive frontend.
---

# Go Web Application Boilerplate (v1.0.0-RC.8 Ready)

Go boilerplate inspired by [Northstar](https://github.com/delaneyj/toolbelt) (Boilerplate for creating real-time Hypermedia applications with [Datastar](https://data-star.dev) and [NATS](https://nats.io)), featuring [Fabric.js](http://fabricjs.com) Whiteboard and Real-time Voice ([LiveKit](https://livekit.io) + [Gemini](https://ai.google.dev)).

## How to Use This Skill

This skill contains bundled resources (templates, examples, database helper) that you should READ and WRITE to the user's project.

**All paths below are relative to this skill's directory.**

### File Paths Reference

| Resource | Path | What to do |
|----------|------|------------|
| **Scaffold** | `assets/scaffold/` | Copy entire directory to new project |
| **Database helper** | `references/database/database.go` | Copy to `<project>/db/database.go` |
| **Examples** | `references/examples/` | Copy relevant examples |

### Scaffolding a New Project

**Step 1:** List and read all scaffold files:
```
assets/scaffold/
├── cmd/web/main.go
├── config/config.go
├── router/router.go
├── nats/nats.go
├── features/
│   ├── common/
│   │   ├── components/
│   │   └── layouts/base.templ
│   ├── whiteboard/          # Standalone Whiteboard (Fabric.js)
│   │   ├── routes.go
│   │   ├── handlers.go
│   │   ├── static.go        # go:embed handler
│   │   ├── static/          # Isolated JS/CSS
│   │   └── pages/
│   ├── voice-training/      # Real-time Voice (LiveKit + Gemini)
│   │   ├── routes.go
│   │   ├── handlers.go
│   │   ├── services/
│   │   └── pages/
│   └── todos/               # Standalone Todos (NATS KV)
│       ├── routes.go
│       ├── handlers.go
│       └── components/
├── go.mod
└── Taskfile.yml
```

**Step 2:** Write each file to the target project.

**Step 3:** Update module name in `go.mod` and replace all `northstar` import paths.

## Feature Module Pattern (Plug & Play)

Each feature is 100% self-contained in `features/<name>/`. For portability, static assets (JS/CSS) should be kept inside the feature folder and served via `go:embed`.

### Isolated Static Files (`static.go`)
```go
package whiteboard
import (
	"embed"
	"io/fs"
	"net/http"
)
//go:embed static/*
var staticEmbed embed.FS

func StaticFS() http.FileSystem {
	fsys, _ := fs.Sub(staticEmbed, "static")
	return http.FS(fsys)
}
```

### Routes with Asset Serving
```go
func SetupRoutes(router chi.Router, ...) {
    // Serve its own static files
    router.Handle("/whiteboard/static/*", http.StripPrefix("/whiteboard/static/", http.FileServer(StaticFS())))
    router.Get("/whiteboard", handlers.Page)
}
```

## Datastar v1 Patterns (RC.8+)

### Templ Attributes
```templ
<div data-init={ "@get('/api/data')" }>
<button data-on:click={ "@post('/api/action')" }>
<input data-bind:signal-name>
<div data-text="$signalName">
```

### Backend (Go)
```go
sse := datastar.NewSSE(w, r)

// Patch with specific mode (Append, Prepend, Outer, etc.)
sse.PatchElementTempl(Component(), datastar.WithModeAppend(), datastar.WithSelector("#target"))

// Standard JSON decoding (prefer over ReadSignals for non-signal payloads)
json.NewDecoder(r.Body).Decode(&req)
```

## Real-time Voice (Gemini + LiveKit)

Implementation of a multimodal voice bot using Google Gemini Live API and LiveKit.

### Key Techniques:
1.  **User Gesture for Audio:** Delay `setMicrophoneEnabled(true)` until an explicit user click (e.g., Mute/Unmute button) to comply with browser Autoplay policies.
2.  **Storage Polyfill:** Inject memory-fallback for `localStorage` to prevent LiveKit crashes in Incognito/Strict environments.
3.  **Turn Tracking:** Use stable IDs for SSE fragments to support Idiomorph morphing during transcript streaming.

## Whiteboard Feature (Fabric.js)

- **Isolation:** CSS and JS are emmbedded within the module.
- **NATS Sync:** Delta-based synchronization (send only changed properties).
- **KISS:** Use standard HTTP POST for mutates and SSE for the stream.

## Best Practices

1.  **Vertical Slices:** One folder per feature. No leaking of assets to global `static/` folders.
2.  **DRY (NATS KV):** Use NATS JetStream for shared state across instances.
3.  **KISS (SSE):** Prefer SSE over WebSockets for one-way server-to-client updates.
4.  **Resilience:** Always wrap browser storage and media calls in `try/catch`.
