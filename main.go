//go:generate swagger generate spec -o swagger.json

package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/MarceloPetrucio/go-scalar-api-reference"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/laurel-public-schools/lps-api/docs"
	"github.com/laurel-public-schools/lps-api/routes"
)

//	@title			LPS API
//	@version		1.0
//	@description	This is the API for the Laurel Public Schools

//	@host	api.laurel.k12.mt.us
//	@BasePath

//	@schemes	http https
//	@produce	application/json
//	@consumes	application/json

// @securityDefinitions.api_key	ApiKey
// @in								header
// @name							x-api-key
// @Routes /email [post]
// @Routes /ic [get]
// @Routes /docs [get]
func main() {
	docs.SwaggerInfo.Title = "LPS API"
	server := &http.Server{Addr: "0.0.0.0:6969", Handler: service()}

	serverCtx, serverStopCtx := context.WithCancel(context.Background())

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	go func() {
		<-sig
		shutdownCtx, _ := context.WithTimeout(serverCtx, 30*time.Second)

		go func() {
			<-shutdownCtx.Done()
			if shutdownCtx.Err() == context.DeadlineExceeded {
				log.Fatal("graceful shutdown timed out...")
			}
		}()

		err := server.Shutdown(shutdownCtx)
		if err != nil {
			log.Fatal(err)
		}
		serverStopCtx()
	}()

	err := server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		log.Fatal(err)
	}

	<-serverCtx.Done()
}

func service() http.Handler {
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
	println("Server is running on port 6969")
	r.Mount("/email", routes.EmailRequest{}.Routes())
	r.Mount("/ic", routes.ICRoute{}.Routes())
	return r
}
