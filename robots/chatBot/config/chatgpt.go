package config

import (
	"errors"
)

type ChatGPT struct {
	OpenAPIKey   string `yaml:"OpenAPIKey" json:",optional"`
	Model        string `yaml:"Model" json:",optional"`
	Organization string `yaml:"Organization" json:",optional"`
	HttpDebug    bool   `yaml:"HttpDebug" json:",optional"`
	BaseUrl      string `yaml:"BaseUrl" json:",optional"`
	APIType      string `yaml:"APIType" json:",optional"`
	APIVersion   string `yaml:"APIVersion" json:",optional"`
}

func (c *ChatGPT) GetName() string {
	return "ChatGPT"
}

func (c *ChatGPT) Validate() error {
	if c.OpenAPIKey == "" {
		return errors.New("OpenAPIKey is required")
	}
	return nil
}
