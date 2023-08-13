package common

import (
	"fmt"
	"log"
	"os"
)

func GetEnvStrict(env string) string {
	val := os.Getenv(env)
	if len(val) == 0 {
		log.Fatalln(fmt.Sprintf("'%s' not found, please set var: %s", env, env))
	}
	return val
}

func GetEnv(env string, defaultVal string) string {
	val := os.Getenv(env)
	if len(val) == 0 {
		return defaultVal
	}
	return val
}
