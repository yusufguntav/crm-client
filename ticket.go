package crmclient

import "fmt"

func (c *Client) CreateTicket(payload TicketPayload) error {
	url := fmt.Sprintf("%s/api/v1/ticket/callback", c.BaseURL)
	c.TicketKey = payload.TicketKey
	return c.sendJSON(url, payload)
}
