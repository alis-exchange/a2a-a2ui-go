// Package genkit bridges Google GenAI ADK function responses to A2A data parts.
//
// [GetA2uiDataPart] inspects a [google.golang.org/genai.Part] for an A2UI tool response and,
// if found, wraps the messages in an [github.com/a2aproject/a2a-go/v2/a2a.Part] with the
// "application/json+a2ui" mime type for downstream A2A transport.
package genkit

import (
	"github.com/a2aproject/a2a-go/v2/a2a"
	"go.alis.build/adk/a2ui/v09/tools"
	"google.golang.org/genai"
)

// GetA2uiDataPart inspects a genai.Part to determine if it contains an A2UI function response.
// If it does, it extracts the A2UI messages and wraps them in an a2a.DataPart with the
// appropriate mimeType ("application/json+a2ui"). Returns the new part and true if successful.
func GetA2uiDataPart(part *genai.Part) (a2uiData *a2a.Part, ok bool) {
	if part != nil && part.FunctionResponse != nil && part.FunctionResponse.Name == tools.GenerateA2UIMessagesToolName {
		if messages, ok := part.FunctionResponse.Response["messages"]; ok {
			dataPart := a2a.NewDataPart(messages)
			dataPart.Metadata = map[string]any{
				"mimeType": "application/json+a2ui",
			}
			return dataPart, true
		}
	}
	return nil, false
}
