package main

import (
	_ "github.com/laurel-public-schools/lps-api/docs"
	"github.com/laurel-public-schools/lps-api/router"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func main() {
	r := router.New()
}
