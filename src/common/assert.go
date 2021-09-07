package common

import "fmt"

func MustInt(input interface{}) int {
	if input == nil {
		return 0
	}
	switch input.(type) {
	case float64:
		return int(input.(float64))
	case int:
		return input.(int)

	}
	return 0
}

func MustInt64(input interface{}) int64 {
	if input == nil {
		return 0
	}
	switch input.(type) {
	case float64:
		return int64(input.(float64))
	case int:
		return int64(input.(int))
	case int64:
		return input.(int64)
	}
	return 0
}

func MustString(input interface{}) string {
	if input == nil {
		return ""
	}
	return fmt.Sprintf("%v", input)

}
