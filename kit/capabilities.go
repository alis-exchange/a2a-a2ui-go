// Package kit contains helpers for A2UI client capability data: attaching it to a [context.Context]
// and parsing catalog fields from a capability params map.
//
// [WithA2UICapabilities] copies A2UI v0.9 capabilities from an executor message
// (Metadata["a2uiClientCapabilities"]["v0.9"]) onto the context when the A2A call context is
// present and the A2UI extension is activated. [CapabilitiesFromContext] reads that map—for example
// to decide whether to expose A2UI tools.
//
// [GetCatalogs] extracts supportedCatalogIds and inlineCatalogs from a capability params map in
// the shape described by the A2UI specification.
package kit

import (
	"context"

	sdka2asrv "github.com/a2aproject/a2a-go/v2/a2asrv"
	a2asrv "go.alis.build/a2a/extension/a2ui/a2asrv"
	"go.alis.build/adk/a2ui/kit"
)

// WithA2UICapabilities attaches the A2UI v0.9 client capability map to ctx when every check
// succeeds; otherwise it returns ctx unchanged.
//
// It requires:
//   - ctx to carry an [sdka2asrv.CallContext] ([sdka2asrv.CallContextFrom]).
//   - the A2UI extension to be activated on that call ([a2asrv.IsActivated]).
//   - execCtx with a non-nil Message and Metadata.
//
// It then reads execCtx.Message.Metadata["a2uiClientCapabilities"] as a map, takes the "v0.9" entry
// as the capability params (e.g. supportedCatalogIds, acceptsInlineCatalogs), and stores that inner
// map on the context for [CapabilitiesFromContext].
func WithA2UICapabilities(ctx context.Context, execCtx *sdka2asrv.ExecutorContext) context.Context {
	callCtx, ok := sdka2asrv.CallContextFrom(ctx)
	if !ok {
		return ctx
	}

	if !a2asrv.IsActivated(callCtx) {
		return ctx
	}

	if execCtx == nil || execCtx.Message == nil || execCtx.Message.Metadata == nil {
		return ctx
	}

	capabilitiesMap, hasCapabilities := execCtx.Message.Metadata["a2uiClientCapabilities"].(map[string]any)
	if !hasCapabilities {
		return ctx
	}

	a2uiCapabilities, hasA2UICapabilities := capabilitiesMap["v0.9"].(map[string]any)
	if !hasA2UICapabilities {
		return ctx
	}

	return kit.WithA2UICapabilities(ctx, a2uiCapabilities)
}

// CapabilitiesFromContext returns the v0.9 capability params map previously stored by
// [WithA2UICapabilities], and whether that store happened. If ok is false, the map must not be used.
func CapabilitiesFromContext(ctx context.Context) (map[string]any, bool) {
	return kit.CapabilitiesFromContext(ctx)
}
