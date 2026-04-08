// Package a2ui is the root of module [go.alis.build/a2a/extension/a2ui], a Go library for the
// [A2UI] (Agent-to-UI) A2A extension. It helps agents emit validated UI message streams and helps
// A2A servers advertise A2UI support to clients.
//
// # Subpackages
//
//   - [go.alis.build/a2a/extension/a2ui/tools] — ADK function tools and JSON Schema for A2UI
//     server-to-client messages (v0.9).
//
//   - [go.alis.build/a2a/extension/a2ui/kit] — Attach v0.9 client capabilities from executor
//     message metadata to [context.Context], and parse catalog fields from capability maps.
//
//   - [go.alis.build/a2a/extension/a2ui/a2asrv] — [github.com/a2aproject/a2a-go/v2/a2asrv] integration:
//     agent extension metadata and optional call interceptors for extension activation.
//
// [A2UI]: https://a2ui.org/
package a2ui
