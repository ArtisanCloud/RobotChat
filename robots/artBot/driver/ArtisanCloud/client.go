package ArtisanCloud

import (
	"context"
	"encoding/json"
	fmt "github.com/ArtisanCloud/RobotChat/pkg/printx"
	"github.com/ArtisanCloud/RobotChat/rcconfig"
	model2 "github.com/ArtisanCloud/RobotChat/robots/artBot/model"
	"github.com/ArtisanCloud/RobotChat/robots/kernel/model"
	"github.com/artisancloud/httphelper"
	"gorm.io/datatypes"
	"io"
	"net/url"
)

type Driver struct {
	config     *rcconfig.ArtBot
	httpClient httphelper.Helper
}

func NewDriver(config *rcconfig.ArtBot) *Driver {

	httpClient, _ := httphelper.NewRequestHelper(&httphelper.Config{
		BaseUrl: config.StableDiffusion.BaseUrl,
	})

	driver := &Driver{
		config:     config,
		httpClient: httpClient,
	}

	return driver
}

// GetConfig 获取基本配置
func (d *Driver) GetConfig() *rcconfig.ArtBot {
	// 实现获取基本配置的逻辑
	return d.config
}

// SetConfig 设置基本配置
func (d *Driver) SetConfig(config *rcconfig.ArtBot) {
	// 实现设置基本配置的逻辑
	d.config = config
}

func (d *Driver) Send(ctx context.Context, endpoint string, message *model.Message) (*model.Message, error) {

	requestUrl, err := d.GetUrlFromEndpoint(endpoint)
	if err != nil {
		return nil, err
	}

	res, err := d.httpClient.Df().WithContext(ctx).
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

func (d *Driver) Query(ctx context.Context, endpoint string) (*model.Message, error) {

	requestUrl, err := d.GetUrlFromEndpoint(endpoint)
	if err != nil {
		return nil, err
	}

	res, err := d.httpClient.Df().WithContext(ctx).
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

func (d *Driver) Text2Image(ctx context.Context, message *model.Message) (*model.Message, error) {
	return d.Send(ctx, "/sdapi/v1/txt2img", message)
}

func (d *Driver) Image2Image(ctx context.Context, message *model.Message) (*model.Message, error) {
	return d.Send(ctx, "/sdapi/v1/img2img", message)
}

func (d *Driver) GetModels(ctx context.Context) ([]*model2.ArtBotModel, error) {
	res, err := d.Query(ctx, "/sdapi/v1/sd-models")
	if err != nil {
		return nil, err
	}

	reply := []*model2.ArtBotModel{}
	//fmt.Dump(string(res.Content))
	err = json.Unmarshal(res.Content, &reply)

	return reply, err
}

func (d *Driver) GetLoras(ctx context.Context) (*model2.ArtBotLorasResponse, error) {
	res, err := d.Query(ctx, "/sdapi/v1/loras")
	if err != nil {
		return nil, err
	}

	reply := []*model.Lora{}
	fmt.Dump(string(res.Content))
	err = json.Unmarshal(res.Content, &reply)

	return &model2.ArtBotLorasResponse{
		Loras: reply,
	}, err
}

func (d *Driver) RefreshLoras(ctx context.Context) error {
	_, err := d.Query(ctx, "/sdapi/v1/loras")

	return err
}

func (d *Driver) Progress(ctx context.Context) (*model2.ProgressResponse, error) {
	res, err := d.Query(ctx, "/sdapi/v1/progress")
	if err != nil {
		return nil, err
	}

	reply := &model2.ProgressResponse{}
	//fmt.Dump(string(res.Content))
	err = json.Unmarshal(res.Content, reply)

	return reply, err
}

func (d *Driver) GetOptions(ctx context.Context) (*model2.OptionsResponse, error) {
	res, err := d.Query(ctx, "/sdapi/v1/options")
	if err != nil {
		return nil, err
	}

	reply := &model2.OptionsResponse{}
	//fmt.Dump(string(res.Content))
	err = json.Unmarshal(res.Content, reply)

	return reply, err
}
func (d *Driver) SetOptions(ctx context.Context, options *model2.OptionsRequest) error {
	content, err := json.Marshal(options)
	if err != nil {
		return nil
	}
	reqMessage := model.NewMessage(model.TextMessage)
	reqMessage.Content = content

	_, err = d.Send(ctx, "/sdapi/v1/options", reqMessage)

	return err
}

func (d *Driver) GetUrlFromEndpoint(endpoint string) (string, error) {
	baseUrl := d.config.StableDiffusion.BaseUrl
	urlObj, err := url.Parse(baseUrl)
	if err != nil {
		return "", err
	}

	urlObj.Path = urlObj.Path + endpoint
	finalUrl := urlObj.String()

	return finalUrl, nil
}
