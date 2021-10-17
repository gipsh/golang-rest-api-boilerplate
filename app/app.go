package app

import "github.com/gin-gonic/gin"

type App struct {
	Router *gin.Engine
}

func (app *App) Run() {

	app.Router = gin.New()

	app.Router.Run()

}
