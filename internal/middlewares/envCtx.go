package middlewares

import (
	"context"
	"net/http"

	"github.com/laurel-public-schools/lps-api/internal/env"
)

type contextKey string

const EnvKey contextKey = "env"

func EnvCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		e := env.GetEnv()
		ctx := context.WithValue(r.Context(), EnvKey, e)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
