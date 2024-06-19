package middleware

import (
	"context"
	"net/http"

	"github.com/laurel-public-schools/lps-api/env"
)

type Auth struct {
	Authorized bool
}

func AuthCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		e := env.GetConfig()
		apiKey := r.Header.Get("x-api-key")
		if apiKey != e.APIKey {
			http.Error(w, "Invalid API Key", http.StatusUnauthorized)
			return
		}

		type contextKey string

		const authKey contextKey = "auth"

		ctx := context.WithValue(r.Context(), authKey, Auth{Authorized: true})
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
