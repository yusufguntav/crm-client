package crmclient

import (
	"fmt"
	"net/http"
)

func (c *Client) UpsertCustomer(payload CustomerPayload, projectKey string) error {
	url := fmt.Sprintf("%s/api/v1/customer/callback", c.BaseURL)
	return c.sendJSON(url, payload, projectKey)
}

func (c *Client) DeleteCustomer(payload CustomerDeletePayload, projectKey string) error {
	url := fmt.Sprintf("%s/api/v1/customer/callback", c.BaseURL)
	return c.sendJSON(url, payload, projectKey, http.MethodDelete)
}
