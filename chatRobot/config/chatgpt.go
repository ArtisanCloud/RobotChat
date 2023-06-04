package config

import (
	"errors"
)

type ChatGPTConfig struct {
	OpenAPIKey   string
	Organization string
	HttpDebug    bool
	ProxyURL     string
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
