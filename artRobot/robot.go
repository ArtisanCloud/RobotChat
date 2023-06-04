package artRobot

import (
	"github.com/artisancloud/robotchat/artRobot/config"
	"github.com/artisancloud/robotchat/artRobot/model"
)

type ArtRobot struct {
	Txt2Img *model.Txt2ImgModel
}

func NewArtRobot(config config.StableDiffusionConfig) (*ArtRobot, error) {
	sdModel, err := model.NewTxt2ImgModel(config)
	if err != nil {
		return nil, err
	}

	return &ArtRobot{
		Txt2Img: sdModel,
	}, nil
}
