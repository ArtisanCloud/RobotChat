package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/ArtisanCloud/RobotChat/pkg"
	"github.com/ArtisanCloud/RobotChat/pkg/objectx"
	fmt2 "github.com/ArtisanCloud/RobotChat/pkg/printx"
	"github.com/ArtisanCloud/RobotChat/rcconfig"
	"github.com/ArtisanCloud/RobotChat/robots/chatBot"
	"github.com/ArtisanCloud/RobotChat/robots/chatBot/driver/ArtisanCloud/chatGLM"
	"github.com/ArtisanCloud/RobotChat/robots/chatBot/driver/contract"
	go_openai "github.com/ArtisanCloud/RobotChat/robots/chatBot/driver/go-openai"
	model2 "github.com/ArtisanCloud/RobotChat/robots/chatBot/model"
	"github.com/ArtisanCloud/RobotChat/robots/kernel/controller"
	"github.com/ArtisanCloud/RobotChat/robots/kernel/model"
	"github.com/artisancloud/httphelper"
)

// Robot Joy is For ChatGPT
var Joy *chatBotService

type chatBotService struct {
	chatBot             *chatBot.ChatBot
	config              *rcconfig.RCConfig
	conversationManager *controller.ConversationManager
}

func NewChatBotService(config *rcconfig.RCConfig) (abs *chatBotService) {

	var driver contract.ChatBotClientInterface
	configChannel := pkg.Lower(config.ChatBot.Channel)
	if configChannel == "" || configChannel == "thudm_glm" {
		// 使用 ArtisanCloud SDK 作为 THUDM_GLM SDK驱动
		driver = chatGLM.NewDriver(&config.ChatBot)
	} else if configChannel == "openai" {
		// 使用 Go-OpenAI 作为 ChatGPT SDK驱动
		driver = go_openai.NewDriver(&config.ChatBot)
	}
	if driver == nil {
		return nil
	}

	robot, err := chatBot.NewChatBot(driver)
	if err != nil {
		panic(err)
	}
	robot.NotifyUrl = config.ChatBot.Queue.NotifyUrl

	abs = &chatBotService{
		chatBot: robot,
		config:  config,
	}
	return abs
}

func (srv *chatBotService) IsAwaken(ctx context.Context) error {
	err := srv.chatBot.IsAwaken(ctx)
	return err
}

func (srv *chatBotService) Launch(ctx context.Context) error {

	// 预处理请求消息
	preProcess := func(ctx context.Context, job *model.Job) (*model.Job, error) {
		srv.chatBot.Logger.Info(srv.chatBot.Name, fmt.Sprintf("I get your message:%s", job.Id))
		return job, nil
	}
	srv.chatBot.SetMessagePreHandler(preProcess)

	// 错误请求处理
	errHandle := func(errReply *model.ErrReply) {
		srv.chatBot.Logger.Error("handle error:", errReply.Job.Id, errReply.Err.Error())
		if errReply.Err != nil {
			errReply.Job.Payload.Metadata.ErrMsg = errReply.Err.Error()
		} else {
			errReply.Job.Payload.Metadata.ErrMsg = "unknown Error from error handle"
		}

		httpClient, err := httphelper.NewRequestHelper(&httphelper.Config{
			BaseUrl: srv.chatBot.NotifyUrl,
		})
		if err != nil {
			srv.chatBot.Logger.Error(srv.chatBot.Name, "handle error new client:", err.Error())
			return
		}

		srv.chatBot.Logger.Error(srv.chatBot.Name, "handle error post notify url:", srv.chatBot.NotifyUrl)

		_, err = httpClient.Df().WithContext(ctx).Method("POST").Json(errReply.Job).Request()
		if err != nil {
			srv.chatBot.Logger.Error(srv.chatBot.Name, "handle error request webhook error:", err.Error())
			return
		}
		return
	}
	srv.chatBot.SetErrorHandler(errHandle)

	// 队列回调请求
	queuePostJobHandle := func(ctx context.Context, job *model.Job) (*model.Job, error) {
		srv.chatBot.Logger.Info("queue has process your request:", job.Id, job.Payload.Content)
		var (
			err     error
			message *model.Message
		)
		message, err = srv.chatBot.Client.CreateChatCompletion(ctx, job.Payload, model.Role(job.Payload.Author))

		if err != nil {
			return job, err
		}
		job.Payload = message

		return job, nil
	}

	queuePostWebhook := func(ctx context.Context, job *model.Job) (*model.Job, error) {

		httpClient, err := httphelper.NewRequestHelper(&httphelper.Config{
			BaseUrl: srv.chatBot.NotifyUrl,
		})
		if err != nil {
			srv.chatBot.Logger.Error(srv.chatBot.Name, "webhook:", err.Error())
			return job, err
		}

		srv.chatBot.Logger.Info(srv.chatBot.Name, "post url:", srv.chatBot.NotifyUrl)
		_, err = httpClient.Df().WithContext(ctx).Method("POST").Json(job).Request()
		if err != nil {
			srv.chatBot.Logger.Error(srv.chatBot.Name, "webhook:", err.Error())
			return job, err
		}

		return job, nil
	}
	srv.chatBot.SetPostMessageHandler(queuePostJobHandle, queuePostWebhook)

	err := srv.chatBot.Start(ctx)

	return err
}

