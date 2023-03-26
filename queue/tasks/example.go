package queueTasks

import (
	"github.com/martzing/go-boilerplate/controller/postgres"
	"github.com/martzing/go-boilerplate/queue"
)

var Tasks = []queue.QueueTask{
	{
		Name:    "example",
		Handler: postgres.CreateCustomerQueueAdapter,
	},
}
