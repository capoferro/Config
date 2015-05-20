package config

import (
	"log"
	"os"
	"strings"
)

var configs = make(map[string]string)
var defaults = make(map[string]string)
var AppName string

func init() {
	if AppName == "" {
		AppName = "app"
	}
}

func Default(key, value string) {
	defaults[key] = value
}

func Get(key string) string {
	if val, exists := configs[key]; exists {
		return val
	} else if val, exists := fromEnv(key); exists {
		return val
	} else if val, exists := defaults[key]; exists {
		return val
	}

	log.Printf("[WARNING] Config \"%s\" does not exist and has no default set.", key)
	return ""
}

func fromEnv(key string) (string, bool) {
	val := os.Getenv(strings.ToUpper(AppName + "_" + key))
	if val == "" {
		return "", false
	}
	configs[key] = val
	return val, true
}
