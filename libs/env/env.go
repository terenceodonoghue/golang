package env

import "os"

func GetOrDefault(key string, defaultValue interface{}) string {
	var val string
	var ok bool

	if val, ok = os.LookupEnv(key); !ok {
		val = defaultValue.(string)
	}

	return val
}
