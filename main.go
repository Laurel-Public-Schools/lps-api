//go:generate swagger generate spec -o swagger.json
package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/MarceloPetrucio/go-scalar-api-reference"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/laurel-public-schools/lps-api/env"
	"github.com/laurel-public-schools/lps-api/routes"
)

// @title LPS API
// @version 1.0
// @description This is the API for the Laurel Public Schools

// @host localhost:6969
// @BasePath /

// @schemes http https
// @produce application/json
// @consumes application/json

// @securityDefinitions.api_key ApiKey
// @in header
// @name x-api-key
func main() {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.RealIP)
	r.Use(authCtx)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("🏳️‍🌈🏳️‍🌈🏳️‍🌈"))
	})
	r.Get("/docs", func(w http.ResponseWriter, r *http.Request) {
		htmlContent, err := scalar.ApiReferenceHTML(&scalar.Options{
			SpecURL: "./docs/swagger.json",
			CustomOptions: scalar.CustomOptions{
				PageTitle: "LPS Api",
			},
			DarkMode: true,
		})
		if err != nil {
			fmt.Printf("%v", err)
		}
		fmt.Fprintln(w, htmlContent)
	})

	r.Mount("/email", routes.EmailRequest{}.Routes())
	http.ListenAndServe(":6969", r)
}

func authCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		e := env.GetConfig()
		apiKey := r.Header.Get("x-api-key")
		if apiKey != e.APIKey {
			http.Error(w, "Invalid API Key", http.StatusUnauthorized)
			return
		}
		ctx := context.WithValue(r.Context(), "authorized", true)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
