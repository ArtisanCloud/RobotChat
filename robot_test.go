package robotchat

import (
	"github.com/artisancloud/robotchat/model/chatgpt"
	"testing"
)

func TestRobot(t *testing.T) {
	rb, err := NewRobot(RobotConfig{
		ChatGPT: &chatgpt.Config{
			OpenAPIKey: "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",
			HttpDebug:  true,
			ProxyURL:   "http://127.0.0.1:33210",
		},
		//TextModelConfig: TextModelConfig{
		//	TemplateDir: "./template",
		//},
	})
	if err != nil {
		t.Fatal(err)
	}
	response, err2 := rb.TextModel.Input(`/image:prompt 我需要一张猫的图片`)
	response, err2 = rb.TextModel.Input(`/text:translation:en 你好世界`)
	if err2 != nil {
		return
	}
	t.Log(response)
}
