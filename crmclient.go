// crmclient.go
package crmclient

import (
	"net/http"
)

var defaultBaseURL = "https://apicrm.vatansoft.net"

type Client struct {
	BaseURL    string
	ProjectKey string
	TicketKey  string
	httpClient *http.Client
}

func New(projectKey string) *Client {
	return &Client{
		ProjectKey: projectKey,
		httpClient: &http.Client{},
		BaseURL:    defaultBaseURL,
	}
}
