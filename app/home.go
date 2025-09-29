package app

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"asta-karya/models"
)

type sheetResponse struct {
	Values [][]string `json:"values"`
}

func fetchMapContent(url string) (string, error) {
	client := http.Client{Timeout: 5 * time.Second}
	resp, err := client.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var payload sheetResponse
	if err := json.NewDecoder(resp.Body).Decode(&payload); err != nil {
		return "", err
	}

	if len(payload.Values) == 0 || len(payload.Values[0]) == 0 {
		return "", fmt.Errorf("no values returned")
	}

	return payload.Values[0][0], nil
}

// LoadJSONFile reads the localized home configuration file for the given language
// and returns the structured content for the landing page.
func LoadJSONFile(lang string) (*models.Home, error) {
	filePath := fmt.Sprintf("locales/home_%s.json", lang)

	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var content models.Home
	if err := json.Unmarshal(data, &content); err != nil {
		return nil, err
	}

	if content.MapSourceURL != "" {
		if mapValue, err := fetchMapContent(content.MapSourceURL); err != nil {
			log.Printf("failed to fetch map content from sheet: %v", err)
		} else if mapValue != "" {
			content.MapContent = mapValue
		}
	}

	return &content, nil
}
