package ArtisanCloud

import (
	"context"
	"github.com/ArtisanCloud/RobotChat/rcconfig"
	"github.com/ArtisanCloud/RobotChat/robots/kernel/model"
)

type Driver struct {
	config *rcconfig.ArtBot
}

func NewDriver(config *rcconfig.ArtBot) *Driver {

	driver := &Driver{
		config: config,
	}

	return driver
}

// GetConfig 获取基本配置
func (d *Driver) GetConfig() *rcconfig.ArtBot {
	// 实现获取基本配置的逻辑
	return d.config
}

// SetConfig 设置基本配置
func (d *Driver) SetConfig(config *rcconfig.ArtBot) {
	// 实现设置基本配置的逻辑
	d.config = config
}

func (d *Driver) Text2Image(ctx context.Context, message *model.Message) (*model.Message, error) {

	return nil, nil
}
