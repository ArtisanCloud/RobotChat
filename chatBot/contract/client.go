package contract

import "github.com/ArtisanCloud/RobotChat/rcconfig"

type ClientInterface interface {
	GetConfig() *rcconfig.ChatBot
	SetConfig(config *rcconfig.ChatBot)
}
