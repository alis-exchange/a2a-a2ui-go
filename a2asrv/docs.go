// Package a2asrv integrates the A2UI A2A extension with the
// [github.com/a2aproject/a2a-go/v2/a2asrv] server runtime.
//
// Advertise A2UI on your agent card or extension list using [AgentExtension] (fixed defaults for
// supported catalog URIs and inline catalogs) or [NewAgentExtension] with [WithSupportedCatalogIds]
// and/or [WithAcceptsInlineCatalogs] when those values need to differ per deployment or agent.
//
// [NewInterceptor] returns a [github.com/a2aproject/a2a-go/v2/a2asrv.CallInterceptor] whose
// [github.com/a2aproject/a2a-go/v2/a2asrv.CallInterceptor.Before] hook inspects requested extension
// URIs on the call and activates [AgentExtension] when the A2UI extension URI is among them.
// The [github.com/a2aproject/a2a-go/v2/a2asrv.CallInterceptor.After] hook is currently a no-op but
// satisfies the interface for wiring and future response-side behavior.
//
// Extension URI, constants, and defaults live in extension.go ([AgentExtension], [NewAgentExtension]).
package a2asrv
