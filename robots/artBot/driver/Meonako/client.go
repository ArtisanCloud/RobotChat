package Meonako

import (
	"context"
	"github.com/ArtisanCloud/RobotChat/pkg/objectx"
	"github.com/ArtisanCloud/RobotChat/rcconfig"
	"github.com/ArtisanCloud/RobotChat/robots/artBot/driver/Meonako/request"
	"github.com/ArtisanCloud/RobotChat/robots/artBot/driver/Meonako/response"
	api "github.com/Meonako/webui-api"
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

func (d *Driver) Text2Image(ctx context.Context, req *request.Text2Image) (*response.Text2Image, error) {

	client := api.New(api.Config{
		BaseURL: d.config.BaseUrl,
	})

	reqDriver := &api.Txt2Image{}
	err := objectx.TransformData(req, reqDriver)
	if err != nil {
		return nil, err
	}
	rs, err := client.Text2Image(reqDriver)

	return &response.Text2Image{
		Images:        rs.Images,
		DecodedImages: rs.DecodedImages,
		Parameters:    rs.Parameters,
		Info:          rs.Info,
	}, err
}
