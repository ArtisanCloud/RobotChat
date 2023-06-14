package artRobot

import (
	"github.com/ArtisanCloud/RobotChat/artRobot/config"
	"github.com/ArtisanCloud/RobotChat/artRobot/model"
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
