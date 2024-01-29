package roleplay

import "github.com/American-Made-Code/herd/pkg/external/gpt"

type Service interface {
	CreateResponseMessage(command CreateResponseMessageCommand) (*Message, error)
}

var gatewayMap = map[gpt.Provider]gpt.GatewayApi{
	gpt.OpenAi:     gpt.NewClient(gpt.OpenAi),
	gpt.GptExample: gpt.NewClient(gpt.GptExample),
}

type service struct {
	g gpt.GatewayApi
}
