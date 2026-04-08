package kit

import (
	"context"

	sdka2asrv "github.com/a2aproject/a2a-go/v2/a2asrv"
	a2asrv "go.alis.build/a2a/extension/a2ui/a2asrv"
)

// a2uiCapabilitiesCtxKey is the private context key type for A2UI capability maps.
// Using a struct type avoids collisions with other packages' context values.
type a2uiCapabilitiesCtxKey struct{}

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

	return context.WithValue(ctx, a2uiCapabilitiesCtxKey{}, a2uiCapabilities)
}

// CapabilitiesFromContext returns the v0.9 capability params map previously stored by
// [WithA2UICapabilities], and whether that store happened. If ok is false, the map must not be used.
func CapabilitiesFromContext(ctx context.Context) (map[string]any, bool) {
	capabilities, ok := ctx.Value(a2uiCapabilitiesCtxKey{}).(map[string]any)
	return capabilities, ok
}
