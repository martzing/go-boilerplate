package utilityHelper

import (
	"strconv"
)

func StrToInt(str string, opts ...int) int {
  value, err := strconv.Atoi(str)

  if err != nil {
    defaultValue := 0

    for _, opt := range opts {
      defaultValue = opt
    }
    return defaultValue
  }

  return value
}

func StrToInt64(str string, opts ...int64) int64 {
  value, err := strconv.ParseInt(str, 10, 64)

  if err != nil {
    var defaultValue int64 = 0

    for _, opt := range opts {
      defaultValue = opt
    }
    return defaultValue
  }

  return value
}
