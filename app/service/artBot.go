package service

import (
	"context"
	request2 "github.com/ArtisanCloud/RobotChat/app/request"
	"github.com/ArtisanCloud/RobotChat/rcconfig"
	"github.com/ArtisanCloud/RobotChat/robots/artBot"
	"github.com/ArtisanCloud/RobotChat/robots/artBot/driver/Meonako"
	"github.com/ArtisanCloud/RobotChat/robots/artBot/request"
	"github.com/ArtisanCloud/RobotChat/robots/artBot/response"
	"github.com/ArtisanCloud/RobotChat/robots/kernel/controller"
	"github.com/ArtisanCloud/RobotChat/robots/kernel/model"
)

var SrvArtBot *ArtBotService

type ArtBotService struct {
	artBot              *artBot.ArtBot
	config              *rcconfig.RCConfig
	conversationManager *controller.ConversationManager
}

func NewArtBotService(config *rcconfig.RCConfig) (abs *ArtBotService) {

	// 使用 Meonako 作为 SDK驱动
	driver := Meonako.NewDriver(&config.ArtBot)
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
	err := srv.artBot.Start(ctx)
	return err
}

func (srv *ArtBotService) Txt2Image(ctx context.Context, req *request.Text2Image) (res *response.Text2Image, err error) {

	//res, err = srv.artBot.Client.Send()

	return res, err
}

func (srv *ArtBotService) ChatTxt2Image(ctx context.Context, req *request2.ParaText2Image) (err error) {

	//conversation := srv.conversationManager.GetConversationByID(req.ConversationId)
	//conversation.GetSessionById[req.SessionId]
	msg := model.NewMessage(model.TextMessage)
	_, err = srv.artBot.Send(ctx, msg)

	return err
}
