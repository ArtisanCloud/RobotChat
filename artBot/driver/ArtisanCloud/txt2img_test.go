package Meonako

import (
	"context"
	artBot2 "github.com/ArtisanCloud/RobotChat/artBot"
	"github.com/ArtisanCloud/RobotChat/artBot/config"
	"github.com/ArtisanCloud/RobotChat/artBot/driver/Meonako/request"
	"github.com/ArtisanCloud/RobotChat/rcconfig"
	"testing"
)

func TestArtBot_Text2Image(t *testing.T) {
	driver := NewDriver(&rcconfig.ArtBot{
		StableDiffusionConfig: config.StableDiffusionConfig{
			BaseUrl: "http://127.0.0.1:7861",
		},
	})
	artBot, err := artBot2.NewArtBot(driver)
	if err != nil {
		t.Error(err)
	}
	ctx := context.Background()
	response, err := artBot.Client.Text2Image(ctx, &request.Text2Image{
		Prompt: "a pretty girl",
	})
	if err != nil {
		t.Error(err)
	}
	t.Log(response)
}
