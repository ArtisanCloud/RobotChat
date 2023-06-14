package contract

import (
	"context"
	"github.com/ArtisanCloud/RobotChat/kernel/model"
	"github.com/ArtisanCloud/RobotChat/rcconfig"
)

// ClientInterface 是与 ChatBot 客户端交互的接口
type ClientInterface interface {
	// GetConfig 获取基本配置
	GetConfig() *rcconfig.ChatBot
	// SetConfig 设置基本配置
	SetConfig(config *rcconfig.ChatBot)

	// SendMessage 向指定对话发送消息
	SendMessage(ctx context.Context, message string, role model.Role) (string, error)

	// GetBotReply 获取指定对话的机器人回复
	GetBotReply(conversation *model.Conversation) (string, error)

	// StartModel 启动 ChatGPT 模型
	StartModel() error

	// StopModel 停止 ChatGPT 模型
	StopModel() error

	// GenerateAnswer 生成无上下文回答
	GenerateAnswer(ctx context.Context, prompt string) (string, error)

	// SetTemperature 设置模型温度
	SetTemperature(temperature float64) error

	// SetMaxAnswerLength 设置回答的最大长度
	SetMaxAnswerLength(length int) error

	// GetConversationHistory 获取指定对话的对话历史记录
	GetConversationHistory(conversation *model.Conversation) []*model.Message
}
