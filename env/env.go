package env

import (
	"os"

	godotenv "github.com/joho/godotenv"
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

func readEnvFile(filePath string) (map[string]string, error) {
	envMap, err := godotenv.Read(filePath)
	if err != nil {
		return nil, err
	}
	return envMap, nil
}

func envToConfig(envMap map[string]string) *Env {
	var config Env
	if user, ok := envMap[User]; ok {
		config.User = user
	}
	if pass, ok := envMap[Pass]; ok {
		config.Pass = pass
	}
	if apiKey, ok := envMap[APIKey]; ok {
		config.APIKey = apiKey
	}
	return &config
}
func GetConfig() *Env {
	envMap, err := readEnvFile(".env")
	if err != nil {
		return GetEnv()
	}
	return envToConfig(envMap)
}
