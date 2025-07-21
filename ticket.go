package crmclient

import "fmt"

func (c *Client) CreateTicket(payload TicketPayload) error {
	url := fmt.Sprintf("%s/api/v1/ticket/callback", c.BaseURL)
	return c.sendJSON(url, payload)
}
