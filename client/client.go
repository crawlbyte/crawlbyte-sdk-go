package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Client struct {
	BaseURL string
	APIKey  string
	Client  *http.Client
}

func New(baseURL, apiKey string) *Client {
	return &Client{
		BaseURL: baseURL,
		APIKey:  apiKey,
		Client: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

func (c *Client) DoRequest(method, path string, body interface{}) ([]byte, error) {
	var reqBodyBytes []byte
	var err error

	if body != nil {
		reqBodyBytes, err = json.Marshal(body)
		if err != nil {
			return nil, err
		}
	}

	url := c.BaseURL + path

	var lastErr error

	for attempt := 1; attempt <= 5; attempt++ {
		reqBody := bytes.NewReader(reqBodyBytes)

		req, reqErr := http.NewRequest(method, url, reqBody)
		if reqErr != nil {
			return nil, reqErr
		}

		req.Header.Set("Authorization", c.APIKey)
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Accept", "application/json")

		resp, reqErr := c.Client.Do(req)
		if reqErr != nil {
			lastErr = reqErr
			time.Sleep(time.Duration(attempt) * 500 * time.Millisecond)
			continue
		}

		respBody, reqErr := io.ReadAll(resp.Body)
		_ = resp.Body.Close()
		if reqErr != nil {
			lastErr = reqErr
			time.Sleep(time.Duration(attempt) * 500 * time.Millisecond)
			continue
		}

		if resp.StatusCode >= 500 || resp.StatusCode == 429 {
			lastErr = fmt.Errorf("HTTP %d: %s", resp.StatusCode, string(respBody))
			time.Sleep(time.Duration(attempt) * 500 * time.Millisecond)
			continue
		}

		if resp.StatusCode >= 400 {
			return nil, fmt.Errorf("HTTP %d: %s", resp.StatusCode, string(respBody))
		}

		return respBody, nil
	}

	return nil, fmt.Errorf("request failed after retries: %w", lastErr)
}
