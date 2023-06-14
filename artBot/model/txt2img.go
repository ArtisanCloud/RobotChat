package model

import (
	"github.com/ArtisanCloud/RobotChat/artBot/config"
	"github.com/ArtisanCloud/RobotChat/artBot/driver/Meonako"
	"github.com/ArtisanCloud/RobotChat/artBot/driver/Meonako/request"
	"github.com/ArtisanCloud/RobotChat/artBot/driver/Meonako/response"
)

type Txt2ImgModel struct {
	Client *Meonako.Client
}

func NewTxt2ImgModel(config config.StableDiffusionConfig) (*Txt2ImgModel, error) {

	client := Meonako.NewClient(config)

	return &Txt2ImgModel{
		Client: client,
	}, nil
}

func (m *Txt2ImgModel) Generate(req *request.Txt2Image) (*response.Txt2Image, error) {
	return m.Client.Txt2Image(req)
}
