package service

import (
	"context"
	"encoding/json"
	request2 "github.com/ArtisanCloud/RobotChat/app/request"
	"github.com/ArtisanCloud/RobotChat/pkg"
	fmt "github.com/ArtisanCloud/RobotChat/pkg/printx"
	"github.com/ArtisanCloud/RobotChat/rcconfig"
	"github.com/ArtisanCloud/RobotChat/robots/artBot"
	"github.com/ArtisanCloud/RobotChat/robots/artBot/driver/Meonako"
	"github.com/ArtisanCloud/RobotChat/robots/artBot/driver/contract"
	"github.com/ArtisanCloud/RobotChat/robots/artBot/request"
	"github.com/ArtisanCloud/RobotChat/robots/artBot/response"
	"github.com/ArtisanCloud/RobotChat/robots/kernel/controller"
	"github.com/ArtisanCloud/RobotChat/robots/kernel/model"
	"log"
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
		driver = Meonako.NewDriver(&config.ArtBot)
	}
	if driver == nil {
		return nil
	}

	robot, err := artBot.NewArtBot(driver)
	if err != nil {
		panic(err)
	}

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

func (srv *ArtBotService) Start(ctx context.Context) error {
	// 启动机器人
	preProcess := func(ctx context.Context, message *model.Message) (*model.Message, error) {
		fmt.Dump("I get your message:", message.Content.String())
		return message, nil
	}
	queueCallback := func(ctx context.Context, job *model.Job) (*model.Job, error) {
		preload := job.Payload.(map[string]interface{})
		fmt.Dump("queue has process your request:", job.Id, preload["content"])
		return job, nil
	}
	errHandle := func(errReply *model.ErrReply) {
		log.Printf("handle error: %s, %s", errReply.Job.Id, errReply.Err.Error())
	}

	srv.artBot.SetPreMessageHandler(preProcess)
	srv.artBot.SetPostMessageHandler(queueCallback)
	srv.artBot.SetErrorHandler(errHandle)
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
	msg := model.NewMessage(model.TextMessage)
	_, err = srv.artBot.Send(ctx, msg)

	return err
}
