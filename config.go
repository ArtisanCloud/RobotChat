package robotchat

import (
	"github.com/artisancloud/robotchat/model/chatgpt"
)

type Config interface {
	GetName() string
	Default()
	Validate() error
}

type RobotConfig struct {
	ChatGPT *chatgpt.Config
	TextModelConfig
}

type TextModelConfig struct {
	TemplateDir string
}

func (r *RobotConfig) GetName() string {
	return "Robot"
}

func (r *RobotConfig) Default() {
	return
}

func (r *RobotConfig) Validate() error {
	if r.ChatGPT == nil {
		return nil
	}
	if err := r.ChatGPT.Validate(); err != nil {
		return err
	}
	return nil
}
