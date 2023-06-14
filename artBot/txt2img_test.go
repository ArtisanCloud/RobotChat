package artBot

import (
	"github.com/ArtisanCloud/RobotChat/artBot/config"
	"github.com/ArtisanCloud/RobotChat/artBot/driver/Meonako/request"
	"testing"
)

func TestArtBot_Txt2Img(t *testing.T) {
	robot, err := NewArtBot(config.StableDiffusionConfig{
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
