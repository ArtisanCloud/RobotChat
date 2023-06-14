package artBot

import (
	"github.com/ArtisanCloud/RobotChat/artBot/config"
	"github.com/ArtisanCloud/RobotChat/artBot/model"
)

type ArtBot struct {
	Txt2Img *model.Txt2ImgModel
}

func NewArtBot(config config.StableDiffusionConfig) (*ArtBot, error) {
	sdModel, err := model.NewTxt2ImgModel(config)
	if err != nil {
		return nil, err
	}

	return &ArtBot{
		Txt2Img: sdModel,
	}, nil
}
