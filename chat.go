package robotchat

import (
	"github.com/artisancloud/robotchat/mode"
	"github.com/artisancloud/robotchat/model"
	"github.com/artisancloud/robotchat/rbconfig"
)

type ChatRobot struct {
	text *mode.TextModeManager
}

func NewChatRobot(config rbconfig.RobotConfig) (*ChatRobot, error) {
	text, err := model.NewChatGPTModel(*config.ChatGPT)
	if err != nil {
		return nil, err
	}
	return &ChatRobot{
		text: mode.NewTextModeManager(text),
	}, nil
}

func (c *ChatRobot) TextModel() *mode.TextModeManager {
	return c.text
}

func (c *ChatRobot) TextResponse(input string) (string, error) {
	return c.text.Response(input)
}
