// Setup and authenticate GMAIL API client
package gmailapi

//GmailClient represents a client to connect with Gmail API
type GmailClient struct {
	service *gmail.Service
}
