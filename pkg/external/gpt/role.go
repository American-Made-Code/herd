package gpt

type Role int

const (
	User Role = iota
	Assistant
	System
)

func (r Role) String() string {
	return roleToString[r]
}

var roleToString = map[Role]string{
	User:      "user",
	Assistant: "assistant",
	System:    "system",
}

var RoleMap = map[string]Role{
	"user":      User,
	"assistant": Assistant,
	"system":    System,
}
