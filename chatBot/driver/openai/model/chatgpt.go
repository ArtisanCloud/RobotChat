package model

import (
	"github.com/ArtisanCloud/RobotChat/chatRobot/config"
	"github.com/artisancloud/openai"
	v1 "github.com/artisancloud/openai/api/v1"
)

type ChatGPTModel struct {
	client *openai.Client
}

type Message struct {
	Role    string
	Content string
}

type Choice struct {
	Reason  string
	Message Message
}

func NewChatGPTModel(config config.ChatGPTConfig) (*ChatGPTModel, error) {
	client, err := openai.NewClient(&openai.V1Config{
		OpenAPIKey:   config.OpenAPIKey,
		Organization: config.Organization,
		HttpDebug:    config.HttpDebug,
		ProxyURL:     config.ProxyURL,
	})
	if err != nil {
		return nil, err
	}
	return &ChatGPTModel{
		client: client,
	}, nil
}

func (c *ChatGPTModel) Response(messages []Message) ([]Choice, error) {
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

func (c *ChatGPTModel) HandleInput(input string) (string, error) {
	oMessages := []Message{
		{
			Role:    "user",
			Content: input,
		},
	}
	choices, err := c.Response(oMessages)
	if err != nil {
		return "", err
	}
	return choices[0].Message.Content, nil
}
