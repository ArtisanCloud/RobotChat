package artBot

import (
	"context"
	"errors"
	"github.com/ArtisanCloud/RobotChat/pkg/objectx"
	"github.com/ArtisanCloud/RobotChat/robots"
	"github.com/ArtisanCloud/RobotChat/robots/artBot/driver/contract"
	"github.com/ArtisanCloud/RobotChat/robots/artBot/response"
	"github.com/ArtisanCloud/RobotChat/robots/kernel/logger"
	"github.com/ArtisanCloud/RobotChat/robots/kernel/model"
	queue2 "github.com/ArtisanCloud/RobotChat/robots/kernel/queue"
)

type ArtBot struct {
	*robots.Robot
	Client contract.ArtBotClientInterface
}

func NewArtBot(client contract.ArtBotClientInterface) (*ArtBot, error) {

	conf := client.GetConfig()
	if conf == nil {
		return nil, errors.New("config file is nil")
	}

	// 初始化机器人
	robot, err := robots.NewRobot(&model.RobotAttributes{
		Name:    "Michelle",
		Version: "1.0",
		Gender:  model.GenderFemale,
		Type:    model.TypeArtBot,
	})

	if err != nil {
		return nil, err
	}

	// 按照需求，加载队列驱动，默认是Redis
	q, err := queue2.LoadQueueDriver(&conf.Queue)
	if err != nil {
		return nil, err
	}
	robot.NotifyUrl = conf.NotifyUrl

	// 测试连接队列驱动
	isConnected := q.IsConnected(context.Background())
	if !isConnected {
		return nil, errors.New("cannot connect queue driver")
	}
	robot.Queue = q

	// 初始化Logger
	robot.Logger, err = logger.NewLogger(conf.Log)
	if err != nil {
		return nil, err
	}

	return &ArtBot{
		Robot:  robot,
		Client: client,
	}, nil
}

func (bot *ArtBot) SendAndWait(ctx context.Context, message *model.Message) (*response.Text2Image, error) {

	// 将消息传递给中间件处理
	for _, middleware := range bot.PreMessageHandlers {
		var err error
		// 执行中间件处理逻辑
		message, err = middleware(ctx, message)
		if err != nil {
			// 如果中间件返回错误，可以选择处理错误或直接返回
			return nil, err
		}
	}

	// 请求操作
	msgReply, err := bot.Client.Text2Image(ctx, message)
	if err != nil {
		return nil, err
	}

	// 返回格式
	res := &response.Text2Image{}
	err = objectx.TransformData(msgReply.Content, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}
