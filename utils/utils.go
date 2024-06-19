// Description: Utils package for helper functions and structs
package utils

import (
	"context"
	"github.com/laurel-public-schools/lps-api/env"
	"github.com/laurel-public-schools/lps-api/middleware"
)

type Utils struct{}

func EnvFromContext(ctx context.Context) *env.Env {
	if v := ctx.Value(middleware.EnvKey); v != nil {
		if envConfig, ok := v.(*env.Env); ok {
			return envConfig
		}
	}
	return nil
}
