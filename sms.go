package crmclient

import (
	"fmt"
	"net/http"
)

func (c *Client) UpsertSenderName(payload SenderNamePayload) error {
	url := fmt.Sprintf("%s/api/v1/customer/sender-name/callback", c.BaseURL)
	return c.sendJSON(url, payload, http.MethodPost)
}

func (c *Client) UpsertSmsCancel(payload SmsCancelPayload) error {
	url := fmt.Sprintf("%s/api/v1/customer/sms-cancel/callback", c.BaseURL)
	return c.sendJSON(url, payload, http.MethodPost)
}
