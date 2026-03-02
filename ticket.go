package crmclient

import "fmt"

func (c *Client) CreateTicket(payload TicketPayload, projectKey string) error {
	url := fmt.Sprintf("%s/api/v1/ticket/callback", c.BaseURL)
	if payload.TicketKey != "" {
		c.TicketKey = payload.TicketKey
	}
	return c.sendJSON(url, payload, projectKey)
}
