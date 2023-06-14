package contract

import "github.com/ArtisanCloud/RobotChat/rcconfig"

type ClientInterface interface {
	// 基本配置
	GetConfig() *rcconfig.ChatBot
	SetConfig(config *rcconfig.ChatBot)

	// 请求对话
	//ChatCompletion
}
