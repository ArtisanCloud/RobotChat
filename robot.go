package robotchat

import (
	"embed"
	"github.com/artisancloud/robotchat/input"
	"github.com/artisancloud/robotchat/model"
	"github.com/artisancloud/robotchat/model/chatgpt"
)

//go:embed template/*.json
var tmplDir embed.FS

type Robot struct {
	TextModel model.TextModel
}

func NewRobot(conf RobotConfig) (*Robot, error) {
	robot := new(Robot)
	if conf.ChatGPT != nil {
		textModel, err := chatgpt.NewChatGPTModel(*conf.ChatGPT)
		if err != nil {
			return nil, err
		}
		robot.TextModel = chatgpt.NewTextAdapter(*textModel)
	}
	// load mode
	tmplSlice, err := input.LoadTemplateForEmbed(tmplDir, "template/text.json")
	if err != nil {
		return nil, err
	}

	for _, tmpl := range tmplSlice {
		robot.TextModel.RegisterMode(tmpl)
	}

	if conf.TextModelConfig.TemplateDir != "" {
		tmplSlice, err := input.LoadTemplate(conf.TextModelConfig.TemplateDir)
		if err != nil {
			return nil, err
		}

		for _, tmpl := range tmplSlice {
			robot.TextModel.RegisterMode(tmpl)
		}
	}

	return robot, nil
}
