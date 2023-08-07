package config

import (
	"errors"
)

type XFYun struct {
	WebSocketUrl string `yaml:"WebSocketUrl" json:",optional"`
	AppId        string `yaml:"AppId" json:",optional"`
	APISecret    string `yaml:"APISecret" json:",optional"`
	APIKey       string `yaml:"APIKey" json:",optional"`
	HttpDebug    bool   `yaml:"HttpDebug" json:",optional"`
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
