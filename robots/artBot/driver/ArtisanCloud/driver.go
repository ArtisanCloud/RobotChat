package ArtisanCloud

import (
	"context"
	"encoding/json"
	"github.com/ArtisanCloud/RobotChat/rcconfig"
	model2 "github.com/ArtisanCloud/RobotChat/robots/artBot/model"
	"github.com/ArtisanCloud/RobotChat/robots/artBot/model/controlNet"
	"github.com/ArtisanCloud/RobotChat/robots/kernel/logger"
	contract2 "github.com/ArtisanCloud/RobotChat/robots/kernel/logger/contract"
	"github.com/ArtisanCloud/RobotChat/robots/kernel/model"
	request2 "github.com/ArtisanCloud/RobotChat/robots/kernel/request"
	response2 "github.com/ArtisanCloud/RobotChat/robots/kernel/response"
	"github.com/artisancloud/httphelper"
	"github.com/artisancloud/httphelper/client"
	"github.com/artisancloud/httphelper/dataflow"
	"gorm.io/datatypes"
	"io"
	"net/http"
	"net/url"
	"time"
)

type Driver struct {
	config     *rcconfig.ArtBot
	HttpClient httphelper.Helper
	Logger     contract2.LoggerInterface

	GetMiddlewareOfLog func(logger contract2.LoggerInterface) dataflow.RequestMiddleware
}

func NewDriver(config *rcconfig.ArtBot) *Driver {

	HttpClient, _ := httphelper.NewRequestHelper(&httphelper.Config{
		BaseUrl: config.StableDiffusion.BaseUrl,
		Config: &client.Config{
			Timeout: time.Duration(config.StableDiffusion.Timeout) * time.Second,
		},
	})

	log, _ := logger.NewLogger(nil, config.Log)

	driver := &Driver{
		config:     config,
		HttpClient: HttpClient,
		Logger:     log,
	}
	driver.OverrideGetMiddlewares()
	driver.RegisterHttpMiddlewares()

	return driver
}

func (d *Driver) RegisterHttpMiddlewares() {

	// log
	logMiddleware := d.GetMiddlewareOfLog

	config := d.GetConfig()

	d.HttpClient.WithMiddleware(
		logMiddleware(d.Logger),
		httphelper.HttpDebugMiddleware(config.Log.HttpDebug),
	)
}

func (d *Driver) OverrideGetMiddlewares() {
	d.OverrideGetMiddlewareOfLog()
}

func (d *Driver) OverrideGetMiddlewareOfLog() {
	d.GetMiddlewareOfLog = func(logger contract2.LoggerInterface) dataflow.RequestMiddleware {
		return dataflow.RequestMiddleware(func(handle dataflow.RequestHandle) dataflow.RequestHandle {
			return func(request *http.Request, response *http.Response) (err error) {

				// 前置中间件
				if d.config.Log.HttpDebug {
					request2.LogRequest(logger, request)
				}

				err = handle(request, response)
				if err != nil {
					return err
				}

				// 后置中间件
				if d.config.Log.HttpDebug {
					response2.LogResponse(logger, response)
				}

				return
			}
		})
	}
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

	// 获取请求地址
	requestUrl, err := d.GetUrlFromEndpoint(endpoint)
	if err != nil {
		return nil, err
	}

	// 请求数据
	res, err := d.HttpClient.Df().WithContext(ctx).
		Url(requestUrl).
		Method("POST").
		Json(message.Content).
		Request()
	if err != nil {
		return nil, err
	}
	msg := model.CopyMessage(message)

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

func (d *Driver) GetSamplers(ctx context.Context) ([]*model2.Sampler, error) {
	res, err := d.Query(ctx, "/sdapi/v1/samplers")
	if err != nil {
		return nil, err
	}

	reply := []*model2.Sampler{}
	//fmt.Dump(string(res.Content))
	err = json.Unmarshal(res.Content, &reply)

	return reply, err
}

func (d *Driver) GetLoras(ctx context.Context) ([]*model.Lora, error) {
	res, err := d.Query(ctx, "/sdapi/v1/loras")
	if err != nil {
		return nil, err
	}

	reply := []*model.Lora{}
	//fmt.Dump(string(res.Content))
	err = json.Unmarshal(res.Content, &reply)

	return reply, err
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

func (d *Driver) GetControlNetModelList(ctx context.Context) (*controlNet.ControlNetModels, error) {

	res, err := d.Query(ctx, "/controlnet/model_list")
	if err != nil {
		return nil, err
	}

	reply := &controlNet.ControlNetModels{}
	//fmt.Dump(string(res.Content))
	err = json.Unmarshal(res.Content, &reply)

	return reply, err
}
func (d *Driver) GetControlNetModuleList(ctx context.Context) (*controlNet.Modules, error) {
	res, err := d.Query(ctx, "/controlnet/module_list")
	if err != nil {
		return nil, err
	}

	reply := &controlNet.Modules{}
	//fmt.Dump(string(res.Content))
	err = json.Unmarshal(res.Content, &reply)

	return reply, err
}

func (d *Driver) GetControlNetControlTypesList(ctx context.Context) (*controlNet.ControlNetTypes, error) {
	res, err := d.Query(ctx, "/controlnet/control_types")
	if err != nil {
		return nil, err
	}

	reply := &controlNet.ControlNetTypes{}
	//fmt.Dump(string(res.Content))
	err = json.Unmarshal(res.Content, &reply)

	return reply, err
}

func (d *Driver) GetControlNetVersion(ctx context.Context) (*controlNet.ControlNetVersion, error) {
	res, err := d.Query(ctx, "/controlnet/version")
	if err != nil {
		return nil, err
	}

	reply := &controlNet.ControlNetVersion{}
	//fmt.Dump(string(res.Content))
	err = json.Unmarshal(res.Content, &reply)

	return reply, err
}
func (d *Driver) GetControlNetSettings(ctx context.Context) (*controlNet.ControlNetSettings, error) {
	res, err := d.Query(ctx, "/controlnet/settings")
	if err != nil {
		return nil, err
	}

	reply := &controlNet.ControlNetSettings{}
	//fmt.Dump(string(res.Content))
	err = json.Unmarshal(res.Content, &reply)

	return reply, err
}
func (d *Driver) DetectControlNet(ctx context.Context, info *controlNet.DetectInfo) (interface{}, error) {
	content, err := json.Marshal(info)
	if err != nil {
		return nil, err
	}
	reqMessage := model.NewMessage(model.TextMessage)
	reqMessage.Content = content

	return d.Send(ctx, "/controlnet/detect", reqMessage)
}
