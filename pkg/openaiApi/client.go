// Setup and authenticate OpenAI API client
package openaiapi

import (
	"net/http"
)

// OpenAIClient represents a client to connect with OpenAI API
type OpenAIClient struct {
	httpClient *http.Client
	apiKey     string
}
