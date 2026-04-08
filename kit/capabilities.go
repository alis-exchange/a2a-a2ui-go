package kit

import (
	"context"
)

// a2uiCapabilitiesCtxKey is the private context key type for A2UI capability maps.
// Using a struct type avoids collisions with other packages' context values.
type a2uiCapabilitiesCtxKey struct{}

// WithA2UICapabilities returns a derived [context.Context] that carries capabilities (typically the
// agent extension Params map for A2UI). Tools and middleware can use [CapabilitiesFromContext] to
// read this value and decide whether to expose A2UI-related tools.
func WithA2UICapabilities(ctx context.Context, capabilities map[string]any) context.Context {
	return context.WithValue(ctx, a2uiCapabilitiesCtxKey{}, capabilities)
}

// CapabilitiesFromContext returns the A2UI capabilities map previously stored with
// [WithA2UICapabilities], and reports whether it was present. If ok is false, the map should not
// be used.
func CapabilitiesFromContext(ctx context.Context) (map[string]any, bool) {
	capabilities, ok := ctx.Value(a2uiCapabilitiesCtxKey{}).(map[string]any)
	return capabilities, ok
}
