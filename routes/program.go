package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ProgramRoute(router *gin.Engine) {
	router.GET("/program", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"title": "Program LPK Asta",
			"program": []string{
				"Perawat Lansia (Kaigo)",
				"Teknisi Mesin",
				"Konstruksi",
				"Peternakan",
			},
		})
	})
}
