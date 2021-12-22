package app

import (
	migrate "github.com/rubenv/sql-migrate"
	"go.uber.org/zap"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"moul.io/zapgorm2"
)

func (app *App) InitDB() {

	logger := zapgorm2.New(app.LOG)
	logger.SetAsDefault()

	var err error

	app.DB, err = gorm.Open(sqlite.Open("db/api.db"), &gorm.Config{
		Logger: logger,
	})

	if err != nil {
		panic(err)
	}

}

func (app *App) MigrateDB() {
	migrations := &migrate.FileMigrationSource{
		Dir: "db/migrations",
	}

	db, err := app.DB.DB()
	if err != nil {
		panic("migration error: cant get db handler")
	}

	n, err := migrate.Exec(db, "postgres", migrations, migrate.Up)
	if err != nil {
		app.LOG.Info("migration error", zap.Error(err))
	}
	app.LOG.Info("Applied migrations!", zap.Int("migrations", n))

}
