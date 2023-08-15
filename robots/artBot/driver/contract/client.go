package contract

import (
	"context"
	"github.com/ArtisanCloud/RobotChat/rcconfig"
	model2 "github.com/ArtisanCloud/RobotChat/robots/artBot/model"
	"github.com/ArtisanCloud/RobotChat/robots/artBot/model/controlNet"
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
	GetSamplers(ctx context.Context) ([]*model2.Sampler, error)
	GetLoras(ctx context.Context) ([]*model.Lora, error)
	RefreshLoras(ctx context.Context) error
	Progress(ctx context.Context) (*model2.ProgressResponse, error)
	GetOptions(ctx context.Context) (*model2.OptionsResponse, error)
	SetOptions(ctx context.Context, options *model2.OptionsRequest) error

	GetControlNetModelList(ctx context.Context) (*controlNet.ControlNetModel, error)
	GetControlNetModuleList(ctx context.Context) (*controlNet.Modules, error)
	GetControlNetVersion(ctx context.Context) (*controlNet.ControlNetVersion, error)
	GetControlNetSettings(ctx context.Context) (*controlNet.ControlNetSettings, error)
	DetectControlNet(ctx context.Context, info *controlNet.DetectInfo) (interface{}, error)
}
