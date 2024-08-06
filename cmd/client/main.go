package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	ddhttp "gopkg.in/DataDog/dd-trace-go.v1/contrib/net/http"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
)

type HelloRequest struct {
	Name string `json:"name"`
}

type HelloResponse struct {
	Message string `json:"message"`
}

func main() {
	// Start the DataDog tracer
	tracer.Start(
		tracer.WithEnv("dev"),
		tracer.WithService("example-client-service"),
		tracer.WithGlobalTag("component", "example-client-component"),
		tracer.WithServiceVersion("v1.3.0"),
		// tc.WithAgentAddr(t.Config.AgentHost),
		tracer.WithAnalytics(true),
		tracer.WithRuntimeMetrics(),
	)

	defer tracer.Stop()

	// Create a traced HTTP client
	client := ddhttp.WrapClient(http.DefaultClient)

	// Prepare the request payload
	request := HelloRequest{Name: "world"}
	requestBody, err := json.Marshal(request)
	if err != nil {
		log.Fatalf("Failed to marshal request: %v", err)
	}

	// Create a new HTTP request with a context that includes the trace span
	req, err := http.NewRequest("POST", "http://localhost:8095/v1/example/echo", bytes.NewBuffer(requestBody))
	if err != nil {
		log.Fatalf("Failed to create request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	// Create a span for the HTTP request
	span, ctx := tracer.StartSpanFromContext(context.Background(), "client.http.request")
	defer span.Finish()
	req = req.WithContext(ctx)
	// Make the HTTP POST request
	response, err := client.Do(req)
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
