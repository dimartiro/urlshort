package controllers

import (
	"net/http"

	"github.com/dimartiro/urlshort/api/src/services"
	"github.com/gin-gonic/gin"
)

func Short(c *gin.Context) {
	url := c.Query("url")

	hashedUrl := services.Short(url)

	c.String(http.StatusOK, hashedUrl)
}
