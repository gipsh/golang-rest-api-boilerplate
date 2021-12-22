package app

import "github.com/gin-gonic/gin"

// context: inject the deps handlers into controllers
func ContextInjectionMiddleware(app App) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("DB", app.DB)
		c.Set("LOG", app.LOG)
		c.Set("PAGINATE", app.P)
		c.Set("CONFIG", app.C)

		c.Next()
	}
}
