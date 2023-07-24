package ArtisanCloud

import (
	"context"
	"encoding/json"
	"github.com/ArtisanCloud/RobotChat/rcconfig"
	"github.com/ArtisanCloud/RobotChat/robots/kernel/model"
	"github.com/artisancloud/httphelper"
	"gorm.io/datatypes"
	"io"
)

type BaseDriver struct {
	HttpClient httphelper.Helper
	Config     *rcconfig.ChatBot

	GetUrlFromEndpoint func(endpoint string) (string, error)
}

// GetConfig 获取基本配置
func (d *BaseDriver) GetConfig() *rcconfig.ChatBot {
	// 实现获取基本配置的逻辑
	return d.Config
}

// SetConfig 设置基本配置
func (d *BaseDriver) SetConfig(config *rcconfig.ChatBot) {
	// 实现设置基本配置的逻辑
	d.Config = config
}

func (d *BaseDriver) Send(ctx context.Context, endpoint string, message *model.Message) (*model.Message, error) {

	requestUrl, err := d.GetUrlFromEndpoint(endpoint)
	if err != nil {
		return nil, err
	}

	res, err := d.HttpClient.Df().WithContext(ctx).
		Url(requestUrl).
		Method("POST").
		Json(message.Content).
		Request()
	if err != nil {
		return nil, err
	}
	msg := model.NewMessage(model.TextMessage)

	// 转化返回的Body
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	var bodyData datatypes.JSON
	err = json.Unmarshal(body, &bodyData)
	msg.Content = bodyData

	// 转化返回的Header
	headerJSON, err := json.Marshal(res.Header)
	if err != nil {
		return nil, err
	}

	var headerDataJSON datatypes.JSON
	err = json.Unmarshal(headerJSON, &headerDataJSON)
	if err != nil {
		return nil, err
	}

	return msg, nil
}

func (d *BaseDriver) Query(ctx context.Context, endpoint string) (*model.Message, error) {

	requestUrl, err := d.GetUrlFromEndpoint(endpoint)
	if err != nil {
		return nil, err
	}

	res, err := d.HttpClient.Df().WithContext(ctx).
		Url(requestUrl).
		Method("GET").
		Request()
	if err != nil {
		return nil, err
	}
	msg := model.NewMessage(model.TextMessage)

	// 转化返回的Body
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	var bodyData datatypes.JSON
	err = json.Unmarshal(body, &bodyData)
	msg.Content = bodyData

	// 转化返回的Header
	headerJSON, err := json.Marshal(res.Header)
	if err != nil {
		return nil, err
	}

	var headerDataJSON datatypes.JSON
	err = json.Unmarshal(headerJSON, &headerDataJSON)
	if err != nil {
		return nil, err
	}

	return msg, nil
}
