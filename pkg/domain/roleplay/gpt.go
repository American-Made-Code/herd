package roleplay

import "github.com/American-Made-Code/herd/pkg/external/gpt"

type Gpt struct {
	Provider gpt.Provider
	Model    string
}
