package a2asrv

import (
	"context"
	"slices"
	"strings"

	sdka2asrv "github.com/a2aproject/a2a-go/v2/a2asrv"
)

// interceptor implements [sdka2asrv.CallInterceptor] for optional A2UI extension activation.
type interceptor struct{}

// NewInterceptor returns a [sdka2asrv.CallInterceptor] that activates [AgentExtension] on the
// current call when the client requested the A2UI extension URI. Register it with your a2a-go
// server configuration alongside other interceptors as needed.
func NewInterceptor() *interceptor {
	return &interceptor{}
}

// Before implements [sdka2asrv.CallInterceptor.Before]. It parses requested extension URIs from
// the call context (splitting comma-separated values), and if [extensionURI] is requested,
// calls [sdka2asrv.CallContext.Extensions.Activate] with [&AgentExtension].
func (i *interceptor) Before(ctx context.Context, callCtx *sdka2asrv.CallContext, req *sdka2asrv.Request) (context.Context, any, error) {
	var requestedURIs []string

	// Some transports may use a comma-separated string of URIs, so we need to split them and trim whitespace.
	for _, uri := range callCtx.Extensions().RequestedURIs() {
		for part := range strings.SplitSeq(uri, ",") {
			requestedURIs = append(requestedURIs, strings.TrimSpace(part))
		}
	}

	if slices.Contains(requestedURIs, extensionURI) {
		callCtx.Extensions().Activate(&AgentExtension)
	}

	return ctx, nil, nil
}

// After implements [sdka2asrv.CallInterceptor.After]. It is a no-op for now; response-side
// A2UI handling may be added later without changing constructor signatures.
func (i *interceptor) After(ctx context.Context, callCtx *sdka2asrv.CallContext, resp *sdka2asrv.Response) error {
	return nil
}
