package app

import (
	"github.com/gin-gonic/gin"
	"github.com/gipsh/golang-rest-api-boilerplate/controllers"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

//
// @title Articles API
// @version 1.0
// @description Articles API server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// host localhost:8080
// @BasePath /api/v1
// @query.collection.format multi

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func (app *App) initRoutes() {

	// controllers
	article := controllers.Article{}

	// api prefix
	apiv1 := app.Router.Group("/api/v1")

	//api version
	apiv1.GET("/version", controllers.GetVersion)

	//	default invalid route
	app.Router.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	// swagger api doc
	//swagger_ip := fmt.Sprintf("http://%s:8080/swagger/doc.json", "localhost")
	//url := ginSwagger.URL(swagger_ip) // The url pointing to API definition
	//app.Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	app.Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// article
	apiv1.GET("/article/:id", article.Get)
	apiv1.GET("/articles", article.List)
	apiv1.POST("/article", article.Create)
	apiv1.PUT("/article/:id", article.Update)
	apiv1.DELETE("/article/:id", article.Delete)
}
