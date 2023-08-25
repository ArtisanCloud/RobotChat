package contract

import (
	"context"
	"github.com/ArtisanCloud/RobotChat/rcconfig"
	"github.com/ArtisanCloud/RobotChat/robots/kernel/model"
)

// ChatBotClientInterface 是与 ChatBot 客户端交互的接口
type ChatBotClientInterface interface {
	// GetConfig 获取基本配置
	GetConfig() *rcconfig.ChatBot
	// SetConfig 设置基本配置
	SetConfig(config *rcconfig.ChatBot)

	CreateCompletion(ctx context.Context, message *model.Message) (*model.Message, error)
	CreateCompletionStream(ctx context.Context, message *model.Message, role model.Role, processStreamData func(data string) error) (*model.Message, error)
	CreateChatCompletion(ctx context.Context, message *model.Message, role model.Role) (*model.Message, error)
	CreateChatCompletionStream(ctx context.Context, message *model.Message, role model.Role, processStreamData func(data string) error) (*model.Message, error)

	// SetTemperature 设置模型温度
	SetTemperature(temperature float64) error

	// SetMaxAnswerLength 设置回答的最大长度
	SetMaxAnswerLength(length int) error
}
