package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AboutRoute(router *gin.Engine) {
	router.GET("/about", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"title":   "Tentang Kami",
			"content": "kami adalah LPK yang membantu tenaga kerja Indonesia di Jepang",
		})
	})
}
