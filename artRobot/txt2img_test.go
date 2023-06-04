package artRobot

import (
	"github.com/artisancloud/robotchat/artRobot/config"
	"github.com/artisancloud/robotchat/artRobot/driver/request"

	"testing"
)

func TestArtRobot_Txt2Img(t *testing.T) {
	robot, err := NewArtRobot(config.StableDiffusionConfig{
		BaseUrl: "http://127.0.0.1:7861",
	})
	if err != nil {
		t.Error(err)
	}
	response, err := robot.Txt2Img.Generate(&request.Txt2Image{
		Prompt: "a pretty girl",
	})
	if err != nil {
		t.Error(err)
	}
	t.Log(response)
}
