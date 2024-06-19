// Description: This helper package is used to load environment variables to immutable struct fields.
package env

import (
	"fmt"
	"os"
	"reflect"

	"github.com/joho/godotenv"
)

// Env struct to hold the environment variables, all fields you intend to use throughout the application should be added here
type Env struct {
	User                    string `env:"EMAIL_USER"`
	Pass                    string `env:"EMAIL_PASSWORD"`
	APIKey                  string `env:"API_KEY"`
	DSN                     string `env:"DSN"`
	REDIS_HOST1             string `env:"REDIS_HOST1"`
	REDIS_HOST2             string `env:"REDIS_HOST2"`
	REDIS_HOST3             string `env:"REDIS_HOST3"`
	REDIS_PORT              string `env:"REDIS_PORT"`
	ONEROSTER_CLIENT_ID     string `env:"ONEROSTER_CLIENT_ID"`
	ONEROSTER_CLIENT_SECRET string `env:"ONEROSTER_CLIENT_SECRET"`
	LHS_SOURCE_ID           string `env:"LHS_SOURCE_ID"`
	XSRF_TOKEN              string `env:"XSRF_TOKEN"`
}

// Description: loadEnv function loads the environment variables from the .env file if the GO_ENV environment variable is set to development
// otherwise it will use the system environment variables
func loadEnv() {
	env := os.Getenv("GO_ENV")
	if env == "development" {
		err := godotenv.Load(".env")

		if err != nil {
			fmt.Println("Error loading .env file")
		}
	}
}

// GetEnv function returns a pointer to the Env struct with the environment variables
func GetEnv() *Env {
	loadEnv()
	var config Env
	val := reflect.ValueOf(&config).Elem()
	typeOfConfig := val.Type()
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		envKey := typeOfConfig.Field(i).Tag.Get("env")
		if envKey == "" {
			continue
		}
		envVal, ok := os.LookupEnv(envKey)
		if !ok {
			panic(fmt.Sprintf("Environment variable %s not set", envKey))
		}
		field.SetString(envVal)
	}
	return &config
}

// GetConfig function returns a pointer to the Env struct with the environment variables
func GetConfig() *Env {
	return GetEnv()
}
