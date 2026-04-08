// Package kit contains small helpers for working with A2UI capability data attached to a
// [context.Context] or deserialized from an A2A agent extension params map.
//
// Use [WithA2UICapabilities] to store capability maps on a context (typically after negotiation or
// client handshake) and [CapabilitiesFromContext] to read them back—for example when deciding
// whether to expose A2UI-related tools.
//
// [GetCatalogs] extracts supported catalog ID strings and normalized inline catalog objects from
// the extension params shape described by the A2UI specification.
package kit
