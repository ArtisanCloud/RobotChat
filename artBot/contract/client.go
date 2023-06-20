package contract

import (
	"context"
	"github.com/ArtisanCloud/RobotChat/artBot/driver/Meonako/request"
	"github.com/ArtisanCloud/RobotChat/artBot/driver/Meonako/response"
	"github.com/ArtisanCloud/RobotChat/rcconfig"
)

// ClientInterface 是与 ArtBot 客户端交互的接口
type ClientInterface interface {
	// GetConfig 获取基本配置
	GetConfig() *rcconfig.ArtBot
	// SetConfig 设置基本配置
	SetConfig(config *rcconfig.ArtBot)

	Text2Image(ctx context.Context, req *request.Text2Image) (*response.Text2Image, error)
}
