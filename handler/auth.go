package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"golang-frontend/config"
	"golang-frontend/datahandler"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Logout(c *gin.Context) {
	token, err := c.Cookie("jwt-token")
	if err != nil {

		// For any other error, return a bad request error
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	url := config.SERVER_URL + "/api/common/logout"

	// Create the request
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	// Set the headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return

	}
	defer resp.Body.Close()
	var responseBody map[string]interface{}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	err = json.Unmarshal(body, &responseBody)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	if msg, exists := responseBody["message"]; exists {
		if msg == "logged out" {
			c.JSON(http.StatusOK, gin.H{"message": msg})
			return
		}

	}
	c.JSON(http.StatusInternalServerError, gin.H{"error": responseBody["error"]})
	return

}
func Login(c *gin.Context) {

	token, err := c.Cookie("jwt-token")
	if err != nil {
		// If the cookie is not found, return an unauthorized error
		if err == http.ErrNoCookie {
			c.HTML(http.StatusInternalServerError, "login", gin.H{"error": err})
			return
		}
		// For any other error, return a bad request error
		c.HTML(http.StatusInternalServerError, "login", gin.H{"error": err})
		return
	}

	url := config.SERVER_URL + "/api/common/verify"

	// Create the request
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "login", gin.H{"error": err})
		return
	}

	// Set the headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "login", gin.H{"error": err})
		return

	}
	defer resp.Body.Close()

	// Read and print the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "login", gin.H{"error": err})
		return
	}
	var responseBody map[string]interface{}
	err = json.Unmarshal(body, &responseBody)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "login", gin.H{"error": err})
		return
	}
	if msg, exists := responseBody["message"]; exists {
		if msg == "logged in" {
			c.HTML(http.StatusOK, "welcome", gin.H{})
			return
		}
	}
	c.HTML(http.StatusUnauthorized, "login", gin.H{})
}

func LoginPost(c *gin.Context) {
	var input datahandler.LoginInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Define the request body as a map or a struct
	myBody := map[string]string{
		"username": input.Username,
		"password": input.Password,
	}

	// Convert the body to JSON
	jsonData, err := json.Marshal(myBody)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Create the request
	url := config.SERVER_URL + "/api/login"
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Set the headers
	req.Header.Set("Content-Type", "application/json")

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer resp.Body.Close()

	//You can also read and print the response body if needed
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var responseBody map[string]interface{}
	err = json.Unmarshal(body, &responseBody)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if err, exists := responseBody["error"]; exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	// Check if the "token" key is present

	if msg := responseBody["message"]; msg == "already logged in" {
		token := responseBody["token"]
		c.SetCookie("new-jwt-token", token.(string), 24*3600, "/", "localhost", false, true)
		c.JSON(http.StatusOK, gin.H{"message": msg})
	} else {
		if token, exists := responseBody["token"]; exists {
			c.SetCookie("jwt-token", token.(string), 24*3600, "/", "localhost", false, true)
			c.JSON(http.StatusOK, gin.H{"message": "logged in"})
		}
	}
}
func LoginConfirm(c *gin.Context) {
	token, err := c.Cookie("new-jwt-token")
	if err != nil {
		// If the cookie is not found, return an unauthorized error
		if err == http.ErrNoCookie {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "No new token found"})
			return
		}
		// For any other error, return a bad request error
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error retrieving token"})
		return
	}
	// Define the request body as a map or a struct
	requestBody := map[string]string{
		// Add your request body fields here
	}

	// Convert the body to JSON
	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return
	}

	// Create the request
	url := config.SERVER_URL + "/api/loginConfirm"
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	// Set the headers
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	req.Header.Set("Content-Type", "application/json")

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	defer resp.Body.Close()

	// Read and print the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	var responseBody map[string]interface{}
	err = json.Unmarshal(body, &responseBody)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	if err, exists := responseBody["error"]; exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	if msg, exists := responseBody["message"]; exists {
		c.SetCookie("jwt-token", token, 24*3600, "/", "localhost", false, true)
		c.JSON(http.StatusOK, gin.H{"message": msg})
	}

}

