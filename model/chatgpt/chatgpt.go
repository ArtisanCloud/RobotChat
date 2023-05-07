package chatgpt

import (
	"errors"
)

type Config struct {
	OpenAPIKey   string
	Organization string
	HttpDebug    bool
	ProxyURL     string
	Model        ModelConfig
}

type ModelConfig struct {
	ChatCompletionModel string
}

func (c *Config) GetName() string {
	return "ChatGPT"
}

func (c *Config) Default() {
	if c.Model.ChatCompletionModel == "" {
		c.Model.ChatCompletionModel = "gpt-3.5-turbo"
	}
}

func (c *Config) Validate() error {
	if c.OpenAPIKey == "" {
		return errors.New("OpenAPIKey is required")
	}
	return nil
}
