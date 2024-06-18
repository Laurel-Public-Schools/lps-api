package db

import (
	"github.com/laurel-public-schools/lps-api/env"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func New() *gorm.DB {
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: env.GetConfig().DSN,
	}))
	if err != nil {
		panic(err)
	}
	return db
}
