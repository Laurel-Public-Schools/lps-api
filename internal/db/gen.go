package db

import (
	"gorm.io/gen"
)

func Gen() {
	g := gen.NewGenerator(gen.Config{
		OutPath: "./queries",
		Mode:    gen.WithQueryInterface,
	})
	db := New()
	g.UseDB(db)

	g.ApplyBasic(
		g.GenerateAllTable()...,
	)
	g.Execute()
}
