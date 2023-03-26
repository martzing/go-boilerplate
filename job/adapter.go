package job

import (
	"encoding/json"
	"fmt"

	utilityHelper "github.com/martzing/go-boilerplate/helpers/utility"
)

func CreateJobAdapter[T comparable](service func(data T)) func(any) {
	logger := utilityHelper.NewLogger()

	adapter := func(body any) {
		logger.Info("start job")

		defer func() {
			if err := recover(); err != nil {
				logger.Errorln(err)
			}
		}()

		var data T

		if body != nil {
			if err := json.Unmarshal([]byte(fmt.Sprintf("%v", body)), &data); err != nil {
				panic(err)
			}
		}

		service(data)

		logger.Info("end job")
	}

	return adapter
}
