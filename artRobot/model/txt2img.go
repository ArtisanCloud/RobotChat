package model

import (
	"github.com/artisancloud/robotchat/artRobot/config"
	"github.com/artisancloud/robotchat/artRobot/driver"
	"github.com/artisancloud/robotchat/artRobot/driver/request"
	"github.com/artisancloud/robotchat/artRobot/driver/response"
)

type Txt2ImgModel struct {
	Client *driver.Client
}

func NewTxt2ImgModel(config config.StableDiffusionConfig) (*Txt2ImgModel, error) {

	client := driver.NewClient(config)

	return &Txt2ImgModel{
		Client: client,
	}, nil
}

func (m *Txt2ImgModel) Generate(req *request.Txt2Image) (*response.Txt2Image, error) {
	return m.Client.Txt2Image(req)
}
