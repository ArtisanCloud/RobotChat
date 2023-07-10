package config

import (
	"errors"
)

type ChatGPT struct {
	OpenAPIKey   string `yaml:"OpenAPIKey"`
	Model        string `yaml:"Model"`
	Organization string `yaml:"Organization"`
	HttpDebug    bool   `yaml:"HttpDebug"`
	BaseUrl      string `yaml:"BaseUrl"`
	APIType      string `yaml:"APIType"`
	APIVersion   string `yaml:"APIVersion"`
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
