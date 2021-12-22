package app

import (
	"fmt"

	"github.com/gipsh/golang-rest-api-boilerplate/controllers"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func (app *App) initRoutes() {

	// api prefix
	apiv1 := app.Router.Group("/api/v1")

	//api version
	apiv1.GET("/version", controllers.GetVersion)

	// default invalid route
	// app.Router.NoRoute(app.MiddlewareFunc(), func(c *gin.Context) {
	// 	c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	// })

	// swagger api doc
	swagger_ip := fmt.Sprintf("http://%s:8080/swagger/doc.json", "localhost")
	url := ginSwagger.URL(swagger_ip) // The url pointing to API definition
	app.Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

}
