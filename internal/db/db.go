package db

import (
	"github.com/laurel-public-schools/lps-api/internal/env"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func New() *gorm.DB {
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: env.GetEnv().DSN,
	}))
	if err != nil {
		panic(err)
	}
	return db
}
