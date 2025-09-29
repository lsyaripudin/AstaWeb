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
	Title       string `json:"title"`
	Home        string `json:"home"`
	Foreword    string `json:"foreword"`
	Profile     string `json:"profile"`
	Background  string `json:"background"`
	Vision      string `json:"vision_mission"`
	Structure   string `json:"organization"`
	Report      string `json:"report"`
	Training    string `json:"training_center"`
	Educators   string `json:"educators"`
	Trainees    string `json:"trainee_strengths"`
	Team        string `json:"team"`
	Curriculum  string `json:"curriculum"`
	Handbook    string `json:"handbook"`
	Recruitment string `json:"recruitment"`
	Media       string `json:"media"`
	Highlights  string `json:"highlights"`
	Galery      string `json:"galery"`
	Youtube     string `json:"youtube"`
	ChangeLang  string `json:"change_language"`
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
