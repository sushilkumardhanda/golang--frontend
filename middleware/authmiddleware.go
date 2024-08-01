package middleware

import (
	"encoding/json"
	"fmt"
	"golang-frontend/config"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

func JwtAuthHtmlMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := c.Cookie("jwt-token")
		if err != nil {
			// If the cookie is not found, return an unauthorized error
			if err == http.ErrNoCookie {
				c.HTML(http.StatusInternalServerError, "login", gin.H{"error": err})
				c.Abort()
				return
			}
			// For any other error, return a bad request error
			c.HTML(http.StatusInternalServerError, "login", gin.H{"error": err})
			c.Abort()
			return
		}

		url := config.SERVER_URL+"/api/common/verify"

		// Create the request
		req, err := http.NewRequest("POST", url, nil)
		if err != nil {
			c.HTML(http.StatusInternalServerError, "login", gin.H{"error": err})
			c.Abort()
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
			c.Abort()
			return

		}
		defer resp.Body.Close()

		// Read and print the response body
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			c.HTML(http.StatusInternalServerError, "login", gin.H{"error": err})
			c.Abort()
			return
		}
		var responseBody map[string]interface{}
		err = json.Unmarshal(body, &responseBody)
		if err != nil {
			c.HTML(http.StatusInternalServerError, "login", gin.H{"error": err})
			c.Abort()
			return
		}
		if msg, exists := responseBody["message"]; exists {
			if msg == "logged in" {
				c.Next()
				return
			}
		}
		c.HTML(http.StatusUnauthorized, "login", gin.H{})
		c.Abort()
	}
}

func JwtAuthJsonMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := c.Cookie("jwt-token")
		if err != nil {
			// If the cookie is not found, return an unauthorized error
			if err == http.ErrNoCookie {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err})
				c.Abort()
				return
			}
			// For any other error, return a bad request error
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			c.Abort()
			return
		}

		url := config.SERVER_URL+"/api/common/verify"

		// Create the request
		req, err := http.NewRequest("POST", url, nil)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			c.Abort()
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
			c.Abort()
			return

		}
		defer resp.Body.Close()

		// Read and print the response body
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			c.Abort()
			return
		}
		var responseBody map[string]interface{}
		err = json.Unmarshal(body, &responseBody)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			c.Abort()
			return
		}
		if msg, exists := responseBody["message"]; exists {
			if msg == "logged in" {
				c.Next()
				return
			}
		}
		c.JSON(http.StatusUnauthorized, gin.H{})
		c.Abort()
		return

	}
}
