package main

import (
	_ "github.com/laurel-public-schools/lps-api/docs"
	"github.com/laurel-public-schools/lps-api/handler"
	"github.com/laurel-public-schools/lps-api/router"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title LPS API
// @version 1.0
// @description This is the API for the Laurel Public Schools

// @host localhost:6969
// @BasePath /api

// @schemes http https
// @produce application/json
// @consumes application/json

// @securityDefinitions.api_key ApiKey
// @in header
// @name Authorization

func main() {
	r := router.New()
	r.GET("/swagger/*", echoSwagger.WrapHandler)

	v1 := r.Group("/api")

	h := handler.NewHandler()
	h.Register(v1)
	r.Logger.Fatal(r.Start(":6969"))
}
