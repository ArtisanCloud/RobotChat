package go_openai

import (
	"context"
	"errors"
	"github.com/ArtisanCloud/RobotChat/rcconfig"
	model2 "github.com/ArtisanCloud/RobotChat/robots/kernel/model"
	"github.com/kr/pretty"
	"github.com/sashabaranov/go-openai"
	"io"
	"strings"
)

type Driver struct {
	Client *openai.Client
	config *rcconfig.ChatBot
}

func NewDriver(config *rcconfig.ChatBot) *Driver {
	openaiConfig := openai.DefaultConfig(config.ChatGPT.OpenAPIKey)
	openaiConfig.BaseURL = config.ChatGPT.BaseUrl
	c := openai.NewClientWithConfig(openaiConfig)
	driver := &Driver{
		Client: c,
		config: config,
	}

	return driver
}

// GetConfig 获取基本配置
func (d *Driver) GetConfig() *rcconfig.ChatBot {
	// 实现获取基本配置的逻辑
	return d.config
}

// SetConfig 设置基本配置
func (d *Driver) SetConfig(config *rcconfig.ChatBot) {
	// 实现设置基本配置的逻辑
	d.config = config
}

// SendMessage 向指定对话发送消息
func (d *Driver) CreateChatCompletion(ctx context.Context, message string, role model2.Role) (string, error) {
	// 实现发送消息的逻辑
	gptModel := openai.GPT3Dot5Turbo
	if d.config.ChatGPT.Model != "" {
		gptModel = d.config.ChatGPT.Model
	}
	req := openai.ChatCompletionRequest{
		Model: gptModel,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    string(role),
				Content: message,
			},
		},
	}
	res, err := d.Client.CreateChatCompletion(ctx, req)
	if err != nil {
		return "", err
	}
	return res.Choices[0].Message.Content, nil

}

func (d *Driver) CreateStreamCompletion(ctx context.Context, message string, role model2.Role) (string, error) {

	gptModel := openai.GPT3Dot5Turbo
	if d.config.ChatGPT.Model != "" {
		gptModel = d.config.ChatGPT.Model
	}

	req := openai.CompletionRequest{
		Model:     gptModel,
		User:      string(role),
		MaxTokens: 100,
		Prompt:    message,
		Stream:    true,
	}

	stream, err := d.Client.CreateCompletionStream(ctx, req)
	if err != nil {
		return "", pretty.Errorf("ChatCompletionStream error: %v", err)
	}
	defer stream.Close()

	var responseContent strings.Builder

	for {
		response, err := stream.Recv()
		if errors.Is(err, io.EOF) {
			break
		}
		if err != nil {
			return "", pretty.Errorf("Stream error: %v", err)
		}
		//fmt.Dump(response.Choices)
		responseContent.WriteString(response.Choices[0].Text)
	}

	return responseContent.String(), nil

}

// GenerateAnswer 生成无上下文回答
func (d *Driver) CreateCompletion(ctx context.Context, prompt string) (string, error) {
	// 实现生成回答的逻辑
	gptModel := openai.GPT3Ada
	if d.config.ChatGPT.Model != "" {
		gptModel = d.config.ChatGPT.Model
	}
	req := openai.CompletionRequest{
		Model:     gptModel,
		Prompt:    prompt,
		MaxTokens: 30,
	}
	//fmt.Dump(req)
	res, err := d.Client.CreateCompletion(ctx, req)
	if err != nil {
		return "", err
	}
	//fmt.Dump(res.Choices)
	return res.Choices[0].Text, nil
}

// StartModel 启动 ChatGPT 模型
func (d *Driver) StartModel() error {
	// 实现启动模型的逻辑
	return nil
}

// StopModel 停止 ChatGPT 模型
func (d *Driver) StopModel() error {
	// 实现停止模型的逻辑
	return nil
}

// SetTemperature 设置模型温度
func (d *Driver) SetTemperature(temperature float64) error {
	// 实现设置模型温度的逻辑
	return nil
}

// SetMaxAnswerLength 设置回答的最大长度
func (d *Driver) SetMaxAnswerLength(length int) error {
	// 实现设置回答最大长度的逻辑
	return nil
}

// GetConversationHistory 获取指定对话的对话历史记录
func (d *Driver) GetConversationHistory(conversation *model2.Conversation) []*model2.Message {
	// 实现获取对话历史记录的逻辑
	return nil
}
