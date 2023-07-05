package config

import (
	"errors"
)

type THUDMGLMConfig struct {
	HttpDebug bool   `yaml:"HttpDebug"`
	BaseUrl   string `yaml:"BaseUrl"`
}

func (c *THUDMGLMConfig) GetName() string {
	return "ChatGPT"
}

func (c *THUDMGLMConfig) Validate() error {
	if c.BaseUrl == "" {
		return errors.New("BaseUrl is required")
	}
	return nil
}
