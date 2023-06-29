package queue

import (
	"errors"
	"github.com/ArtisanCloud/RobotChat/pkg"
	"github.com/ArtisanCloud/RobotChat/rcconfig"
	queue2 "github.com/ArtisanCloud/RobotChat/robots/kernel/queue/driver/go-redis"
)

func LoadQueueDriver(config *rcconfig.Queue) (queue QueueInterface, err error) {
	if config == nil {
		return nil, errors.New("queue config is nil")
	}

	// default is redis
	if config.Driver == "" || pkg.Lower(config.Driver) == "redis" {
		queue = queue2.NewRedisQueue(&config.Redis)
		return queue, nil

	} else {
		return nil, errors.New("queue driver not support")
	}
}
