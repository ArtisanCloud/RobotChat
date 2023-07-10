package contract

import (
	"context"
	"github.com/ArtisanCloud/RobotChat/rcconfig"
	model2 "github.com/ArtisanCloud/RobotChat/robots/artBot/model"
	"github.com/ArtisanCloud/RobotChat/robots/kernel/model"
)

// ArtBotClientInterface 是与 ArtBot 客户端交互的接口
type ArtBotClientInterface interface {

	// GetConfig 获取基本配置
	GetConfig() *rcconfig.ArtBot
	// SetConfig 设置基本配置
	SetConfig(config *rcconfig.ArtBot)

	Text2Image(ctx context.Context, message *model.Message) (*model.Message, error)
	Image2Image(ctx context.Context, message *model.Message) (*model.Message, error)
	GetModels(ctx context.Context) ([]*model2.ArtBotModel, error)
	Progress(ctx context.Context) (*model2.ProgressResponse, error)
	GetOptions(ctx context.Context) (*model2.OptionsResponse, error)
	SetOptions(ctx context.Context, options *model2.OptionsRequest) error
}
