// Package tools provides Google ADK (Agent Development Kit) tools for generating and validating
// A2UI server-to-client messages against the v0.9 JSON Schema.
//
// The primary entry points are [GenerateA2UIMessages], which returns a [google.golang.org/adk/tool.Tool]
// that validates tool arguments, and [NewA2UIToolset], which wraps that tool in a filtered toolset
// exposed only when A2UI capabilities are present on the agent context (see
// [go.alis.build/a2a/extension/a2ui/kit.CapabilitiesFromContext]).
//
// Schema validation uses [github.com/google/jsonschema-go/jsonschema]. Additional semantic checks
// (for example, requiring a component with id "root" for each created surface) are implemented in
// this package alongside the schema.
package tools
