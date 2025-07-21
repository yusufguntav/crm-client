package crmclient

import (
	"fmt"
	"net/http"
)

func (c *Client) UpsertCustomer(payload CustomerPayload) error {
	url := fmt.Sprintf("%s/api/v1/customer/callback", c.BaseURL)
	return c.sendJSON(url, payload)
}

func (c *Client) DeleteCustomer(projectID string, idInProject string) error {
	url := fmt.Sprintf("%s/api/v1/customer/callback/%s", c.BaseURL, projectID)
	body := map[string]string{"id_in_project": idInProject}
	return c.sendJSON(url, body, http.MethodDelete)
}
