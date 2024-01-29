package gpt

type Provider int

const (
	OpenAi Provider = iota
	GptExample
)

type Auth struct {
	Provider Provider
	ApiKey   string
}

type GatewayApi interface {
	CreateChatResponse(command CreateChatCommand) (*ChatResponse, error)
	New(apiKey string) GatewayApi
}

var providerMap = map[Provider]GatewayApi{
	OpenAi:     openAiProvider{},
	GptExample: exampleProvider{},
}

type Gateway map[Provider]GatewayApi

func NewGateway(providers []Auth) Gateway {
	gatewayMap := providerMap

	for _, p := range providers {
		gatewayMap[p.Provider] = providerMap[p.Provider].New(p.ApiKey)
	}

	return gatewayMap
}
