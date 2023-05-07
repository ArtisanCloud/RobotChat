package chatgpt

import (
	"github.com/artisancloud/openai"
	v1 "github.com/artisancloud/openai/api/v1"
)

type ChatGPT struct {
	client *openai.Client
	conf   *Config
}

func NewChatGPTModel(config Config) (*ChatGPT, error) {
	client, err := openai.NewClient(&openai.V1Config{
		OpenAPIKey:   config.OpenAPIKey,
		Organization: config.Organization,
		HttpDebug:    config.HttpDebug,
		ProxyURL:     config.ProxyURL,
	})
	if err != nil {
		return nil, err
	}
	return &ChatGPT{
		client: client,
		conf:   &config,
	}, nil
}

type Message struct {
	Role    string
	Content string
}

type Choice struct {
	Reason  string
	Message Message
}

func (c *ChatGPT) Response(messages []Message) ([]Choice, error) {
	var oMessages []v1.Message
	for _, message := range messages {
		oMessages = append(oMessages, v1.Message{
			Role:    message.Role,
			Content: message.Content,
		})
	}
	completion, err := c.client.V1.Chat.CreateChatCompletion(&v1.CreateChatCompletionRequest{
		Model:    "gpt-3.5-turbo",
		Messages: oMessages,
	})
	if err != nil {
		return nil, err
	}
	choices := make([]Choice, 0, len(completion.Choices))
	for _, choice := range completion.Choices {
		choices = append(choices, Choice{
			Reason: choice.FinishReason,
			Message: Message{
				Role:    choice.Message.Role,
				Content: choice.Message.Content,
			},
		})
	}
	return choices, nil
}
