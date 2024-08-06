package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type HelloRequest struct {
	Name string `json:"name"`
}

type HelloResponse struct {
	Message string `json:"message"`
}

func main() {
	// Prepare the request payload
	request := HelloRequest{Name: "world"}
	requestBody, err := json.Marshal(request)
	if err != nil {
		log.Fatalf("Failed to marshal request: %v", err)
	}

	// Make the HTTP POST request
	response, err := http.Post("http://localhost:8095/v1/example/echo", "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		log.Fatalf("Failed to make request: %v", err)
	}
	defer response.Body.Close()

	// Read and process the response
	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatalf("Failed to read response body: %v", err)
	}

	// Unmarshal the response
	var helloResponse HelloResponse
	if err := json.Unmarshal(body, &helloResponse); err != nil {
		log.Fatalf("Failed to unmarshal response: %v", err)
	}

	// Print the response
	fmt.Printf("Response: %v\n", helloResponse.Message)
}
