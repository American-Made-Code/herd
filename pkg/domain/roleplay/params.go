package roleplay

type CreateResponseMessageCommand struct {
	AllParticipants          []Participant
	ResponseParticipantIndex int
	ResponseGpt              Gpt
	Messages                 []Message
}
