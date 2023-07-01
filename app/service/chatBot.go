package service

import (
	"context"
	"github.com/ArtisanCloud/PowerLibs/v3/object"
	"github.com/ArtisanCloud/RobotChat/pkg"
	fmt "github.com/ArtisanCloud/RobotChat/pkg/printx"
	"github.com/ArtisanCloud/RobotChat/rcconfig"
	"github.com/ArtisanCloud/RobotChat/robots/artBot/response"
	"github.com/ArtisanCloud/RobotChat/robots/chatBot"
	"github.com/ArtisanCloud/RobotChat/robots/chatBot/driver/contract"
	go_openai "github.com/ArtisanCloud/RobotChat/robots/chatBot/driver/go-openai"
	"github.com/ArtisanCloud/RobotChat/robots/chatBot/request"
	"github.com/ArtisanCloud/RobotChat/robots/kernel/controller"
	"github.com/ArtisanCloud/RobotChat/robots/kernel/model"
	"gorm.io/datatypes"
	"log"
)

var SrvChatBot *ChatBotService

type ChatBotService struct {
	chatBot             *chatBot.ChatBot
	config              *rcconfig.RCConfig
	conversationManager *controller.ConversationManager
}

func NewChatBotService(config *rcconfig.RCConfig) (abs *ChatBotService) {

	var driver contract.ChatBotClientInterface
	if config.ChatBot.Channel == "" || pkg.Lower(config.ChatBot.Channel) == "openai" {
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

	abs = &ChatBotService{
		chatBot: robot,
		config:  config,
	}
	return abs
}

func (srv *ChatBotService) IsAwaken(ctx context.Context) error {
	err := srv.chatBot.IsAwaken(ctx)
	return err
}

func (srv *ChatBotService) Launch(ctx context.Context) error {
	// 启动机器人
	preProcess := func(ctx context.Context, message *model.Message) (*model.Message, error) {
		fmt.Dump("I get your message:", message.Content.String())
		return message, nil
	}
	queueCallback := func(ctx context.Context, job *model.Job) (*model.Job, error) {
		fmt.Dump("queue has process your request:", job.Id, job.Payload)
		return job, nil
	}
	errHandle := func(errReply *model.ErrReply) {
		log.Printf("handle error: %s, %s", errReply.Job.Id, errReply.Err.Error())
	}

	srv.chatBot.SetPreMessageHandler(preProcess)
	srv.chatBot.SetPostMessageHandler(queueCallback)
	srv.chatBot.SetErrorHandler(errHandle)

	err := srv.chatBot.Start(ctx)

	return err
}

func (srv *ChatBotService) Completion(ctx context.Context, req *request.ChatCompletionRequest) (res *response.Text2Image, err error) {

	//res, err = srv.chatBot.Client.Send()

	return res, err
}

func (srv *ChatBotService) ChatCompletion(ctx context.Context, req *request.ChatCompletionRequest) (err error) {

	//conversation := srv.conversationManager.GetConversationByID(req.ConversationId)
	//conversation.GetSessionById[req.SessionId]
	msg := model.NewMessage(model.TextMessage)
	strReq, err := object.JsonEncode(req)
	if err != nil {
		return err
	}
	msg.Content = datatypes.JSON(strReq)
	_, err = srv.chatBot.Send(ctx, msg)

	return err
}
