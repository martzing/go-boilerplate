package boot

import (
	"github.com/martzing/go-boilerplate/job"
	"github.com/martzing/go-boilerplate/job/schedule"
)

func bootJob() {
	q := job.New()

	// add task to this array
	tasks := append(
		schedule.ExampleTask,
	)

	for _, task := range tasks {
		q.CreateJob(job.Schedule(task))
	}

	q.Start()
}
