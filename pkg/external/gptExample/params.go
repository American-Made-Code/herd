package gptExample

type CreateChatCompletionCommand struct {
	// A list of Messages comprising the conversation so far.
	// [Example Python Code]
	// (https://cookbook.openai.com/examples/how_to_format_inputs_to_chatgpt_models)
	Messages []Message `json:"messages"`
	// ID of the Model to use. See the
	// [Model endpoint compatibility table]
	// (https://platform.openai.com/docs/models/Model-endpoint-compatibility)
	// for details on which models work with the Chat API.
	Model string `json:"model"`
	// Number between -2.0 and 2.0. Positive values penalize new tokens
	// based on their existing frequency in the text so far, decreasing the model's
	// likelihood to repeat the same line verbatim.
	// [See more information about frequency and presence penalties.]
	// (https://platform.openai.com/docs/guides/text-generation/parameter-details)
	Frequency_penalty *float32 `json:"frequency_penalty,omitempty"`
	// The maximum number of [tokens] (https://platform.openai.com/tokenizer)
	// that can be generated in the chat completion. The total length of input
	// tokens and generated tokens is limited by the model's context length.
	// [Example Python code for counting tokens.]
	// (https://cookbook.openai.com/examples/how_to_count_tokens_with_tiktoken)
	Max_tokens *int `json:"max_tokens,omitempty"`
	// How many chat completion choices to generate for each input message.
	// Note that you will be charged based on the number of generated tokens
	//  across all of the choices. Keep N as 1 to minimize costs.
	N *int `json:"n,omitempty"`
	// Number between -2.0 and 2.0. Positive values penalize new tokens based
	// on whether they appear in the text so far, increasing the model's
	// likelihood to talk about new topics.
	// [See more information about frequency and presence penalties.]
	// (https://platform.openai.com/docs/guides/text-generation/parameter-details)
	Presence_penalty *float32 `json:"presence_penalty,omitempty"`
	// An object specifying the format that the model must output.
	// Compatible with gpt-4-1106-preview and gpt-3.5-turbo-1106.
	// Setting to { "type": "json_object" } enables JSON mode,
	// which guarantees the message the model generates is valid JSON.
	// Important: when using JSON mode, you must also instruct the model to
	// produce JSON yourself via a system or user message. Without this,
	// the model may generate an unending stream of whitespace until the generation
	// reaches the token limit, resulting in a long-running and seemingly "stuck" request.
	// Also note that the message content may be partially cut off if
	// finish_reason="length", which indicates the generation exceeded max_tokens
	// or the conversation exceeded the max context length.
	Response_format *ResponseFormat `json:"response_format,omitempty"`
	// Up to 4 sequences where the API will Stop generating further tokens.
	Stop *[]string `json:"stop,omitempty"`
	// If set, partial message deltas will be sent, like in ChatGPT.
	// Tokens will be sent as data-only [server-sent events]
	// (https://developer.mozilla.org/en-US/docs/Web/API/Server-sent_events/Using_server-sent_events#Event_stream_format)
	// as they become available, with the Stream terminated by a data:
	// [DONE] message. [Example Python code.]
	// (https://cookbook.openai.com/examples/how_to_stream_completions)
	Stream *bool `json:"stream,omitempty"`
	// What sampling Temperature to use, between 0 and 2. Higher values
	// like 0.8 will make the output more random, while lower values like 0.2
	// will make it more focused and deterministic. We generally recommend
	// altering this or top_p but not both.
	Temperature *float32 `json:"temperature,omitempty"`
	// An alternative to sampling with temperature, called nucleus sampling,
	// where the model considers the results of the tokens with Top_p
	// probability mass. So 0.1 means only the tokens comprising the top
	// 10% probability mass are considered. We generally recommend altering
	// this or temperature but not both.
	Top_p *float32 `json:"top_p,omitempty"`
	// A unique identifier representing your end-User, which can help OpenAI
	// to monitor and detect abuse. [Learn more.]
	// (https://platform.openai.com/docs/guides/safety-best-practices/end-User-ids)
	User *string `json:"user,omitempty"`
}
