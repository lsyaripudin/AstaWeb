package controllers

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type LoginResponse struct {
	ResultCode string `json:"resultCode"`
	Message    string `json:"message"`
	Data       struct {
		AccessToken string `json:"accessToken"`
	} `json:"data"`
}

type Siswa struct {
	Name              string `form:"name"`
	Nik               string `form:"nik"`
	Email             string `form:"email"`
	Tempat_Lahir      string `form:"tempat_lahir"`
	Tanggal_Lahir     string `form:"tanggal_lahir"`
	Jenis_Kelamin     string `form:"jenis_kelamin"`
	Alamat            string `form:"alamat"`
	RT_RW             string `form:"rt_rw"`
	Kelurahan_Desa    string `form:"kelurahan_desa"`
	Kecamatan         string `form:"kecamatan"`
	Agama             string `form:"agama"`
	Status_Perkawinan string `form:"status_perkawinan"`
	Pekerjaan         string `form:"pekerjaan"`
	Kewarganegaraan   string `form:"kewarganegaraan"`
	Berlaku_Hingga    string `form:"berlaku_hingga"`
}

type ApiResponse struct {
	ResultCode string `json:"resultCode"`
	Message    string `json:"message"`
}

func GetToken() string {
	loginData := map[string]string{
		"username": "testuser",
		"password": "12345",
	}

	loginJSON, _ := json.Marshal(loginData)
	loginResp, err := http.Post("http://localhost:8096/login", "application/json", bytes.NewBuffer(loginJSON))
	if err != nil {
		log.Fatalf("Error while logging in: %v", err)
	}
	defer loginResp.Body.Close()

	var loginResponse LoginResponse
	body, _ := io.ReadAll(loginResp.Body)
	if err := json.Unmarshal(body, &loginResponse); err != nil {
		log.Fatalf("Error while parsing login response: %v", err)
	}

	if loginResponse.ResultCode != "00" {
		log.Fatalf("Login failed: %s", loginResponse.Message)
	}

	return loginResponse.Data.AccessToken
}

func SubmitForm(c *gin.Context) {
	var siswa Siswa
	if err := c.ShouldBind(&siswa); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tanggalLahir, err := time.Parse("2006-01-02", siswa.Tanggal_Lahir)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid date format"})
		return
	}

	siswa.Tanggal_Lahir = tanggalLahir.Format(time.RFC3339)

	accessToken := GetToken()
	siswaJSON, err := json.Marshal(siswa)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to marshal JSON"})
		return
	}

	req, err := http.NewRequest("POST", "http://localhost:8096/api/insertsiswa", bytes.NewBuffer(siswaJSON))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create request"})
		return
	}

	req.Header.Set("Authorization", "Bearer "+accessToken)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send request"})
		return
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read response"})
		return
	}

	var apiResponse ApiResponse
	if err := json.Unmarshal(respBody, &apiResponse); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse response"})
		return
	}

	c.HTML(http.StatusOK, "signup.html", gin.H{"message": apiResponse.Message})
	c.HTML(http.StatusOK, "galery.html", gin.H{"message": apiResponse.Message})
}
