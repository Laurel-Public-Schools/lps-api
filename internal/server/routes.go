package server

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/laurel-public-schools/lps-api/internal/middlewares"
	"github.com/laurel-public-schools/lps-api/internal/routes"

	"net/http"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.RealIP)
	r.Use(middlewares.AuthCtx)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("🏳️‍🌈🏳️‍🌈🏳️‍🌈"))
	})
	r.Mount("/email", routes.EmailRequest{}.Routes())
	r.Mount("/ic", routes.ICRoute{}.Routes())
	return r
}
