package model

import "github.com/ArtisanCloud/RobotChat/robots/kernel/model"

type ArtBotLorasResponse struct {
	SDResponse
	Loras []*model.Lora
}
