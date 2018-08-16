package app

import (
	"os"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func Start() {
	ConfigureRouter()
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	router.Run(":" + port)
}

func ConfigureRouter() {
	router = gin.Default()
	router.RedirectFixedPath = false
	router.RedirectTrailingSlash = true

	mapUrlsToControllers()
}
