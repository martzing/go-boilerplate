package queue

import (
	"context"

	"github.com/martzing/go-boilerplate/configs"
	redisHelper "github.com/martzing/go-boilerplate/helpers/redis"
	"github.com/vmihailenco/taskq/v3"
	"github.com/vmihailenco/taskq/v3/redisq"
)

type QueueTask struct {
	Name    string
	Handler func(*taskq.Message)
}

type QueueManage struct {
	queue   taskq.Queue
	context context.Context
}

var Manager *QueueManage

func New() *QueueManage {
	redis := redisHelper.GetRedisInstance()
	redisClient := redis.GetClient()
	QueueFactory := redisq.NewFactory()

	queue := QueueFactory.RegisterQueue(&taskq.QueueOptions{
		Name:  *configs.ServiceName,
		Redis: redisClient,
	})

	ctx := context.Background()
	err := queue.Consumer().Start(ctx)

	if err != nil {
		panic(err)
	}

	Manager = &QueueManage{
		queue:   queue,
		context: ctx,
	}

	return Manager
}

func (qm *QueueManage) CreateQueue(task QueueTask) {
	taskq.RegisterTask(&taskq.TaskOptions{
		Name:    task.Name,
		Handler: task.Handler,
	})
}

func (qm *QueueManage) AddTask(taskName string, payload []byte) {
	err := qm.queue.Add(&taskq.Message{
		TaskName: taskName,
		Ctx:      qm.context,
		ArgsBin:  payload,
	})

	if err != nil {
		panic(err)
	}
}
