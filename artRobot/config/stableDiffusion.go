package config

import (
	"github.com/artisancloud/robotchat/rbconfig"
)

type StableDiffusionConfig struct {
	rbconfig.ConfigInterface

	BaseUrl   string
	PrefixUri string
	Version   string
	HttpDebug bool
	ProxyURL  string
}

func (c *StableDiffusionConfig) GetName() string {
	return "StableDiffusion"
}

func (c *StableDiffusionConfig) Validate() error {
	return nil
}
