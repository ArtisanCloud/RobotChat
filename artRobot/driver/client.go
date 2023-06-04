package driver

import (
	webapi "github.com/Meonako/webui-api"
	"github.com/artisancloud/robotchat/artRobot/config"
	"github.com/artisancloud/robotchat/artRobot/driver/request"
	"github.com/artisancloud/robotchat/artRobot/driver/response"
)

type Client struct {
	Config config.StableDiffusionConfig
}

func NewClient(config config.StableDiffusionConfig) *Client {
	return &Client{
		Config: config,
	}
}

func (c *Client) Txt2Image(req *request.Txt2Image) (*response.Txt2Image, error) {

	client := webapi.New(webapi.Config{
		BaseURL: c.Config.BaseUrl,
	})

	rs, err := client.Text2Image(&webapi.Txt2Image{
		Prompt: req.Prompt,
	})

	return &response.Txt2Image{
		Images:        rs.Images,
		DecodedImages: rs.DecodedImages,
		Parameters:    rs.Parameters,
		Info:          rs.Info,
	}, err
}
