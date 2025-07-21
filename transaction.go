package crmclient

import "fmt"

func (c *Client) SendTransaction(payload TransactionPayload) error {
	url := fmt.Sprintf("%s/api/v1/customer/transaction/callback", c.BaseURL)
	return c.sendJSON(url, payload)
}
