package config

import (
	"errors"
)

type XFYunConfig struct {
	WebSocketUrl string `yaml:"WebSocketUrl"`
	AppId        string `yaml:"AppId"`
	APISecret    string `yaml:"APISecret"`
	APIKey       string `yaml:"APIKey"`
	HttpDebug    bool   `yaml:"HttpDebug"`
}

func (c *XFYunConfig) GetName() string {
	return "ChatGPT"
}

func (c *XFYunConfig) Validate() error {
	if c.WebSocketUrl == "" {
		return errors.New("WebSocketUrl is required")
	}
	return nil
}
