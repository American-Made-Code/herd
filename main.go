package main

import (
	"fmt"
	"os"

	"github.com/American-Made-Code/herd/pkg/domain/roleplay"
	"github.com/American-Made-Code/herd/pkg/external/gpt"
)

// main is the entry point of the program.
func main() {
	apiKey := os.Getenv("OPENAI_API_KEY")
	providers := []gpt.Auth{
		{Provider: gpt.OpenAi, ApiKey: apiKey},
	}

	roleplayService := roleplay.NewService(providers)

	first := roleplay.Participant{
		Index:        0,
		SystemPrompt: "You are a wizard.",
		Name:         "Gandalf",
	}

	second := roleplay.Participant{
		Index:        1,
		SystemPrompt: "You are a barbarian.",
		Name:         "Conan",
	}

	third := roleplay.Participant{
		Index:        2,
		SystemPrompt: "You are a rogue.",
		Name:         "Bilbo",
	}

	command := roleplay.CreateResponseMessageCommand{
		AllParticipants:          []roleplay.Participant{first, second, third},
		ResponseParticipantIndex: 1,
		ResponseGpt: roleplay.Gpt{
			Provider: gpt.OpenAi,
			Model:    "gpt-3.5-turbo",
		},
		Messages: []roleplay.Message{
			{
				Content:     "Hello my friends, we are on an adventure to defeat a horde of goblins. What do you bring to the table?",
				Participant: first,
			},
		},
	}

	res, err := roleplayService.CreateResponseMessage(command)

	if err != nil {
		panic(err)
	}

	fmt.Printf("\n MESSAGE:\n%v", res.GeneratedMessages[0].Content)
	fmt.Printf("\n PARTICIPANT:\n%+v", res.GeneratedMessages[0].Participant)
}
