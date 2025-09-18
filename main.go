package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"asta-karya/app"
	"asta-karya/controllers"

	"github.com/gin-gonic/gin"
)

type Header struct {
	Title    string `json:"title"`
	Home     string `json:"home"`
	About    string `json:"about"`
	Program  string `json:"program"`
	Training string `json:"training"`
	Media    string `json:"media"`
	Galery   string `json:"galery"`
}

func loadJSONFile(filePath string, target interface{}) error {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, target)
}

func main() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	router.LoadHTMLGlob("views/*")
	router.Static("/static", "./static")

	router.GET("/", func(c *gin.Context) {
		lang := c.DefaultQuery("lang", "id")

		var header Header
		headerFile := fmt.Sprintf("locales/header_%s.json", lang)
		if err := loadJSONFile(headerFile, &header); err != nil {
			log.Println("Error loading header:", err)
		}

		homeData, err := app.LoadJSONFile(lang)
		if err != nil {
			log.Println("Error loading home data", err)
			homeData = nil
		}

		c.HTML(http.StatusOK, "layout.html", gin.H{
			"Header": header,
			"Home":   homeData,
			"Lang":   lang,
		})
	})

	router.GET("/signup", func(c *gin.Context) {
		c.HTML(http.StatusOK, "signup.html", nil)
	})

	router.GET("/galery", func(c *gin.Context) {
		c.HTML(http.StatusOK, "galery.html", nil)
	})

	router.POST("/submit", controllers.SubmitForm)

	log.Println("Server running on port 8095")
	router.Run(":8095")
}
