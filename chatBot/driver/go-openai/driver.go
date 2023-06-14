package go_openai

import (
	"context"
	"github.com/ArtisanCloud/RobotChat/kernel/model"
	"github.com/ArtisanCloud/RobotChat/rcconfig"
	"github.com/sashabaranov/go-openai"
)

type Driver struct {
	Client *openai.Client
	config *rcconfig.ChatBot
}

func NewDriver(config *rcconfig.ChatBot) *Driver {
	openaiConfig := openai.DefaultConfig(config.OpenAPIKey)
	openaiConfig.BaseURL = config.BaseURL
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
func (d *Driver) SendMessage(ctx context.Context, message string, role model.Role) (string, error) {
	// 实现发送消息的逻辑
	req := openai.ChatCompletionRequest{
		Model: openai.GPT3Dot5Turbo,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleUser,
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

// GetBotReply 获取指定对话的机器人回复
func (d *Driver) GetBotReply(conversation *model.Conversation) (string, error) {
	// 实现获取机器人回复的逻辑
	return "", nil
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

// GenerateAnswer 生成无上下文回答
func (d *Driver) GenerateAnswer(ctx context.Context, prompt string) (string, error) {
	// 实现生成回答的逻辑
	req := openai.CompletionRequest{
		Model:  openai.GPT3Ada,
		Prompt: prompt,
	}
	res, err := d.Client.CreateCompletion(ctx, req)
	if err != nil {
		return "", err
	}
	return res.Choices[0].Text, nil
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
func (d *Driver) GetConversationHistory(conversation *model.Conversation) []*model.Message {
	// 实现获取对话历史记录的逻辑
	return nil
}
