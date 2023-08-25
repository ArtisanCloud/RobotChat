package ArtisanCloud

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/ArtisanCloud/RobotChat/rcconfig"
	"github.com/ArtisanCloud/RobotChat/robots/kernel/logger"
	contract2 "github.com/ArtisanCloud/RobotChat/robots/kernel/logger/contract"
	"github.com/ArtisanCloud/RobotChat/robots/kernel/model"
	request2 "github.com/ArtisanCloud/RobotChat/robots/kernel/request"
	response2 "github.com/ArtisanCloud/RobotChat/robots/kernel/response"
	"github.com/artisancloud/httphelper"
	"github.com/artisancloud/httphelper/dataflow"
	"gorm.io/datatypes"
	"io"
	"net/http"
)

type BaseDriver struct {
	HttpClient httphelper.Helper
	Config     *rcconfig.ChatBot
	Logger     contract2.LoggerInterface

	GetMiddlewareOfLog func(logger contract2.LoggerInterface) dataflow.RequestMiddleware
	GetUrlFromEndpoint func(endpoint string) (string, error)
}

func NewDriver(config *rcconfig.ChatBot) *BaseDriver {
	log, _ := logger.NewLogger(nil, config.Log)
	return &BaseDriver{
		Logger: log,
	}
}

func (d *BaseDriver) OverrideGetMiddlewares() {
	d.OverrideGetMiddlewareOfLog()
}

func (d *BaseDriver) RegisterHttpMiddlewares() {

	// log
	logMiddleware := d.GetMiddlewareOfLog

	config := d.GetConfig()

	d.HttpClient.WithMiddleware(
		logMiddleware(d.Logger),
		httphelper.HttpDebugMiddleware(config.Log.HttpDebug),
	)
}

func (d *BaseDriver) OverrideGetMiddlewareOfLog() {
	d.GetMiddlewareOfLog = func(logger contract2.LoggerInterface) dataflow.RequestMiddleware {
		return dataflow.RequestMiddleware(func(handle dataflow.RequestHandle) dataflow.RequestHandle {
			return func(request *http.Request, response *http.Response) (err error) {

				// 前置中间件
				if d.Config.Log.HttpDebug {
					request2.LogRequest(logger, request)
				}

				err = handle(request, response)
				if err != nil {
					return err
				}

				// 后置中间件
				if d.Config.Log.HttpDebug {
					response2.LogResponse(logger, response)
				}

				return
			}
		})
	}
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
		return message, err
	}

	msg := model.NewMessage(model.TextMessage)

	// 转化返回的Body
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return message, err
	}
	var bodyData datatypes.JSON
	err = json.Unmarshal(body, &bodyData)
	if err != nil {
		return message, err
	}
	msg.Content = bodyData

	// 转化返回的Header
	headerJSON, err := json.Marshal(res.Header)
	if err != nil {
		return message, err
	}

	var headerDataJSON datatypes.JSON
	err = json.Unmarshal(headerJSON, &headerDataJSON)
	if err != nil {
		return message, err
	}

	// 如果返回的接口不是200状态位
	if res.StatusCode != http.StatusOK {
		return message, errors.New(string(msg.Content))
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
