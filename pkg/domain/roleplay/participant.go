package roleplay

import "github.com/American-Made-Code/herd/pkg/external/gpt"

type Participant struct {
	Index        int
	SystemPrompt string
	Name         string
}

func (p Participant) toSystemMessage() gpt.Message {
	return gpt.Message{
		Content: p.SystemPrompt,
		Role:    gpt.System,
		Name:    &p.Name,
	}
}
