package service

import "github.com/ArtisanCloud/RobotChat/rcconfig"

func InitService(config *rcconfig.RCConfig) {
	SrvArtBot = NewArtBotService(config)
}
