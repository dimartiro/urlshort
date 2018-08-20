package controllers

import (
	"fmt"
	"net/http"

	"github.com/dimartiro/urlshort/api/src/dao"
	"github.com/gin-gonic/gin"
)

func RedirectToUrl(c *gin.Context) {
	hash := c.Param("hash")

	url, err := dao.GetUrlFromHashFromDB(hash)

	if err != nil {
		fmt.Println("Error getting URL " + err.Error())
		return
	}

	if url != "" {
		c.Redirect(http.StatusMovedPermanently, url)
	} else {
		c.Status(http.StatusNotFound)
	}

}
