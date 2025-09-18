package routes

import (
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HomeRoute(router *gin.Engine) {
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "layout.html", gin.H{
			"Title": "Home",
			"Content": template.HTML(`
				{{ template "content" . }}
			`),
		})
	})
}
