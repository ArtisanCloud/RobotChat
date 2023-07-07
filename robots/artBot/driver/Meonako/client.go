package Meonako

import (
	"context"
	"encoding/json"
	"github.com/ArtisanCloud/RobotChat/pkg/objectx"
	"github.com/ArtisanCloud/RobotChat/rcconfig"
	model2 "github.com/ArtisanCloud/RobotChat/robots/artBot/model"
	"github.com/ArtisanCloud/RobotChat/robots/kernel/model"
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

func (d *Driver) Text2Image(ctx context.Context, message *model.Message) (*model.Message, error) {

	client := api.New(api.Config{
		BaseURL: d.config.BaseUrl,
	})

	reqDriver := &api.Txt2Image{}
	err := objectx.TransformData(message.Content, reqDriver)
	if err != nil {
		return nil, err
	}
	rs, err := client.Text2Image(reqDriver)
	if err != nil {
		return nil, err
	}

	strRes, err := json.Marshal(rs)
	if err != nil {
		return nil, err
	}
	mesReply := model.NewMessage(model.TextMessage)
	mesReply.Content = strRes

	return mesReply, err
}

func (d *Driver) Image2Image(ctx context.Context, message *model.Message) (*model.Message, error) {

	client := api.New(api.Config{
		BaseURL: d.config.BaseUrl,
	})

	reqDriver := &api.Img2Img{}
	err := objectx.TransformData(message.Content, reqDriver)
	if err != nil {
		return nil, err
	}
	rs, err := client.Image2Image(reqDriver)
	if err != nil {
		return nil, err
	}

	strRes, err := json.Marshal(rs)
	if err != nil {
		return nil, err
	}
	mesReply := model.NewMessage(model.TextMessage)
	mesReply.Content = strRes

	return mesReply, err
}

func (d *Driver) Progress(ctx context.Context) (*model2.ProgressResponse, error) {

	client := api.New(api.Config{
		BaseURL: d.config.BaseUrl,
	})

	res, err := client.Progress()
	if err != nil {
		return nil, err
	}

	reply := &model2.ProgressResponse{}
	err = objectx.TransformData(res, reply)

	return reply, err
}
func (d *Driver) GetOptions(ctx context.Context) (*model2.OptionsResponse, error) {
	client := api.New(api.Config{
		BaseURL: d.config.BaseUrl,
	})

	res, err := client.Options()
	if err != nil {
		return nil, err
	}

	reply := &model2.OptionsResponse{}
	err = objectx.TransformData(res, reply)

	return reply, err
}
func (d *Driver) SetOptions(ctx context.Context, options *model2.OptionsRequest) error {
	client := api.New(api.Config{
		BaseURL: d.config.BaseUrl,
	})

	reqDriver := &api.Options{}
	err := objectx.TransformData(options, reqDriver)
	if err != nil {
		return err
	}
	err = client.SetOptions(reqDriver)

	return err
}
