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
		Content:     "Hello my friends, we have been hired by the mayor of Neverwinter to investigate the disappearance of several children. Are you ready to embark on this adventure?",
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
		fmt.Printf("\n=============================\n")
		fmt.Printf("\nINDEX: %v \n", index)
		fmt.Printf("\nCURRENT MESSAGES (%v) ----------\n", len(roleplayMessages))
		for _, message := range roleplayMessages {
			fmt.Printf("\n<%v> (\n%v\n)\n", message.Participant.Name, message.Content)
		}
		fmt.Printf("\n----------\n")
		command.ResponseParticipantIndex = index
		command.Messages = roleplayMessages

		res, err := roleplayService.CreateResponseMessage(command)

		if err != nil {
			panic(err)
		}

		fmt.Printf("\nNEW MESSAGES (%v) ----------\n", len(res.GeneratedMessages))
		for _, message := range res.GeneratedMessages {
			fmt.Printf("\n<%v> (\n%v\n)\n", message.Participant.Name, message.Content)
		}
		fmt.Printf("\n----------\n")

		roleplayMessages = append(roleplayMessages, res.GeneratedMessages[0])
	}

	fmt.Printf("\n FINAL: (%v) ----------\n", len(roleplayMessages))
	for _, message := range roleplayMessages {
		fmt.Printf("\n<%v> (\n%v\n)\n", message.Participant.Name, message.Content)
	}
	fmt.Printf("\n----------\n")
}
