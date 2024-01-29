package gpt

import "github.com/American-Made-Code/herd/pkg/external/gptExample"

type exampleProvider struct {
	r gptExample.Client
}

func (p exampleProvider) New(apiKey string) GatewayApi {
	p = exampleProvider{
		r: gptExample.NewClient(apiKey),
	}
	return p
}

func (p exampleProvider) CreateChatResponse(command CreateChatCommand) (*ChatResponse, error) {
	// Convert the command to the provider's command
	rawCommand := command.toExampleCommand()
	// Make the request to the provider
	rawRes, err := p.r.CreateChatCompletion(rawCommand)
	// Convert the response to the gateway's response
	res := ChatResponse{}.fromExampleSchema(rawRes)
	// Check for errors
	if err != nil {
		return nil, err
	}
	// Return the gateway response
	return &res, nil
}
