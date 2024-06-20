// Description: Utils package for helper functions and structs
package utils

import (
	"context"

	"github.com/laurel-public-schools/lps-api/internal/env"
	"github.com/laurel-public-schools/lps-api/internal/middlewares"
)

type Utils struct{}

func EnvFromContext(ctx context.Context) *env.Env {
	if v := ctx.Value(middlewares.EnvKey); v != nil {
		if envConfig, ok := v.(*env.Env); ok {
			return envConfig
		}
	}
	return nil
}
