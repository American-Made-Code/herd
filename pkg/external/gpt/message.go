package gpt

import (
	"github.com/American-Made-Code/herd/pkg/external/gptExample"
	"github.com/American-Made-Code/herd/pkg/external/openAi"
)

type Message struct {
	// The contents of the message.
	Content string `json:"content"`
	// The Role of the messages author, can be
	// "system", "user", "assistant", or "tool".
	Role string `json:"role"`
	// An optional Name for the participant.
	// Provides the model information to differentiate
	// between participants of the same role.
	Name *string `json:"name,omitempty"`
}

func (model Message) fromOpenAiSchema(schema *openAi.Message) Message {
	model = Message{
		Content: schema.Content,
		Role:    schema.Role,
		Name:    schema.Name,
	}

	return model
}

func (model Message) toOpenAiSchema() openAi.Message {
	schema := openAi.Message{
		Content: model.Content,
		Role:    model.Role,
		Name:    model.Name,
	}

	return schema
}

func (model Message) fromExampleSchema(schema *gptExample.Message) Message {
	model = Message{
		Content: schema.Content,
		Role:    schema.Role,
		Name:    schema.Name,
	}

	return model
}

func (model Message) toExampleSchema() gptExample.Message {
	schema := gptExample.Message{
		Content: model.Content,
		Role:    model.Role,
		Name:    model.Name,
	}

	return schema
}
