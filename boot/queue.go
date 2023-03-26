package boot

import (
	utilityHelper "github.com/martzing/go-boilerplate/helpers/utility"
	"github.com/martzing/go-boilerplate/queue"
	queueTasks "github.com/martzing/go-boilerplate/queue/tasks"
)

func bootQueue() {
	logger := utilityHelper.NewLogger()

	logger.Infoln("Initial Queue")

	q := queue.New()

	for _, task := range queueTasks.Tasks {
		q.CreateQueue(task)
	}
}
