package gpt

import (
	"fmt"

	"github.com/American-Made-Code/herd/pkg/external/gptExample"
	"github.com/American-Made-Code/herd/pkg/external/openAi"
)

type Provider int

const (
	OpenAi Provider = iota
	GptExample
)

type Gateway interface {
	CreateChatResponse(command CreateChatCommand) (*ChatResponse, error)
}

type gateway struct {
	p          Provider
	openAi     openAi.Client
	gptExample gptExample.Client
}

func NewClient(provider Provider) Gateway {
	switch provider {
	case OpenAi:
		return &gateway{
			p:      provider,
			openAi: openAi.NewClient(),
			// Assuming other fields are not needed for this provider
		}
	case GptExample:
		return &gateway{
			p:          provider,
			gptExample: gptExample.NewClient(),
			// Assuming other fields are not needed for this provider
		}
	default:
		// Handle unknown provider case
		// This could be an error return or a default provider
		panic("unknown provider")
	}
}

func (g *gateway) CreateChatResponse(command CreateChatCommand) (*ChatResponse, error) {
	switch g.p {
	case OpenAi:
		return g.createChatResponseOpenAi(command)
	case GptExample:
		return g.createChatResponseExample(command)
	default:
		return nil, fmt.Errorf("unsupported provider")
	}
}