package a2asrv

import (
	"slices"

	"github.com/a2aproject/a2a-go/v2/a2a"
	sdka2asrv "github.com/a2aproject/a2a-go/v2/a2asrv"
)

// Well-known URIs for the A2UI v0.9 A2A extension and reference catalogs used in [AgentExtension].
const (
	// extensionURI is the A2A extension identifier for A2UI v0.9. Clients request this URI to
	// negotiate A2UI; the server should advertise the same URI in agent metadata.
	extensionURI = "https://a2ui.org/a2a-extension/a2ui/v0.9"

	// vuetifyCatalogURI points to the public Vuetify catalog used as a default supported catalog.
	vuetifyCatalogURI = "https://raw.githubusercontent.com/alis-exchange/a2ui-vuetify-renderer/main/catalog/vuetify-catalog.json"

	// basicCatalogURI points to the A2UI basic catalog from the specification.
	basicCatalogURI = "https://a2ui.org/specification/v0_9/basic_catalog.json"
)

// AgentExtension describes the A2UI extension for [a2a.AgentExtension] lists and agent cards.
// It is not marked Required so clients may omit A2UI; when present, Params advertise which
// catalog IDs are pre-supported and whether inline catalog payloads are allowed.
var AgentExtension = a2a.AgentExtension{
	URI:         extensionURI,
	Description: "Enables the agent to generate rich, interactive user interfaces that render natively across web, mobile, and desktop, without executing arbitrary code.",
	Required:    false,
	Params: map[string]any{
		"supportedCatalogIds": []string{
			vuetifyCatalogURI,
			basicCatalogURI,
		},
		"acceptsInlineCatalogs": true,
	},
}

// IsActivated reports whether the A2UI v0.9 extension ([extensionURI]) is currently activated on
// callCtx—i.e. present in [sdka2asrv.CallContext.Extensions.ActivatedURIs]. Typical use is after an
// interceptor or handler has activated the extension in response to a client request.
func IsActivated(callCtx *sdka2asrv.CallContext) bool {
	return slices.Contains(callCtx.Extensions().ActivatedURIs(), extensionURI)
}

// NewAgentExtensionOptions holds values merged into the A2UI extension [a2a.AgentExtension.Params]
// by [NewAgentExtension]. Callers normally configure it only via [NewAgentExtensionOption] functions.
type NewAgentExtensionOptions struct {
	// SupportedCatalogIds lists catalog URIs this agent advertises as pre-supported (A2UI extension
	// param "supportedCatalogIds"). Clients may rely on these IDs when resolving component schemas.
	SupportedCatalogIds []string
	// AcceptsInlineCatalogs is the A2UI extension param "acceptsInlineCatalogs": when true, the
	// agent may embed catalog definitions inline in messages instead of referencing only the listed URIs.
	AcceptsInlineCatalogs bool
}

// NewAgentExtensionOption configures [NewAgentExtensionOptions] before [NewAgentExtension] builds
// the returned [a2a.AgentExtension].
type NewAgentExtensionOption func(*NewAgentExtensionOptions)

// WithSupportedCatalogIds sets the advertised supported catalog URIs. The slice is stored by
// reference; do not mutate it after passing it to [NewAgentExtension].
func WithSupportedCatalogIds(supportedCatalogIds []string) NewAgentExtensionOption {
	return func(options *NewAgentExtensionOptions) {
		options.SupportedCatalogIds = supportedCatalogIds
	}
}

// WithAcceptsInlineCatalogs sets whether the agent accepts inline catalog payloads (extension param
// "acceptsInlineCatalogs").
func WithAcceptsInlineCatalogs(acceptsInlineCatalogs bool) NewAgentExtensionOption {
	return func(options *NewAgentExtensionOptions) {
		options.AcceptsInlineCatalogs = acceptsInlineCatalogs
	}
}

// NewAgentExtension returns an [a2a.AgentExtension] for A2UI v0.9 ([extensionURI]) with the same
// URI, description, and Required flag as [AgentExtension], but with Params derived from options.
//
// Defaults match [AgentExtension]: [vuetifyCatalogURI] and [basicCatalogURI] in SupportedCatalogIds,
// and AcceptsInlineCatalogs true. Pass [WithSupportedCatalogIds], [WithAcceptsInlineCatalogs], or both
// to override defaults for agent cards or dynamic registration.
func NewAgentExtension(opts ...NewAgentExtensionOption) a2a.AgentExtension {
	options := NewAgentExtensionOptions{
		SupportedCatalogIds: []string{
			vuetifyCatalogURI,
			basicCatalogURI,
		},
		AcceptsInlineCatalogs: true,
	}
	for _, opt := range opts {
		opt(&options)
	}

	return a2a.AgentExtension{
		URI:         extensionURI,
		Description: "Enables the agent to generate rich, interactive user interfaces that render natively across web, mobile, and desktop, without executing arbitrary code.",
		Required:    false,
		Params: map[string]any{
			"supportedCatalogIds":   options.SupportedCatalogIds,
			"acceptsInlineCatalogs": options.AcceptsInlineCatalogs,
		},
	}
}
