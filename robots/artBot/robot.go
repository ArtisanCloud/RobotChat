package artBot

import (
	"context"
	"errors"
	"github.com/ArtisanCloud/RobotChat/robots/artBot/contract"
	"github.com/ArtisanCloud/RobotChat/robots/kernel/model"
	queue2 "github.com/ArtisanCloud/RobotChat/robots/kernel/queue"
	"sync"
)

type ArtBot struct {
	Client contract.ClientInterface
	Queue  queue2.QueueInterface

	errorHandler        model.HandleError
	preMessageHandlers  []model.HandlePreSend
	postMessageHandlers []model.HandlePostReply
	isWorking           bool
	mutex               sync.Mutex
	errorChan           chan *model.ErrReply
}

func NewArtBot(client contract.ClientInterface) (*ArtBot, error) {
	conf := client.GetConfig()
	if conf == nil {
		return nil, errors.New("config file is nil")
	}

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

	return &ArtBot{
		Client:    client,
		Queue:     q,
		isWorking: false,
		mutex:     sync.Mutex{},
		errorChan: make(chan *model.ErrReply),
	}, nil
}

func (bot *ArtBot) SetPreMessageHandler(handles ...model.HandlePreSend) {
	bot.preMessageHandlers = append(bot.preMessageHandlers, handles...)
}

func (bot *ArtBot) SetPostMessageHandler(handles ...model.HandlePostReply) {
	bot.postMessageHandlers = append(bot.postMessageHandlers, handles...)
}

func (bot *ArtBot) SetErrorHandler(handle model.HandleError) {
	bot.errorHandler = handle
}

func (bot *ArtBot) IsAwaken(ctx context.Context) error {

	if bot.Client == nil {
		return errors.New("robot is not existed")
	}

	if !bot.Queue.IsConnected(ctx) {
		return errors.New("queue is not connected")
	}

	return nil
}

func (bot *ArtBot) Start(ctx context.Context) error {

	// 检查是否已经唤醒
	if bot.isWorking {
		return nil
	}

	// 加锁
	bot.mutex.Lock()
	defer bot.mutex.Unlock()

	//
	err := bot.IsAwaken(ctx)
	if err != nil {
		return err
	}

	// 启动消费消息的 Goroutine
	go bot.Receive(ctx)

	// 监听错误通道并处理错误
	go func() {
		for errReply := range bot.errorChan {
			// 调用错误处理函数进行处理
			if bot.errorHandler != nil {
				bot.errorHandler(errReply)
			}
		}
	}()

	// 设置唤醒标志位
	bot.isWorking = true

	return nil
}

func (bot *ArtBot) Stop() {
	// 加锁
	bot.mutex.Lock()
	defer bot.mutex.Unlock()

	// 关闭 errorChan
	close(bot.errorChan)

	// 设置工作状态为停止
	bot.isWorking = false
}

func (bot *ArtBot) Send(ctx context.Context, message *model.Message) (*model.Job, error) {

	// 将消息传递给中间件处理
	for _, middleware := range bot.preMessageHandlers {
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

func (bot *ArtBot) Receive(ctx context.Context) {
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
				bot.errorChan <- errReply
				// 可以选择进行错误处理或直接返回
				continue
			}

			// 将消息传递给中间件处理
			for _, middleware := range bot.postMessageHandlers {
				var err error
				// 执行中间件处理逻辑
				job, err = middleware(ctx, job)
				if err != nil {
					errReply := &model.ErrReply{
						ctx, job, err,
					}
					bot.errorChan <- errReply
					// 如果中间件返回错误，可以选择处理错误或直接返回
					continue
				}
			}
		}
	}
}
