package app

import (
	"time"

	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type App struct {
	Router *gin.Engine
	DB     *gorm.DB
	LOG    *zap.Logger
}

func (app *App) Run() {

	app.initLogger()

	app.Router = gin.New()

	app.Router.Use(ginzap.Ginzap(app.LOG, time.RFC3339, true))

	app.Router.Use(ginzap.RecoveryWithZap(app.LOG, true))

	app.InitDB()

	app.MigrateDB()

	app.initRoutes()

	app.Router.Run()

}

func (app *App) initLogger() {
	logger, _ := zap.NewDevelopment()
	app.LOG = logger
}
