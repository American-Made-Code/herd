package gpt

import "github.com/American-Made-Code/herd/pkg/external/openAi"

type openAiProvider struct {
	r openAi.Client
}

func (p openAiProvider) New(apiKey string) GatewayApi {
	p = openAiProvider{
		r: openAi.NewClient(apiKey),
	}
	return p
}

func (p openAiProvider) CreateChatResponse(command CreateChatCommand) (*ChatResponse, error) {
	// Convert the command to the provider's command
	rawCommand := command.toOpenAiCommand()
	// Make the request to the provider
	rawRes, err := p.r.CreateChatCompletion(rawCommand)
	// Convert the response to the gateway's response
	res := ChatResponse{}.fromOpenAiSchema(rawRes)
	// Check for errors
	if err != nil {
		return nil, err
	}
	// Return the gateway response
	return &res, nil
}
