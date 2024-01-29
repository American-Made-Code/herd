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

	allParticipants := []roleplay.Participant{first, second, third}

	var roleplayMessages []roleplay.Message
	initialMessage := roleplay.Message{
		Content:     "Hello my friends, we are on an adventure to defeat a horde of goblins. What do you bring to the table?",
		Participant: first,
	}
	roleplayMessages = append(roleplayMessages, initialMessage)

	responseOrder := []int{1, 2, 0, 1, 2, 0}

	command := roleplay.CreateResponseMessageCommand{
		AllParticipants:          allParticipants,
		ResponseParticipantIndex: 1,
		ResponseGpt: roleplay.Gpt{
			Provider: gpt.OpenAi,
			Model:    "gpt-3.5-turbo",
		},
		Messages: roleplayMessages,
	}

	for _, index := range responseOrder {
		command.ResponseParticipantIndex = index
		res, err := roleplayService.CreateResponseMessage(command)

		if err != nil {
			panic(err)
		}

		fmt.Printf("\n<%v> \t%v", res.GeneratedMessages[0].Participant.Name, res.GeneratedMessages[0].Content)

		roleplayMessages = append(roleplayMessages, res.GeneratedMessages[0])
	}

	fmt.Printf("\n\n DONE!")
}
