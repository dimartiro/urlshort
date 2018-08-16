package app

import (
	"github.com/dimartiro/urlshort/controllers"
)

func mapUrlsToControllers() {
	router.POST("/short", controllers.Short)
	router.GET("/:hash", controllers.RedirectToUrl)
}
