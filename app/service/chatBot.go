package service

import (
	"context"
	"errors"
	"github.com/ArtisanCloud/RobotChat/pkg"
	"github.com/ArtisanCloud/RobotChat/pkg/objectx"
	fmt "github.com/ArtisanCloud/RobotChat/pkg/printx"
	"github.com/ArtisanCloud/RobotChat/rcconfig"
	"github.com/ArtisanCloud/RobotChat/robots/chatBot"
	"github.com/ArtisanCloud/RobotChat/robots/chatBot/driver/ArtisanCloud/chatGLM"
	"github.com/ArtisanCloud/RobotChat/robots/chatBot/driver/contract"
	go_openai "github.com/ArtisanCloud/RobotChat/robots/chatBot/driver/go-openai"
	model2 "github.com/ArtisanCloud/RobotChat/robots/chatBot/model"
	"github.com/ArtisanCloud/RobotChat/robots/kernel/controller"
	"github.com/ArtisanCloud/RobotChat/robots/kernel/model"
	"log"
)

// Robot Joy is For ChatGPT
var Joy *ChatBotService

type ChatBotService struct {
	chatBot             *chatBot.ChatBot
	config              *rcconfig.RCConfig
	conversationManager *controller.ConversationManager
}

func NewChatBotService(config *rcconfig.RCConfig) (abs *ChatBotService) {

	var driver contract.ChatBotClientInterface
	configChannle := pkg.Lower(config.ChatBot.Channel)
	if configChannle == "" || configChannle == "thudm_glm" {
		// 使用 ArtisanCloud SDK 作为 THUDM_GLM SDK驱动
		driver = chatGLM.NewDriver(&config.ChatBot)
	} else if configChannle == "openai" {
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

func (srv *ChatBotService) Completion(ctx context.Context, req *model2.CompletionRequest) (res *model2.CompletionResponse, err error) {

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
		res.Error = "glm服务器返回错误信息"
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

func (srv *ChatBotService) ChatCompletion(ctx context.Context, req *model2.ChatCompletionRequest) (res *model2.ChatCompletionResponse, err error) {

	reqMsg := req.Messages[0]

	message, err := srv.chatBot.CreateMessage(model.TextMessage, reqMsg.Content)
	if err != nil {
		return nil, err
	}

	resMes, err := srv.chatBot.CreateChatCompletion(ctx, message, model.Role(reqMsg.Role))

	return &model2.ChatCompletionResponse{
		Choices: []model2.ChatCompletionChoice{
			{
				Message: model2.ChatCompletionMessage{
					Content: resMes.Content.String(),
				},
			},
		},
	}, err
}

func (srv *ChatBotService) SteamCompletion(ctx context.Context, req *model2.CompletionRequest) (res *model2.CompletionResponse, err error) {

	message, err := srv.chatBot.CreateMessage(model.TextMessage, req.Prompt.(string))
	if err != nil {
		return nil, err
	}

	resMes, err := srv.chatBot.CreateStreamCompletion(ctx, message, model.Role(req.User))

	return &model2.CompletionResponse{
		Choices: []model2.CompletionChoice{
			{
				Text: resMes.Content.String(),
			},
		},
	}, err
}
