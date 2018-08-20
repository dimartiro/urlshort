package controllers

import (
	"net/http"

	"github.com/dimartiro/urlshort/api/src/services"
	"github.com/gin-gonic/gin"
)

func Short(c *gin.Context) {
	url := c.Query("url")

	hashedUrl, err := services.Short(url)

	if err != nil {
		c.String(http.StatusInternalServerError, "Error generating and saving URL")
		return
	}

	c.String(http.StatusOK, hashedUrl)
}
