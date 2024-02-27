// Setup and authenticate GMAIL API client
package gmailapi

import (
	"context"
	"log"
	"net/http"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/gmail/v1"

// GmailClient struct holds the necessary details for the Gmail client
type GmailClient struct {
	service *gmail.Service
}

// NewGmailClient initializes a new Gmail client with the necessary credentials
func NewGmailClient(ctx context.Context, config *oauth2.Config) *GmailClient {
	// TODO: Implement a token retrieval and create an http.Client with the token
	// client := config.Client(ctx, token)

	srv, err := gmail.New(client)
	if err != nil {
		log.Fatalf("Unable to create a Gmail client: %v", err)
	}

	return &GmailClient{
		Service: srv,
	}
}

//GmailClient represents a client to connect with Gmail API
