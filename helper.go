package crmclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func (c *Client) sendJSON(url string, body any, methodOverride ...string) error {
	data, _ := json.Marshal(body)
	method := http.MethodPost
	if len(methodOverride) > 0 {
		method = methodOverride[0]
	}

	req, err := http.NewRequest(method, url, bytes.NewBuffer(data))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("ProjectSecretKey", c.ProjectKey)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 300 {
		return fmt.Errorf("unexpected response code: %d", resp.StatusCode)
	}
	return nil
}
