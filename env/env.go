package env

import (
	"os"
)

type Env struct {
	User   string `json:"user"`
	Pass   string `json:"pass"`
	APIKey string `json:"api_key"`
}

const (
	User   = "EMAIL_USER"
	Pass   = "EMAIL_PASSWORD"
	APIKey = "API_KEY"
)

func GetEnv() *Env {
	var config Env
	if val, ok := os.LookupEnv(User); ok {
		config.User = val
	}
	if val, ok := os.LookupEnv(Pass); ok {
		config.Pass = val
	}
	if val, ok := os.LookupEnv(APIKey); ok {
		config.APIKey = val
	}
	return &config
}

func GetConfig() *Env {
	return GetEnv()
}
