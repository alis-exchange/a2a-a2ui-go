// Package a2asrv integrates the A2UI A2A extension with the
// [github.com/a2aproject/a2a-go/v2/a2asrv] server runtime.
//
// Use [AgentExtension] in your agent card or extension list so clients discover A2UI support,
// including default supported catalog URIs and whether inline catalogs are accepted.
//
// [NewInterceptor] returns a [github.com/a2aproject/a2a-go/v2/a2asrv.CallInterceptor] whose
// [github.com/a2aproject/a2a-go/v2/a2asrv.CallInterceptor.Before] hook inspects requested extension
// URIs on the call and activates [AgentExtension] when the A2UI extension URI is among them.
// The [github.com/a2aproject/a2a-go/v2/a2asrv.CallInterceptor.After] hook is currently a no-op but
// satisfies the interface for wiring and future response-side behavior.
//
// Extension URI and catalog defaults are defined alongside [AgentExtension] in extension.go.
package a2asrv
