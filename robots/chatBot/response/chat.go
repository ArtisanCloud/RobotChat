package response

import (
	"github.com/ArtisanCloud/RobotChat/robots/chatBot"
	"github.com/ArtisanCloud/RobotChat/robots/chatBot/request"
)

type ChatCompletionChoice struct {
	Index        int                           `json:"index"`
	Message      request.ChatCompletionMessage `json:"message"`
	FinishReason string                        `json:"finish_reason"`
}

// ChatCompletionResponse represents a response structure for chat completion API.
type ChatCompletionResponse struct {
	ID      string                 `json:"id"`
	Object  string                 `json:"object"`
	Created int64                  `json:"created"`
	Model   string                 `json:"model"`
	Choices []ChatCompletionChoice `json:"choices"`
	Usage   chatBot.Usage          `json:"usage"`
}
