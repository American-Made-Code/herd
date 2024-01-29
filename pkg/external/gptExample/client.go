package gptExample

import (
	"fmt"

	"github.com/go-resty/resty/v2"
)

type Client interface {
	CreateChatCompletion(command CreateChatCompletionCommand) (*ChatCompletion, error)
}

type client struct {
	r *resty.Client
}

func NewClient(apiKey string) Client {
	httpClient := resty.New().
		SetAuthToken(apiKey)

	return &client{
		r: httpClient,
	}
}

func (c *client) CreateChatCompletion(command CreateChatCompletionCommand) (*ChatCompletion, error) {
	return nil, fmt.Errorf("example provider is not supported")
}
