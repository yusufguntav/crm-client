// crmclient.go
package crmclient

import (
	"net/http"
)

type Client struct {
	BaseURL    string
	ProjectKey string
	httpClient *http.Client
}

func New(baseURL, projectKey string) *Client {
	return &Client{
		BaseURL:    baseURL,
		ProjectKey: projectKey,
		httpClient: &http.Client{},
	}
}
