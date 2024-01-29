package roleplay

import "github.com/American-Made-Code/herd/pkg/external/gpt"

type Service interface {
	CreateResponseMessage(command CreateResponseMessageCommand) (*Response, error)
}

type service struct {
	g gpt.Gateway
}

func NewService(providers []gpt.Auth) Service {
	return &service{
		g: gpt.NewGateway(providers),
	}
}

func (s *service) CreateResponseMessage(command CreateResponseMessageCommand) (*Response, error) {
	var currentParticipant Participant
	var gptMessages []gpt.Message

	// Get the roleplay participant that needs to respond
	for _, participant := range command.AllParticipants {
		if participant.Index == command.ResponseParticipantIndex {
			currentParticipant = participant
		}
	}

	// Get the system prompt from the current participant
	systemMessage := currentParticipant.toSystemMessage()
	// Add the system prompt as the first message
	gptMessages = append(gptMessages, systemMessage)

	// Add the user messages to the gpt messages
	for _, message := range command.Messages {
		if message.Participant.Index == currentParticipant.Index {
			assitantMessage := message.toAssistantMessage()
			gptMessages = append(gptMessages, assitantMessage)
		} else {
			userMessage := message.toUserMessage()
			gptMessages = append(gptMessages, userMessage)
		}
	}

	gptClient := s.g[command.ResponseGpt.Provider]

	gptCommand := gpt.CreateChatCommand{
		Model:    command.ResponseGpt.Model,
		Messages: gptMessages,
	}

	res, err := gptClient.CreateChatResponse(gptCommand)

	if err != nil {
		return nil, err
	}

	response := Response{}.fromChatResponse(res).withParticipant(currentParticipant)

	return &response, nil
}
