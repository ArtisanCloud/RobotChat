package Meonako

import (
	"context"
	"encoding/json"
	"github.com/ArtisanCloud/RobotChat/pkg/objectx"
	"github.com/ArtisanCloud/RobotChat/rcconfig"
	model2 "github.com/ArtisanCloud/RobotChat/robots/artBot/model"
	"github.com/ArtisanCloud/RobotChat/robots/artBot/model/controlNet"
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
		BaseURL: d.config.StableDiffusion.BaseUrl,
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
		BaseURL: d.config.StableDiffusion.BaseUrl,
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

func (d *Driver) GetModels(ctx context.Context) ([]*model2.ArtBotModel, error) {
	client := api.New(api.Config{
		BaseURL: d.config.StableDiffusion.BaseUrl,
	})

	res, err := client.SDModels()
	if err != nil {
		return nil, err
	}

	models := []*model2.ArtBotModel{}
	err = objectx.TransformData(res, models)

	return models, err
}

func (d *Driver) GetSamplers(ctx context.Context) ([]*model2.Sampler, error) {
	return nil, nil
}

func (d *Driver) GetLoras(ctx context.Context) ([]*model.Lora, error) {

	return nil, nil
}

func (d *Driver) RefreshLoras(ctx context.Context) error {

	return nil
}

func (d *Driver) Progress(ctx context.Context) (*model2.ProgressResponse, error) {

	client := api.New(api.Config{
		BaseURL: d.config.StableDiffusion.BaseUrl,
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
		BaseURL: d.config.StableDiffusion.BaseUrl,
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
		BaseURL: d.config.StableDiffusion.BaseUrl,
	})

	reqDriver := &api.Options{}
	err := objectx.TransformData(options, reqDriver)
	if err != nil {
		return err
	}
	err = client.SetOptions(reqDriver)

	return err
}

func (d *Driver) GetControlNetModelList(ctx context.Context) (*controlNet.ControlNetModels, error) {
	return nil, nil
}
func (d *Driver) GetControlNetModuleList(ctx context.Context) (*controlNet.Modules, error) {
	return nil, nil
}
func (d *Driver) GetControlNetControlTypesList(ctx context.Context) (*controlNet.ControlNetTypes, error) {
	return nil, nil
}
func (d *Driver) GetControlNetVersion(ctx context.Context) (*controlNet.ControlNetVersion, error) {
	return nil, nil
}
func (d *Driver) GetControlNetSettings(ctx context.Context) (*controlNet.ControlNetSettings, error) {
	return nil, nil
}
func (d *Driver) DetectControlNet(ctx context.Context, info *controlNet.DetectInfo) (interface{}, error) {
	return nil, nil
}
