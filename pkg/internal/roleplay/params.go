package roleplay

import "github.com/American-Made-Code/herd/pkg/external/gpt"

type RoleplayParticipant struct {
	Index        int
	Provider     gpt.Provider
	Model        string
	SystemPrompt string
	Name         string
}

func (p RoleplayParticipant) toSystemMessage() gpt.Message {
	return gpt.Message{
		Content: p.SystemPrompt,
		Role:    gpt.System,
		Name:    &p.Name,
	}
}

type CreateResponseMessageCommand struct {
	AllParticipants          []RoleplayParticipant
	ResponseParticipantIndex int
	Messages                 []Message
}
