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
