package main

import (
	"fmt"
	"github.com/MarceloPetrucio/go-scalar-api-reference"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/laurel-public-schools/lps-api/routes"
	"net/http"
)

func main() {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.RealIP)
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
