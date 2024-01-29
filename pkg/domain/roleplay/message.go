package roleplay

import (
	"fmt"

	"github.com/American-Made-Code/herd/pkg/external/gpt"
)

type Message struct {
	Content     string
	Participant Participant
}

func (m Message) toAssistantMessage() gpt.Message {
	return gpt.Message{
		Content: m.Content,
		Role:    gpt.Assistant,
	}
}

func (m Message) toUserMessage() gpt.Message {
	return gpt.Message{
		Content: m.Content,
		Role:    gpt.User,
		Name:    &m.Participant.Name,
	}
}

func (m Message) fromGptMessage(source gpt.Message) Message {
	m.Content = source.Content

	fmt.Printf("\n m: %+v\n", m)

	return m
}

func (m Message) withParticipant(source Participant) Message {
	m.Participant = source

	return m
}
