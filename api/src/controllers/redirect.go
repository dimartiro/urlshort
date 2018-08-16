package controllers

import (
	"net/http"

	"github.com/dimartiro/urlshort/api/src/model"
	"github.com/gin-gonic/gin"
)

func RedirectToUrl(c *gin.Context) {
	hash := c.Param("hash")

	if url := model.GetUrlFromHash(hash); url != "" {
		c.Redirect(http.StatusMovedPermanently, url)
	} else {
		c.Status(http.StatusNotFound)
	}

}
