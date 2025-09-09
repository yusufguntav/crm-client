package crmclient

import (
	"fmt"
	"net/http"
)

func (c *Client) UpsertCustomer(payload CustomerPayload) error {
	url := fmt.Sprintf("%s/api/v1/customer/callback", c.BaseURL)
	return c.sendJSON(url, payload)
}

func (c *Client) DeleteCustomer(payload CustomerDeletePayload) error {
	url := fmt.Sprintf("%s/api/v1/customer/callback", c.BaseURL)
	return c.sendJSON(url, payload, http.MethodDelete)
}
