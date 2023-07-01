package ArtisanCloud

import (
	"context"
	"encoding/json"
	"github.com/ArtisanCloud/RobotChat/rcconfig"
	artBot2 "github.com/ArtisanCloud/RobotChat/robots/artBot"
	"github.com/ArtisanCloud/RobotChat/robots/artBot/config"
	"github.com/ArtisanCloud/RobotChat/robots/artBot/request"
	"github.com/ArtisanCloud/RobotChat/robots/kernel/model"
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

	req := &request.Text2Image{
		Prompt: "long hair, skinny, narrow waist, gothic lolita, twintails",
		NegativePrompt: api.BuildPrompt(
			"(worst quality, low quality:1.4)",
			"simple background, white background",
		),
		DoNotSendImages: true,
		BatchSize:       1,
		BatchCount:      1,
		Steps:           5,
	}
	strReq, err := json.Marshal(req)
	if err != nil {
		t.Error(err)
	}

	ctx := context.Background()
	response, err := artBot.Client.Text2Image(ctx, &model.Message{
		Content: strReq,
	})
	if err != nil {
		t.Error(err)
	}
	t.Log(response)
}
