package a2asrv

import "github.com/a2aproject/a2a-go/v2/a2a"

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
