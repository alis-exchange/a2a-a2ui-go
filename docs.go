// Package a2ui is the root of module [go.alis.build/a2a/extension/a2ui], a Go library for the
// [A2UI] (Agent-to-UI) A2A extension. It helps A2A servers advertise A2UI support to clients
// and bridges ADK function responses to A2A data parts.
//
// Transport-agnostic ADK tools and JSON Schema validation live in
// [go.alis.build/adk/a2ui]; this module provides the A2A-specific wiring.
//
// # Subpackages
//
//   - [go.alis.build/a2a/extension/a2ui/a2asrv] — [github.com/a2aproject/a2a-go/v2/a2asrv] integration:
//     agent extension metadata and optional call interceptors for extension activation.
//
//   - [go.alis.build/a2a/extension/a2ui/kit] — Attach v0.9 client capabilities from A2A executor
//     message metadata to [context.Context], delegating to [go.alis.build/adk/a2ui/kit].
//
//   - [go.alis.build/a2a/extension/a2ui/genkit] — Bridges [google.golang.org/genai.Part] A2UI
//     function responses to [github.com/a2aproject/a2a-go/v2/a2a.Part] for A2A transport.
//
// [A2UI]: https://a2ui.org/
package a2ui
