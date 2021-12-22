package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/gipsh/golang-rest-api-boilerplate/build"
)

func GetVersion(c *gin.Context) {
	c.JSON(200, gin.H{
		"version": build.Version,
		"build":   build.Build,
		"time":    build.Time,
		"author":  build.User,
	})
}
