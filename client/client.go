package client

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type Body struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

type Client struct {
	http    *http.Client
	baseURL string
}

func New(baseURL string, http *http.Client) *Client {
	return &Client{
		http:    http,
		baseURL: baseURL,
	}
}

func (c *Client) GetTodo(ctx context.Context) (*Body, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, c.baseURL+"todos/1", nil)
	if err != nil {
		return nil, fmt.Errorf("build request: %w", err)
	}
	req.Header.Set("Accept", "application/json")

	resp, err := c.http.Do(req)
	if err != nil {
		return nil, fmt.Errorf("do request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status %d", resp.StatusCode)
	}

	var body Body
	if err := json.NewDecoder(resp.Body).Decode(&body); err != nil {
		return nil, fmt.Errorf("decode response: %w", err)
	}
	return &body, nil
}
