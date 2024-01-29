package roleplay

import "github.com/American-Made-Code/herd/pkg/external/gpt"

type Message struct {
	Content     string
	Participant RoleplayParticipant
}

func (m Message) toAssistantMessage() gpt.Message {
	return gpt.Message{
		Content: m.Content,
		Role:    gpt.Assistant,
		Name:    &m.Participant.Name,
	}
}

func (m Message) toUserMessage() gpt.Message {
	return gpt.Message{
		Content: m.Content,
		Role:    gpt.User,
		Name:    &m.Participant.Name,
	}
}
