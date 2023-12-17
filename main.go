package main

import (
	"github.com/ayushvaish2511/CRUD_GOFR/handler"
	"github.com/ayushvaish2511/CRUD_GOFR/migrations"
	"github.com/ayushvaish2511/CRUD_GOFR/store"
	"gofr.dev/cmd/gofr/migration"
	dbmigration "gofr.dev/cmd/gofr/migration/dbMigration"
	"gofr.dev/pkg/gofr"
)

func main() {
	app := gofr.New()

	s := store.New()
	h := handler.New(s)

	appName := app.Config.Get("APP_NAME")
	
	err := migration.Migrate(appName, dbmigration.NewGorm(app.GORM()), migrations.All(),
		dbmigration.UP, app.Logger)
	if err != nil {
		app.Logger.Error(err)

		return
	}

	app.GET("/customer", h.Get)
	app.GET("/customer/{id}", h.GetById)
	app.POST("/customer", h.Create)
	app.PUT("/customer/{id}", h.Update)
	app.DELETE("/customer/{id}", h.Delete)
	

	app.Server.HTTP.Port = 9090
	app.Server.MetricsPort = 2325
	app.Start()
}
