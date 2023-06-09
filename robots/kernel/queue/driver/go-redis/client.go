package queue

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/ArtisanCloud/RobotChat/rcconfig"
	"github.com/ArtisanCloud/RobotChat/robots/kernel/model"
	"github.com/redis/go-redis/v9"
)

type RedisQueue struct {
	Client *redis.Client
}

func NewRedisQueue(config *rcconfig.Redis) (queue *RedisQueue) {
	c := redis.NewClient(&redis.Options{
		Addr:       config.Addr,
		ClientName: config.ClientName,
		Username:   config.Username,
		Password:   config.Password,
		DB:         config.DB,
		MaxRetries: config.MaxRetries,
	})
	//fmt.Dump( c.Options().ClientName)
	queue = &RedisQueue{
		Client: c,
	}

	return queue
}

func (q *RedisQueue) Connect(ctx context.Context) error {
	// 测试连接
	_, err := q.Client.Ping(ctx).Result()
	if err != nil {
		return err
	}

	return nil
}

func (q *RedisQueue) IsConnected(ctx context.Context) bool {
	pong, err := q.Client.Ping(ctx).Result()
	if err != nil {
		return false
	}
	return pong == "PONG"
}

func (q *RedisQueue) ProduceMessage(ctx context.Context, job *model.Job) error {

	if job == nil || job.Payload == nil {
		return errors.New("job object is invalid")
	}

	key := q.GetQueueName()

	payloadBytes, err := json.Marshal(job) // 将有效负载序列化为 JSON 字符串
	if err != nil {
		return err
	}

	_, err = q.Client.LPush(ctx, key, payloadBytes).Result()
	if err != nil {
		return err
	}

	return nil
}

func (q *RedisQueue) ConsumeMessage(ctx context.Context) (*model.Job, error) {
	key := q.GetQueueName()
	result, err := q.Client.BRPop(ctx, 0, key).Result()
	if err != nil {
		return nil, err
	}
	job := &model.Job{}
	err = json.Unmarshal([]byte(result[1]), job)

	return job, err
}

func (q *RedisQueue) QueueLength(ctx context.Context) (int, error) {
	length, err := q.Client.LLen(ctx, "my_queue").Result()
	if err != nil {
		return 0, err
	}
	return int(length), nil
}

func (q *RedisQueue) Close(ctx context.Context) error {
	err := q.Client.Close()
	if err != nil {
		return err
	}

	// 在这里可以执行其他关闭操作或处理特定的上下文取消

	return nil
}

func (q *RedisQueue) GetQueueName() string {
	return q.Client.Options().ClientName
}
