package service

import (
	"context"
	"encoding/json"
	request2 "github.com/ArtisanCloud/RobotChat/app/request"
	"github.com/ArtisanCloud/RobotChat/pkg"
	"github.com/ArtisanCloud/RobotChat/rcconfig"
	"github.com/ArtisanCloud/RobotChat/robots/artBot"
	"github.com/ArtisanCloud/RobotChat/robots/artBot/driver/ArtisanCloud"
	"github.com/ArtisanCloud/RobotChat/robots/artBot/driver/contract"
	"github.com/ArtisanCloud/RobotChat/robots/artBot/request"
	"github.com/ArtisanCloud/RobotChat/robots/artBot/response"
	"github.com/ArtisanCloud/RobotChat/robots/kernel/controller"
	"github.com/ArtisanCloud/RobotChat/robots/kernel/model"
	"github.com/artisancloud/httphelper"
)

var SrvArtBot *ArtBotService

type ArtBotService struct {
	artBot              *artBot.ArtBot
	config              *rcconfig.RCConfig
	conversationManager *controller.ConversationManager
}

func NewArtBotService(config *rcconfig.RCConfig) (abs *ArtBotService) {

	var driver contract.ArtBotClientInterface
	if config.ArtBot.Channel == "" || pkg.Lower(config.ArtBot.Channel) == "stablediffusion" {
		// 使用 Meonako 作为 SD SDK驱动
		//driver = Meonako.NewDriver(&config.ArtBot)
		driver = ArtisanCloud.NewDriver(&config.ArtBot)
	}
	if driver == nil {
		return nil
	}

	robot, err := artBot.NewArtBot(driver)
	if err != nil {
		panic(err)
	}
	robot.NotifyUrl = config.ArtBot.Queue.NotifyUrl

	abs = &ArtBotService{
		artBot: robot,
		config: config,
	}
	return abs
}

func (srv *ArtBotService) IsAwaken(ctx context.Context) error {
	err := srv.artBot.IsAwaken(ctx)
	return err
}

func (srv *ArtBotService) Launch(ctx context.Context) error {

	// 预处理请求消息
	preProcess := func(ctx context.Context, message *model.Message) (*model.Message, error) {
		srv.artBot.Logger.Info("I get your message:", srv.artBot.Name, message.Content.String())
		return message, nil
	}
	srv.artBot.SetPreMessageHandler(preProcess)

	// 错误请求处理
	errHandle := func(errReply *model.ErrReply) {
		srv.artBot.Logger.Error("handle error:", errReply.Job.Id, errReply.Err.Error())
	}
	srv.artBot.SetErrorHandler(errHandle)

	// 队列回调请求
	queuePostJobHandle := func(ctx context.Context, job *model.Job) (*model.Job, error) {

		srv.artBot.Logger.Info("queue has process your request:", job.Id, job.Payload.Content)
		message, err := srv.artBot.Client.Text2Image(ctx, job.Payload)
		if err != nil {
			return job, err
		}
		job.Payload = message

		return job, nil
	}

	queuePostWebhook := func(ctx context.Context, job *model.Job) (*model.Job, error) {

		httpClient, err := httphelper.NewRequestHelper(&httphelper.Config{
			BaseUrl: srv.artBot.NotifyUrl,
		})
		if err != nil {
			srv.artBot.Logger.Error(srv.artBot.Name, "webhook:", err.Error())
			return job, err
		}

		_, err = httpClient.Df().WithContext(ctx).Method("POST").Json(job).Request()
		if err != nil {
			srv.artBot.Logger.Error(srv.artBot.Name, "webhook:", err.Error())
			return job, err
		}

		return job, nil
	}

	srv.artBot.SetPostMessageHandler(queuePostJobHandle, queuePostWebhook)

	// 启动机器人
	err := srv.artBot.Start(ctx)

	return err
}

func (srv *ArtBotService) Txt2Image(ctx context.Context, req *request.Text2Image) (res *response.Text2Image, err error) {

	strReq, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	message := model.NewMessage(model.TextMessage)
	message.Content = strReq
	res, err = srv.artBot.SendAndWait(ctx, message)

	return res, err
}

func (srv *ArtBotService) ChatTxt2Image(ctx context.Context, req *request2.ParaText2Image) (err error) {

	//conversation := srv.conversationManager.GetConversationByID(req.ConversationId)
	//conversation.GetSessionById[req.SessionId]
	strReq, err := json.Marshal(req)
	if err != nil {
		return err
	}
	message := model.NewMessage(model.TextMessage)
	message.Content = strReq

	_, err = srv.artBot.Send(ctx, message)

	return err
}

func (srv *ArtBotService) WebhookText(ctx context.Context, notify *request2.ParaQueueNotify) {
	srv.artBot.Logger.Info("test notify:", "info key", notify)
}
