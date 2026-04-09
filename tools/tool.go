package tools

import (
	"fmt"

	"github.com/a2aproject/a2a-go/v2/a2a"
	"github.com/google/jsonschema-go/jsonschema"
	"go.alis.build/a2a/extension/a2ui/kit"
	"google.golang.org/adk/agent"
	"google.golang.org/adk/tool"
	"google.golang.org/adk/tool/functiontool"
	"google.golang.org/genai"
)

// ToolName is the ADK tool identifier for [GenerateA2UIMessages]. Agents and prompts should refer
// to this constant when configuring allowed tools.
const (
	ToolName = "generate_a2ui_messages"
)

// toolDescription is the human-readable instruction block shown to the model. It summarizes valid
// message shapes, emphasizes the required "root" component id, and includes minimal JSON examples.
var toolDescription = `
Validates an array of A2UI messages against the A2UI schema. Use this tool when showing components, forms, or custom UI surfaces.

**CRITICAL RENDERING RULES:**
- For every surfaceId created, at least one component in updateComponents.components MUST have "id": "root".
- The chat renderer mounts <ComponentNode id="root" />; without root the UI will not render.

**Example (minimal valid):**
{
  "messages": [
    {
      "version": "v0.9",
      "createSurface": {
        "surfaceId": "my-surface",
        "catalogId": "https://raw.githubusercontent.com/alis-exchange/a2ui-vuetify-renderer/main/catalog/vuetify-catalog.json"
      }
    },
    {
      "version": "v0.9",
      "updateComponents": {
        "surfaceId": "my-surface",
        "components": [
          { "component": "Card", "id": "root", "child": "text-1" },
          { "component": "Text", "id": "text-1", "text": "Hello" }
        ]
      }
    }
  ]
}

**Example (updateDataModel):**
{
  "messages": [
    {
      "version": "v0.9",
      "updateDataModel": {
        "surfaceId": "my-surface",
        "path": "/user/name",
        "value": "John Doe"
      }
    }
  ]
}

**Example (deleteSurface):**
{
  "messages": [
    {
      "version": "v0.9",
      "deleteSurface": {
        "surfaceId": "my-surface"
      }
    }
  ]
}

You MUST use the exact keys "createSurface", "updateComponents", "updateDataModel", or "deleteSurface".
`

// GenerateA2UIToolArgs is the JSON input/output shape for the generate A2UI messages tool.
// Messages is a heterogeneous list of A2UI server-to-client message objects (each map typically
// contains exactly one of createSurface, updateComponents, updateDataModel, or deleteSurface).
type GenerateA2UIToolArgs struct {
	Messages []map[string]any `json:"messages"`
}

// JSONSchema returns the tool argument JSON Schema for [GenerateA2UIToolArgs] so functiontool can
// expose it to the model: an object with required property "messages" whose value matches the
// inlined A2UI v0.9 server-to-client list schema (see schema.go).
func (GenerateA2UIToolArgs) JSONSchema() *jsonschema.Schema {
	return &jsonschema.Schema{
		Type: "object",
		Properties: map[string]*jsonschema.Schema{
			"messages": &a2UiServerToClientListSchema,
		},
		Required: []string{"messages"},
	}
}

// GenerateA2UIMessages builds an ADK [tool.Tool] that accepts [GenerateA2UIToolArgs], validates
// args.Messages against the resolved JSON Schema, then applies extra semantic checks (root
// component per surface). Validation errors are returned as tool errors so the model can self-correct.
func GenerateA2UIMessages() (tool.Tool, error) {
	handler := func(ctx tool.Context, args *GenerateA2UIToolArgs) (*GenerateA2UIToolArgs, error) {
		rs, err := a2UiServerToClientListSchema.Resolve(nil)
		if err != nil {
			return nil, fmt.Errorf("failed to resolve A2UI schema: %v", err)
		}

		if err := rs.Validate(args.Messages); err != nil {
			return nil, fmt.Errorf("validation failed. Please correct the following A2UI schema error and try again:\n- %s", err.Error())
		}

		if err := validateA2UISemantics(args.Messages); err != nil {
			return nil, fmt.Errorf("validation failed. %v", err)
		}

		return args, nil
	}

	return functiontool.New(functiontool.Config{
		Name:         ToolName,
		Description:  toolDescription,
		InputSchema:  GenerateA2UIToolArgs{}.JSONSchema(),
		OutputSchema: GenerateA2UIToolArgs{}.JSONSchema(),
	}, handler)
}

// NewA2UIToolset returns a named [tool.Toolset] containing [GenerateA2UIMessages], filtered so the
// tool is only visible when [kit.CapabilitiesFromContext] finds A2UI capabilities on the agent
// context (for example after extension negotiation).
func NewA2UIToolset() (tool.Toolset, error) {
	a2uiTool, err := GenerateA2UIMessages()
	if err != nil {
		return nil, err
	}

	var toolSet tool.Toolset = &a2uiToolset{
		name:  "a2ui",
		tools: []tool.Tool{a2uiTool},
	}
	toolSet = tool.FilterToolset(toolSet, func(ctx agent.ReadonlyContext, tool tool.Tool) bool {
		if _, ok := kit.CapabilitiesFromContext(ctx); !ok {
			return false
		}

		return true
	})

	return toolSet, nil
}

// a2uiToolset is a small adapter that implements [tool.Toolset] for a fixed slice of tools.
type a2uiToolset struct {
	name  string
	tools []tool.Tool
}

// Name returns the toolset name ("a2ui").
func (t *a2uiToolset) Name() string {
	return t.name
}

// Tools returns the tools registered in this toolset.
func (t *a2uiToolset) Tools(_ agent.ReadonlyContext) ([]tool.Tool, error) {
	return t.tools, nil
}

// GetA2uiDataPart inspects a genai.Part to determine if it contains an A2UI function response.
// If it does, it extracts the A2UI messages and wraps them in an a2a.DataPart with the
// appropriate mimeType ("application/json+a2ui"). Returns the new part and true if successful.
func GetA2uiDataPart(part *genai.Part) (a2uiData *a2a.Part, ok bool) {
	// Check if the part is a function response from the A2UI tool
	if part != nil && part.FunctionResponse != nil && part.FunctionResponse.Name == ToolName {
		// Extract the "messages" array from the response
		if messages, ok := part.FunctionResponse.Response["messages"]; ok {
			// Wrap the messages in an A2A data part
			dataPart := a2a.NewDataPart(messages)
			// Set the metadata to indicate it's an A2UI payload
			dataPart.Metadata = map[string]any{
				"mimeType": "application/json+a2ui",
			}
			return dataPart, true
		}
	}

	return nil, false
}
