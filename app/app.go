package app

import (
	"time"

	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"github.com/gipsh/golang-rest-api-boilerplate/app/config"
	"github.com/morkid/paginate"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type App struct {
	C      config.AppConfig
	Router *gin.Engine
	DB     *gorm.DB
	LOG    *zap.Logger
	P      *paginate.Pagination
}

func (app *App) Run() {

	app.loadConfig()

	app.initLogger()

	app.Router = gin.New()

	app.Router.Use(ginzap.Ginzap(app.LOG, time.RFC3339, true))

	app.Router.Use(ginzap.RecoveryWithZap(app.LOG, true))

	app.InitDB()

	app.MigrateDB()

	app.initPaginate()

	app.Router.Use(ContextInjectionMiddleware(*app))

	app.Router.Use(CORSMiddleware())

	app.initRoutes()

	app.Router.Run()

}

func (app *App) initLogger() {
	logger, _ := zap.NewDevelopment()
	app.LOG = logger
}

func (app *App) initPaginate() {

	app.P = paginate.New(
		&paginate.Config{
			DefaultSize: 50,
		})
}

func (app *App) loadConfig() {

	var err error
	app.C, err = config.LoadConfig(".")
	if err != nil {
		panic("cant load config")
	}
}
