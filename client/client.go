package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type TodoBody struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

type PostBody struct {
	UserID int    `json:"userId"`
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
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

// GetTodo sends a GET request to the 'todos/1' endpoint and returns the response body and error.
func (c *Client) GetTodo(ctx context.Context) (*TodoBody, error) {
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

	var todoBody TodoBody
	if err := json.NewDecoder(resp.Body).Decode(&todoBody); err != nil {
		return nil, fmt.Errorf("decode response: %w", err)
	}
	return &todoBody, nil
}

// PostPost sends a POST request to the 'posts' endpoint and returns the response body and error.
func (c *Client) PostPost(ctx context.Context, post PostBody) (*PostBody, error) {
	body, err := json.Marshal(post)
	if err != nil {
		return nil, fmt.Errorf("marshal: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, c.baseURL+"posts", bytes.NewReader(body))
	if err != nil {
		return nil, fmt.Errorf("build request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	resp, err := c.http.Do(req)
	if err != nil {
		return nil, fmt.Errorf("do request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("unexpected status %d", resp.StatusCode)
	}

	var postBody PostBody
	if err := json.NewDecoder(resp.Body).Decode(&postBody); err != nil {
		return nil, fmt.Errorf("decode response: %w", err)
	}
	return &postBody, nil
}

// PostDelete sends a DELETE request to the 'posts/2' endpoint and only returns an error because the response body is empty.
func (c *Client) DeletePost(ctx context.Context) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodDelete, c.baseURL+"posts/2", nil)
	if err != nil {
		return fmt.Errorf("build request: %w", err)
	}
	req.Header.Set("Accept", "application/json")

	resp, err := c.http.Do(req)
	if err != nil {
		return fmt.Errorf("do request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status %d", resp.StatusCode)
	}

	return nil
}
