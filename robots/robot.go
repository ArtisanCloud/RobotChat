package robots

import (
	"context"
	"errors"
	"github.com/ArtisanCloud/RobotChat/robots/kernel/logger"
	"github.com/ArtisanCloud/RobotChat/robots/kernel/model"
	queue2 "github.com/ArtisanCloud/RobotChat/robots/kernel/queue"
	"sync"
)

type RobotType int8

const (
	TypeChatBot RobotType = iota
	TypeArtBot
)

type RobotGender int8

const (
	GenderMale RobotGender = iota
	GenderFemale
	GenderNeutral
)

type RobotAttributes struct {
	// attributes
	Name    string
	Version string
	Gender  RobotGender
	Type    RobotType
}

type Robot struct {
	*RobotAttributes

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

func NewRobot(attributes *RobotAttributes) (*Robot, error) {

	// 返回Robot
	return &Robot{
		RobotAttributes: attributes,
		IsWorking:       false,
		Mutex:           sync.Mutex{},
		ErrorChan:       make(chan *model.ErrReply),
	}, nil
}

func (bot *Robot) SetPreMessageHandler(handles ...model.HandlePreSend) {
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

func (bot *Robot) Send(ctx context.Context, message *model.Message) (*model.Job, error) {

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

	job := &model.Job{
		Id:      model.GenerateId(),
		Payload: message,
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
					continue
				}
			}
		}
	}
}
