package contract

import (
	"context"
	"github.com/ArtisanCloud/RobotChat/rcconfig"
	"github.com/ArtisanCloud/RobotChat/robots/kernel/model"
)

// ClientInterface 是与 ChatBot 客户端交互的接口
type ClientInterface interface {

	// GetConfig 获取基本配置
	GetConfig() *rcconfig.ChatBot
	// SetConfig 设置基本配置
	SetConfig(config *rcconfig.ChatBot)

	CreateChatCompletion(ctx context.Context, message string, role model.Role) (string, error)
	CreateStreamCompletion(ctx context.Context, message string, role model.Role) (string, error)
	CreateCompletion(ctx context.Context, prompt string) (string, error)

	// SetTemperature 设置模型温度
	SetTemperature(temperature float64) error

	// SetMaxAnswerLength 设置回答的最大长度
	SetMaxAnswerLength(length int) error
}
