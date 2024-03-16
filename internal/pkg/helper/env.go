package helper

import (
	"os"
	"strconv"
	"time"
)

type env string

func GetEnvWithDefault(key string, def string) env {
	val := os.Getenv(key)
	if val != "" {
		return env(val)
	}

	return env(def)
}

func (e env) ToString() string {
	return string(e)
}

func (e env) ToDuration() time.Duration {
	t, _ := time.ParseDuration(e.ToString())
	return t
}

func (e env) ToInt() int {
	val, _ := strconv.Atoi(e.ToString())
	return val
}
