package chatGLM

import (
	"context"
	"encoding/json"
	"github.com/ArtisanCloud/RobotChat/rcconfig"
	"github.com/ArtisanCloud/RobotChat/robots/chatBot/driver/ArtisanCloud"
	"github.com/ArtisanCloud/RobotChat/robots/chatBot/model"
	"github.com/ArtisanCloud/RobotChat/robots/kernel/logger"
	model2 "github.com/ArtisanCloud/RobotChat/robots/kernel/model"
	"github.com/artisancloud/httphelper"
	"net/url"
)

type Driver struct {
	*ArtisanCloud.BaseDriver
}

func NewDriver(config *rcconfig.ChatBot) *Driver {

	httpClient, _ := httphelper.NewRequestHelper(&httphelper.Config{
		BaseUrl: config.THUDMGLM.BaseUrl,
	})

	baseDriver := ArtisanCloud.NewDriver(config)
	baseDriver.Config = config
	baseDriver.HttpClient = httpClient

	log, _ := logger.NewLogger(nil, config.Log)
	baseDriver.Logger = log

	driver := &Driver{
		BaseDriver: baseDriver,
	}

	driver.GetUrlFromEndpoint = driver.OverrideGetUrlFromEndpoint()
	driver.OverrideGetMiddlewares()
	driver.RegisterHttpMiddlewares()

	return driver
}

// SendMessage 向指定对话发送消息
func (d *Driver) CreateChatCompletion(ctx context.Context, message *model2.Message, role model2.Role) (*model2.Message, error) {

	req := GLMRequest{}
	req.Prompt = string(message.Content)
	strReq, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	message.Content = strReq

	return d.Send(ctx, "/", message)

}

func (d *Driver) CreateChatCompletionStream(ctx context.Context, message *model2.Message, role model2.Role, processStreamData func(data string, status model.ChatStatus) error) (*model2.Message, error) {
	return nil, nil
}

// GenerateAnswer 生成无上下文回答
func (d *Driver) CreateCompletion(ctx context.Context, message *model2.Message) (*model2.Message, error) {
	return d.Send(ctx, "/", message)

}

func (d *Driver) CreateCompletionStream(ctx context.Context, message *model2.Message, role model2.Role, processStreamData func(data string, status model.ChatStatus) error) (*model2.Message, error) {
	return nil, nil
}

// StartModel 启动 ChatGPT 模型
func (d *Driver) StartModel() error {
	// 实现启动模型的逻辑
	return nil
}

// StopModel 停止 ChatGPT 模型
func (d *Driver) StopModel() error {
	// 实现停止模型的逻辑
	return nil
}

// SetTemperature 设置模型温度
func (d *Driver) SetTemperature(temperature float64) error {
	// 实现设置模型温度的逻辑
	return nil
}

// SetMaxAnswerLength 设置回答的最大长度
func (d *Driver) SetMaxAnswerLength(length int) error {
	// 实现设置回答最大长度的逻辑
	return nil
}

func (d *Driver) OverrideGetUrlFromEndpoint() func(endpoint string) (string, error) {

	return func(endpoint string) (string, error) {
		baseUrl := d.Config.THUDMGLM.BaseUrl

		urlObj, err := url.Parse(baseUrl)
		if err != nil {
			return "", err
		}

		urlObj.Path = urlObj.Path + endpoint
		finalUrl := urlObj.String()

		return finalUrl, nil
	}

}
