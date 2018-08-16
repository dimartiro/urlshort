package app

import (
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func Start() {
	ConfigureRouter()
	router.Run(":8080")
}

func ConfigureRouter() {
	router = gin.Default()
	router.RedirectFixedPath = false
	router.RedirectTrailingSlash = true

	mapUrlsToControllers()
}
