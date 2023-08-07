package config

import (
	"errors"
)

type THUDMGLM struct {
	HttpDebug bool   `yaml:"HttpDebug" json:",optional"`
	BaseUrl   string `yaml:"BaseUrl" json:",optional"`
}

func (c *THUDMGLM) GetName() string {
	return "ChatGPT"
}

func (c *THUDMGLM) Validate() error {
	if c.BaseUrl == "" {
		return errors.New("BaseUrl is required")
	}
	return nil
}
