package schedule

import (
	"github.com/martzing/go-boilerplate/controller/postgres"
	"github.com/martzing/go-boilerplate/job"
)

var ExampleTask = []job.Schedule{
	{
		Id:       "demo-001",
		Schedule: "*/10 * * * * *",
		Body: map[string]interface{}{
			"Data": "Hello",
		},
		Func: postgres.CreateCustomerJobAdapter,
	},
}
