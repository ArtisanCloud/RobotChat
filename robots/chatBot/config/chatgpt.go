package config

import (
	"errors"
)

type ChatGPTConfig struct {
	OpenAPIKey   string `yaml:"OpenAPIKey"`
	Model        string `yaml:"Model"`
	Organization string `yaml:"Organization"`
	HttpDebug    bool   `yaml:"HttpDebug"`
	BaseURL      string `yaml:"BaseURL"`
	APIType      string `yaml:"APIType"`
	APIVersion   string `yaml:"APIVersion"`
}

func (c *ChatGPTConfig) GetName() string {
	return "ChatGPT"
}

func (c *ChatGPTConfig) Validate() error {
	if c.OpenAPIKey == "" {
		return errors.New("OpenAPIKey is required")
	}
	return nil
}
