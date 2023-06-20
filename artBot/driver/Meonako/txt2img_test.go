package Meonako

import (
	"context"
	artBot2 "github.com/ArtisanCloud/RobotChat/artBot"
	"github.com/ArtisanCloud/RobotChat/artBot/config"
	"github.com/ArtisanCloud/RobotChat/artBot/driver/Meonako/request"
	"github.com/ArtisanCloud/RobotChat/rcconfig"
	api "github.com/Meonako/webui-api"
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
		Prompt: "long hair, skinny, narrow waist, gothic lolita, twintails",
		NegativePrompt: api.BuildPrompt(
			"(worst quality, low quality:1.4)",
			"simple background, white background",
		),
		DoNotSendImages: true,
		BatchSize:       1,
		BatchCount:      1,
		Steps:           5,
	})
	if err != nil {
		t.Error(err)
	}
	t.Log(response)
}
