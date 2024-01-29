package openAi

import (
	"bufio"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/go-resty/resty/v2"
)

type Client interface {
	CreateChatCompletion(command CreateChatCompletionCommand) (*ChatCompletion, error)
	CreateChatCompletionStream(command CreateChatCompletionCommand) (<-chan ChatCompletionChunk, error)
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
	var res *ChatCompletion
	var err error

	// Make the post request
	resp, err := c.r.R().
		SetBody(command).
		SetResult(ChatCompletion{}).
		SetError(OpenAiErrResponse{}).
		Post("https://api.openai.com/v1/chat/completions")

	if err != nil {
		fmt.Println("RESTY ERROR:")
		fmt.Println(err)
	}

	if resp.IsSuccess() {
		res = resp.Result().(*ChatCompletion)
	}

	if resp.IsError() {
		errResponse := resp.Error().(*OpenAiErrResponse)

		fmt.Println("OPEN AI ERROR:")
		fmt.Println(errResponse)

		err = &errResponse.Error
	}

	return res, err
}

func (c *client) CreateChatCompletionStream(command CreateChatCompletionCommand) (<-chan ChatCompletionChunk, error) {
	var err error

	// Make the post request
	resp, err := c.r.R().
		SetBody(command).
		SetDoNotParseResponse(true).
		Post("https://api.openai.com/v1/chat/completions")

	if err != nil {
		fmt.Println("RESTY ERROR:")
		fmt.Println(err)
		return nil, err
	}

	if resp.Header().Get("Content-Type") != "text/event-stream" {
		fmt.Println("Invalid content type")
	}

	ch := make(chan ChatCompletionChunk)
	go func() {
		defer close(ch)
		defer resp.RawBody().Close()

		reader := bufio.NewReader(resp.RawBody())

		for {
			line, err := reader.ReadString('\n')
			if err != nil {
				// handle the error (end of stream, etc.)
				break
			}

			// Process the line (an SSE message)
			if strings.HasPrefix(line, "data:") {
				var event ChatCompletionChunk

				// Here we get the actual data from the event
				dataString := strings.TrimSpace(strings.TrimPrefix(line, "data:"))

				// Unmarshal the data into a struct
				err := json.Unmarshal([]byte(dataString), &event)
				if err != nil {
					fmt.Println("Error unmarshalling data")
					fmt.Println(err)
					continue
				}

				fmt.Printf("Received event: %+v\n", event)

				// TODO add a transform here
				ch <- event
			}
		}

	}()

	return ch, nil
}
