package config

import (
	"errors"
)

type XFYun struct {
	WebSocketUrl string `yaml:"WebSocketUrl"`
	AppId        string `yaml:"AppId"`
	APISecret    string `yaml:"APISecret"`
	APIKey       string `yaml:"APIKey"`
	HttpDebug    bool   `yaml:"HttpDebug"`
}

func (c *XFYun) GetName() string {
	return "ChatGPT"
}

func (c *XFYun) Validate() error {
	if c.WebSocketUrl == "" {
		return errors.New("WebSocketUrl is required")
	}
	return nil
}
