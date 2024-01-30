package ApiClient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	apiKey      = "Your-api-key-will-be-here"
	apiEndpoint = "https://api.openai.com/v1/chat/completions"
)

func HttpApiRequest(c *gin.Context) {

	// Define the system, chat history, and user messages
	system := []map[string]interface{}{{"role": "system", "content": "You are HappyBot."}}
	chatHistory := []map[string]interface{}{} // Add previous chat history here if available
	user := []map[string]interface{}{{"role": "user", "content": "Are you fully operational?"}}

	// Combine the messages for the chat completion
	messages := append(system, append(chatHistory, user...)...)

	// Create the request body
	requestBody := map[string]interface{}{
		"messages":   messages,
		"model":      "gpt-3.5-turbo",
		"max_tokens": 25,
		"top_p":      0.9,
	}

	// Convert the request body to JSON
	requestBodyBytes, err := json.Marshal(requestBody)
	if err != nil {
		log.Fatalf("Error encoding request body: %v", err)
	}

	// Send the HTTP request to the OpenAI API
	req, err := http.NewRequest("POST", apiEndpoint, bytes.NewBuffer(requestBodyBytes))
	if err != nil {
		log.Fatalf("Error creating HTTP request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiKey))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error sending HTTP request: %v", err)
	}
	defer resp.Body.Close()

	// Read the response body
	responseBodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading response body: %v", err)
	}

	// Unmarshal the response JSON
	var responseData map[string]interface{}
	err = json.Unmarshal(responseBodyBytes, &responseData)
	if err != nil {
		log.Fatalf("Error decoding JSON response: %v", err)
	}

	// Extract the content from the JSON response
	//content := responseData["choices"].([]interface{})[0].(map[string]interface{})["message"].(map[string]interface{})["content"].(string)
	//fmt.Println(content)

	c.JSON(200, gin.H{
		"status": "Ok",
		"Data":   responseData,
	})

}
