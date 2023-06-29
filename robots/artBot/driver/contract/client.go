package contract

import (
	"context"
	"github.com/ArtisanCloud/RobotChat/rcconfig"
	"github.com/ArtisanCloud/RobotChat/robots/kernel/model"
)

// ArtBotClientInterface 是与 ArtBot 客户端交互的接口
type ArtBotClientInterface interface {

	// GetConfig 获取基本配置
	GetConfig() *rcconfig.ArtBot
	// SetConfig 设置基本配置
	SetConfig(config *rcconfig.ArtBot)

	Text2Image(ctx context.Context, message *model.Message) (*model.Message, error)
}