func (srv *chatBotService) Completion(ctx context.Context, req *model2.CompletionRequest) (res *model2.CompletionResponse, err error) {

	res = &model2.CompletionResponse{
		Choices: []model2.CompletionChoice{},
	}

	// 创建消息
	message, err := srv.chatBot.CreateMessage(model.TextMessage, req)
	if err != nil {
		return nil, err
	}

	// 请求数据
	resMes, err := srv.chatBot.Client.CreateCompletion(ctx, message)
	if err != nil {
		return nil, err
	}

	// 解析数据
	glmReply := &chatGLM.GLMResponse{}
	err = objectx.TransformData(resMes.Content, glmReply)
	if err != nil {
		return nil, err
	}

	if glmReply.Status != 200 {
		res.Detail = resMes.Content.String()
		res.Error = "服务器返回错误信息"
		return res, errors.New(res.Error)
	}

	// 返回数据
	res.Choices = []model2.CompletionChoice{
		{
			Text: glmReply.Response,
		},
	}

	return res, err
}

func (srv *chatBotService) CompletionStream(ctx context.Context, req *model2.ChatCompletionRequest) (message *model.Message, err error) {
	reqMsg := req.Messages[0]

	message, err = srv.chatBot.CreateMessage(model.TextMessage, reqMsg.Content)
	if err != nil {
		return nil, err
	}

	message, err = srv.chatBot.CreateCompletionStream(ctx, message, reqMsg.Role, func(data string, status model2.ChatStatus) error {
		fmt2.Dump(data)
		return nil
	})

	return message, err
}

func (srv *chatBotService) ChatCompletion(ctx context.Context, req *model2.ChatCompletionRequest) (message *model.Message, err error) {

	reqMsg := req.Messages[0]

	message, err = srv.chatBot.CreateMessage(model.TextMessage, reqMsg.Content)
	if err != nil {
		return nil, err
	}
	message, err = srv.chatBot.CreateChatCompletion(ctx, message, reqMsg.Role)

	return message, err
}

func (srv *chatBotService) ChatCompletionStream(ctx context.Context, req *model2.ChatCompletionRequest) (message *model.Message, err error) {
	reqMsg := req.Messages[0]

	message, err = srv.chatBot.CreateMessage(model.TextMessage, reqMsg.Content)
	if err != nil {
		return nil, err
	}

	message, err = srv.chatBot.CreateChatCompletionStream(ctx, message, reqMsg.Role, func(data string, status model2.ChatStatus) error {
		fmt2.Dump(data)
		return nil
	})

	return message, err
}
