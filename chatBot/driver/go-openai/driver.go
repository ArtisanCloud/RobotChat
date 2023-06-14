package go_openai

import (
	"github.com/ArtisanCloud/RobotChat/rcconfig"
	"github.com/sashabaranov/go-openai"
)

type Driver struct {
	Client *openai.Client
	config *rcconfig.ChatBot
}

func NewDriver(config *rcconfig.ChatBot) *Driver {
	openaiConfig := openai.DefaultConfig(config.OpenAPIKey)
	openaiConfig.BaseURL = config.BaseURL
	openaiConfig.OrgID = config.Organization
	openaiConfig.APIType = openai.APIType(config.APIType)
	openaiConfig.APIVersion = config.APIVersion
	c := openai.NewClientWithConfig(openaiConfig)

	driver := &Driver{
		Client: c,
		config: config,
	}

	return driver
}

func (d Driver) GetConfig() *rcconfig.ChatBot {
	return d.config
}
func (d Driver) SetConfig(config *rcconfig.ChatBot) {
	d.config = config
}
