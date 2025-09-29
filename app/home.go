package app

import (
	"encoding/json"
	"fmt"
	"os"

	"asta-karya/models"
)

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

	return &content, nil
}
