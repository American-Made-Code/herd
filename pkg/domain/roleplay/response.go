package roleplay

import "github.com/American-Made-Code/herd/pkg/external/gpt"

type Response struct {
	// A unique identifier from the gpt provider.
	ExternalID string
	// A list of response message choices. Can be more than one if n is greater than 1.
	GeneratedMessages []Message
	// The Unix timestamp (in seconds) of when the chat completion was Created.
	Created int
	// Number of tokens in the generated completion.
	Response_tokens int
	// Number of tokens in the prompt.
	Prompt_tokens int
	// Total number of tokens used in the request (prompt + completion).
	Total_tokens int
}

func (r Response) fromChatResponse(source *gpt.ChatResponse) Response {
	var generatedMessages []Message

	for _, choice := range source.Choices {
		generatedMessage := Message{}.fromGptMessage(choice.Message)
		generatedMessages = append(generatedMessages, generatedMessage)
	}

	r = Response{
		ExternalID:        source.ID,
		GeneratedMessages: generatedMessages,
		Created:           source.Created,
		Response_tokens:   source.Completion_tokens,
		Prompt_tokens:     source.Prompt_tokens,
		Total_tokens:      source.Total_tokens,
	}

	return r
}
func (r Response) withParticipant(source Participant) Response {
	var generatedMessages []Message

	for _, message := range r.GeneratedMessages {
		generatedMessage := message.withParticipant(source)
		generatedMessages = append(generatedMessages, generatedMessage)
	}

	r.GeneratedMessages = generatedMessages

	return r
}
