package crmclient

import (
	"fmt"
	"net/http"
)

func (c *Client) UpsertSenderName(payload SenderNamePayload, projectKey string) error {
	url := fmt.Sprintf("%s/api/v1/customer/sender-name/callback", c.BaseURL)
	return c.sendJSON(url, payload, projectKey, http.MethodPost)
}

func (c *Client) UpsertSmsCancel(payload SmsCancelPayload, projectKey string) error {
	url := fmt.Sprintf("%s/api/v1/customer/sms-cancel/callback", c.BaseURL)
	return c.sendJSON(url, payload, projectKey, http.MethodPost)
}
