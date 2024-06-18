package env

import (
	"os"
	// godotenv "github.com/joho/godotenv"
)

type Env struct {
	User        string `json:"user"`
	Pass        string `json:"pass"`
	APIKey      string `json:"api_key"`
	DSN         string `json:"dsn"`
	REDIS_HOST1 string `json:"redis_host1"`
	REDIS_HOST2 string `json:"redis_host2"`
	REDIS_HOST3 string `json:"redis_host3"`
	REDIS_PORT  string `json:"redis_port"`
}

const (
	User        = "EMAIL_USER"
	Pass        = "EMAIL_PASSWORD"
	APIKey      = "API_KEY"
	DSN         = "DSN"
	REDIS_HOST1 = "REDIS_HOST1"
	REDIS_HOST2 = "REDIS_HOST2"
	REDIS_HOST3 = "REDIS_HOST3"
	REDIS_PORT  = "REDIS_PORT"
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
	if val, ok := os.LookupEnv(DSN); ok {
		config.DSN = val
	}
	if val, ok := os.LookupEnv(REDIS_HOST1); ok {
		config.REDIS_HOST1 = val
	}
	if val, ok := os.LookupEnv(REDIS_HOST2); ok {
		config.REDIS_HOST2 = val
	}
	if val, ok := os.LookupEnv(REDIS_HOST3); ok {
		config.REDIS_HOST3 = val
	}
	if val, ok := os.LookupEnv(REDIS_PORT); ok {
		config.REDIS_PORT = val
	}
	return &config
}

func GetConfig() *Env {
	return GetEnv()
}

// func readEnvFile(filePath string) map[string]string {
// 	envMap, err := godotenv.Read(filePath)
// 	if err != nil {
// 		panic(err)
// 	}
// 	return envMap
// }

// func envToConfig(envMap map[string]string) *Env {
// 	var config Env
// 	if user, ok := envMap[User]; ok {
// 		config.User = user
// 	}
// 	if pass, ok := envMap[Pass]; ok {
// 		config.Pass = pass
// 	}
// 	if apiKey, ok := envMap[APIKey]; ok {
// 		config.APIKey = apiKey
// 	}
// 	return &config
// }

// func GetConfig() *Env {
// 	envMap := readEnvFile(".env")
// 	return envToConfig(envMap)
// }
