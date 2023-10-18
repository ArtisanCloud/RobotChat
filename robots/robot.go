package robots

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/ArtisanCloud/RobotChat/robots/kernel/logger"
	"github.com/ArtisanCloud/RobotChat/robots/kernel/model"
	queue2 "github.com/ArtisanCloud/RobotChat/robots/kernel/queue"
	"sync"
)

type Robot struct {
	*model.RobotAttributes

	// components
	Queue  queue2.QueueInterface
	Logger *logger.Logger

	// Middlewares
	ErrorHandler        model.HandleError
	PreMessageHandlers  []model.HandlePreSend
	PostMessageHandlers []model.HandlePostReply

	// functional
	IsWorking bool
	Mutex     sync.Mutex
	ErrorChan chan *model.ErrReply

	// webhook
	NotifyUrl string
}

func NewRobot(attributes *model.RobotAttributes) (*Robot, error) {

	// 返回Robot
	return &Robot{
		RobotAttributes: attributes,
		IsWorking:       false,
		Mutex:           sync.Mutex{},
		ErrorChan:       make(chan *model.ErrReply),
	}, nil
}

func (bot *Robot) SetMessagePreHandler(handles ...model.HandlePreSend) {
	bot.PreMessageHandlers = append(bot.PreMessageHandlers, handles...)
}

func (bot *Robot) SetPostMessageHandler(handles ...model.HandlePostReply) {
	bot.PostMessageHandlers = append(bot.PostMessageHandlers, handles...)
}

func (bot *Robot) SetErrorHandler(handle model.HandleError) {
	bot.ErrorHandler = handle
}

func (bot *Robot) IsAwaken(ctx context.Context) error {

	if !bot.Queue.IsConnected(ctx) {
		return errors.New("queue is not connected")
	}

	return nil
}

func (bot *Robot) Start(ctx context.Context) error {

	// 检查是否已经唤醒
	if bot.IsWorking {
		return nil
	}

	// 加锁
	bot.Mutex.Lock()
	defer bot.Mutex.Unlock()

	//
	err := bot.IsAwaken(ctx)
	if err != nil {
		return err
	}

	// 启动消费消息的 Goroutine
	go bot.Receive(ctx)

	// 监听错误通道并处理错误
	go func() {
		for errReply := range bot.ErrorChan {
			// 调用错误处理函数进行处理
			if bot.ErrorHandler != nil {
				bot.ErrorHandler(errReply)
			}
		}
	}()

	// 设置唤醒标志位
	bot.IsWorking = true

	return nil
}

func (bot *Robot) Stop() {
	// 加锁
	bot.Mutex.Lock()
	defer bot.Mutex.Unlock()

	// 关闭 errorChan
	close(bot.ErrorChan)

	// 设置工作状态为停止
	bot.IsWorking = false
}

func (bot *Robot) CreateTextMessage(content interface{}) (*model.Message, error) {
	return bot.CreateMessage(model.TextMessage, content)
}

func (bot *Robot) CreateImageMessage(content interface{}) (*model.Message, error) {
	return bot.CreateMessage(model.ImageMessage, content)
}

func (bot *Robot) CreateMessage(messageType model.MessageType, content interface{}) (*model.Message, error) {
	strContent, err := json.Marshal(content)
	if err != nil {
		return nil, err
	}

	message := model.NewMessage(messageType)
	message.MessageType = messageType
	message.Content = strContent

	message.Metadata = model.MetaData{
		Robot:        bot.RobotAttributes,
		Conversation: nil,
	}

	return message, nil
}

func (bot *Robot) Send(ctx context.Context, jobType int, message *model.Message) (*model.Job, error) {

	job := &model.Job{
		Id:      model.GenerateId(),
		Payload: message,
		Type:    jobType,
	}

	// 将消息传递给中间件处理
	for _, middleware := range bot.PreMessageHandlers {
		var err error
		// 执行中间件处理逻辑
		job, err = middleware(ctx, job)
		if err != nil {
			// 如果中间件返回错误，可以选择处理错误或直接返回
			return nil, err
		}
	}

	err := bot.Queue.ProduceMessage(ctx, job)
	if err != nil {
		return nil, err
	}

	return job, nil
}

func (bot *Robot) Receive(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			// 从队列中获取消息
			job, err := bot.Queue.ConsumeMessage(ctx)
			if err != nil {
				errReply := &model.ErrReply{
					ctx, job, err,
				}
				// 处理获取消息的错误
				bot.ErrorChan <- errReply
				// 可以选择进行错误处理或直接返回
				continue
			}

			// 将消息传递给中间件处理
			for _, middleware := range bot.PostMessageHandlers {
				var err error
				// 执行中间件处理逻辑
				job, err = middleware(ctx, job)
				if err != nil {
					errReply := &model.ErrReply{
						ctx, job, err,
					}
					bot.ErrorChan <- errReply
					// 如果中间件返回错误，可以选择处理错误或直接返回
					//continue
					break
				}
			}
		}
	}
}
