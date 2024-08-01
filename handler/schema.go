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

func SchemaList(c *gin.Context) {
	token, err := c.Cookie("jwt-token")
	if err != nil {

		// For any other error, return a bad request error
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	url := config.SERVER_URL + "/api/common/getSchemaList"

	var input datahandler.SelectedITR

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	jsonData, err := json.Marshal(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// Create the request
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	// Set the headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))

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
	if schemaList, exists := responseBody["schemaList"]; exists {
		c.JSON(http.StatusOK, gin.H{"schemaList": schemaList})
		return
	}

}

func Schedule(c *gin.Context) {

	token, err := c.Cookie("jwt-token")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	itr := c.Query("itr")
	schema := c.Query("schema")

	// Define the request body as a map or a struct
	myBody := map[string]string{
		"ITR":    itr,
		"Schema": schema,
	}

	// Convert the body to JSON
	jsonData, err := json.Marshal(myBody)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Create the request
	url := config.SERVER_URL + "/api/common/getScheduleTree"
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// Set the headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
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
	if tree, exists := responseBody["tree"]; exists {
		c.HTML(http.StatusOK, "schedule", gin.H{"tree": tree})
	}

}
func ViewSchedule(c *gin.Context) {

	token, err := c.Cookie("jwt-token")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	itr := c.Query("itr")
	schema := c.Query("schema")
	ElementID := c.Query("elementID")
	// Define the request body as a map or a struct
	myBody := map[string]string{
		"ITR":    itr,
		"Schema": schema,
	}

	// Convert the body to JSON
	jsonData, err := json.Marshal(myBody)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Create the request
	url := config.SERVER_URL + "/api/common/getScheduleTree"
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// Set the headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
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
	if tree, exists := responseBody["tree"]; exists {
		var treenode datahandler.TreeNode
		data, err := json.Marshal(tree)
		if err != nil {
			fmt.Println("Error marshalling:", err)
			return
		}
		err = json.Unmarshal(data, &treenode)
		if err != nil {
			fmt.Println("Error unmarshalling:", err)
			return
		}
		for _, node := range treenode.Children {
			if node.NodeData.ElementID == ElementID {
				c.HTML(http.StatusOK, "viewSchedule", gin.H{"node": node})
				return
			}
		}
	}
}

func Welcome(c *gin.Context) {
	c.HTML(http.StatusOK, "welcome", gin.H{})
}

func EditSchedule(c *gin.Context) {
	token, err := c.Cookie("jwt-token")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	itr := c.Query("itr")
	schema := c.Query("schema")
	ElementID := c.Query("elementID")
	// Define the request body as a map or a struct
	myBody := map[string]string{
		"ITR":       itr,
		"Schema":    schema,
		"ElementID": ElementID,
	}
	jsonData, err := json.Marshal(myBody)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Create the request
	url := config.SERVER_URL + "/api/common/getElement"
	req, err := http.NewRequest("GET", url, bytes.NewBuffer(jsonData))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// Set the headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer resp.Body.Close()
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

	if element, exists := responseBody["element"]; exists {
		c.HTML(http.StatusOK, "editSchedule", gin.H{"element": element})
	}
}
func AdvanceSetting(c *gin.Context) {
	token, err := c.Cookie("jwt-token")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	itr := c.Query("itr")
	schema := c.Query("schema")
	ElementID := c.Query("elementID")
	// Define the request body as a map or a struct
	myBody := map[string]string{
		"ITR":       itr,
		"Schema":    schema,
		"ElementID": ElementID,
	}
	jsonData, err := json.Marshal(myBody)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Create the request
	url := config.SERVER_URL + "/api/common/getElement"
	req, err := http.NewRequest("GET", url, bytes.NewBuffer(jsonData))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// Set the headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer resp.Body.Close()
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

	if element, exists := responseBody["element"]; exists {
		c.HTML(http.StatusOK, "advanceSetting", gin.H{"element": element})
	}

}
