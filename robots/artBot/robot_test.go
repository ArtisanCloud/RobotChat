package artBot

import (
	"context"
	fmt "github.com/ArtisanCloud/RobotChat/pkg/printx"
	"github.com/ArtisanCloud/RobotChat/rcconfig"
	"github.com/ArtisanCloud/RobotChat/robots/artBot/driver/Meonako"
	"github.com/ArtisanCloud/RobotChat/robots/kernel/model"
	"github.com/stretchr/testify/assert"
	"gorm.io/datatypes"
	"log"
	"os"
	"path/filepath"
	"testing"
	"time"
)

func TestQueue_Working(t *testing.T) {
	exePath, err := os.Getwd()
	assert.NoError(t, err)
	//fmt.Dump(exePath)
	configDir := filepath.Dir(exePath)
	configPath := filepath.Join(configDir, "../config.yml")
	config := rcconfig.LoadRCConfigByPath(configPath)

	// 创建一个假的队列驱动
	driver := Meonako.NewDriver(&config.ArtBot)
	bot, err := NewArtBot(driver)

	// 启动机器人
	ctx := context.Background()
	preProcess := func(ctx context.Context, job *model.Job) (*model.Job, error) {
		fmt.Dump("I get your message:", job.Payload.Content.String())
		return job, nil
	}
	queueCallback := func(ctx context.Context, job *model.Job) (*model.Job, error) {
		preload := job.Payload
		fmt.Dump("queue has process your request:", job.Id, preload.Content)
		return job, nil
	}
	errHandle := func(errReply *model.ErrReply) {
		log.Printf("handle error: %s, %s", errReply.Job.Id, errReply.Err.Error())
	}

	bot.SetMessagePreHandler(preProcess)
	bot.SetPostMessageHandler(queueCallback)
	bot.SetErrorHandler(errHandle)
	err = bot.Start(ctx)
	assert.NoError(t, err)

	// 创建一个测试消息
	message := model.NewMessage(model.TextMessage)

	// 发送消息到队列
	message.Content = datatypes.JSON(`{"prompt": "Are you a Robot?"}`)
	_, err = bot.Send(ctx, message)
	assert.NoError(t, err)

	message.Content = datatypes.JSON(`{"prompt": "how old are you?"}`)
	_, err = bot.Send(ctx, message)
	assert.NoError(t, err)

	message.Content = datatypes.JSON(`{"prompt": "你会说中文么?"}`)
	_, err = bot.Send(ctx, message)
	assert.NoError(t, err)

	// 延迟一段时间，等待机器人处理消息
	time.Sleep(time.Second)

	// 停止机器人
	bot.Stop()

	// 断言机器人的工作状态已经停止
	assert.False(t, bot.IsWorking)
}
