package env

import (
	"fmt"
	"os"
)

func GetOrDefault(key string, defaultValue interface{}) string {
	val, ok := os.LookupEnv(key)
	if !ok {
		val = fmt.Sprintf("%v", defaultValue)
	}

	return val
}
