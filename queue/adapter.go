package queue

import (
	"encoding/json"
	"fmt"

	"github.com/vmihailenco/taskq/v3"
)

func CreateQueueAdapter[T comparable](service func(data T)) func(*taskq.Message) {
	adapter := func(message *taskq.Message) {
		fmt.Println("Start Queue")

		var data T

		json.Unmarshal(message.ArgsBin, &data)

		service(data)

		fmt.Println("End Queue")
	}

	return adapter
}
