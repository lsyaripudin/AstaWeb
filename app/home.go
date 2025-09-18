package app

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"asta-karya/models"
)

func FetchGoogleSheetData() (string, error) {
	apiURL := "https://sheets.googleapis.com/v4/spreadsheets/1V8c2Zx4j8Lf6XgD4g3eskE-QCw3e5S8qPaBIJrsYlPY/values/Sheet1!B2?key=AIzaSyBof-SD7PX06Fpx_uZbZhpI5w4S8iAXqm4"

	resp, err := http.Get(apiURL)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("API request failed with status: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var sheetData models.GoogleSheetResponse
	err = json.Unmarshal(body, &sheetData)
	if err != nil {
		return "", err
	}

	if len(sheetData.Values) > 0 && len(sheetData.Values[0]) > 0 {
		return sheetData.Values[0][0], nil
	}

	return "No Data", nil
}

func UpdateJSONFile(lang string, newMapContent string) error {
	filePath := fmt.Sprintf("locales/home_%s.json", lang)

	jsonData, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	var sheet models.Home
	err = json.Unmarshal(jsonData, &sheet)
	if err != nil {
		return err
	}

	sheet.MapContent = newMapContent

	updatedJSON, err := json.MarshalIndent(sheet, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(filePath, updatedJSON, 0644)
}

func LoadJSONFile(lang string) (*models.Home, error) {
	homeContent, err := FetchGoogleSheetData()
	if err != nil {
		log.Println("Error fetching Google Sheets data:", err)
		homeContent = "Error fetching data"
	} else {
		err = UpdateJSONFile(lang, homeContent)
		if err != nil {
			log.Println("Error updating JSON file:", err)
		}
	}

	filePath := fmt.Sprintf("locales/home_%s.json", lang)

	jsonData, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var sheet models.Home
	err = json.Unmarshal(jsonData, &sheet)
	if err != nil {
		return nil, err
	}

	return &sheet, nil
}
