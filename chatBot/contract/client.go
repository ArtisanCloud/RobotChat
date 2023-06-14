package contract

import "github.com/ArtisanCloud/RobotChat/rcconfig"

type ClientInterface interface {
	GetConfig() *rcconfig.ChatRobot
	SetConfig(config *rcconfig.ChatRobot)
}
