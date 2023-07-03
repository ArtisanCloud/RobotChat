package chatBot

import (
	"context"
	"errors"
	"github.com/ArtisanCloud/RobotChat/robots"
	"github.com/ArtisanCloud/RobotChat/robots/chatBot/driver/contract"
	"github.com/ArtisanCloud/RobotChat/robots/kernel/logger"
	model2 "github.com/ArtisanCloud/RobotChat/robots/kernel/model"
	queue2 "github.com/ArtisanCloud/RobotChat/robots/kernel/queue"
)

type ChatBot struct {
	*robots.Robot
	Client contract.ChatBotClientInterface
}

func NewChatBot(client contract.ChatBotClientInterface) (*ChatBot, error) {

	conf := client.GetConfig()
	if conf == nil {
		return nil, errors.New("config file is nil")
	}

	// 初始化机器人
	robot, err := robots.NewRobot(&robots.RobotAttributes{
		Name:    "Joy",
		Version: "1.0",
		Gender:  robots.GenderFemale,
		Type:    robots.TypeChatBot,
	})

	// 按照需求，加载队列驱动，默认是Redis
	q, err := queue2.LoadQueueDriver(&conf.Queue)
	if err != nil {
		return nil, err
	}
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

	return &ChatBot{
		Robot:  robot,
		Client: client,
	}, nil
}

// SendMessage 向指定对话发送消息
func (bot *ChatBot) CreateChatCompletion(ctx context.Context, message string, role model2.Role) (string, error) {
	return bot.Client.CreateChatCompletion(ctx, message, role)
}

func (bot *ChatBot) CreateStreamCompletion(ctx context.Context, message string, role model2.Role) (string, error) {
	return bot.Client.CreateStreamCompletion(ctx, message, role)
}

// GenerateAnswer 生成无上下文回答
func (bot *ChatBot) CreateCompletion(ctx context.Context, prompt string) (string, error) {
	return bot.Client.CreateCompletion(ctx, prompt)
}
