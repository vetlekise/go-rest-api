// Package main is the entry point for the go-rest-api tool.
package main

import (
	"context"
	"flag"
	"log/slog"
	"net/http"
	"time"

	"github.com/vetlekise/go-rest-api/client"
)

var url string
var httpClient = &http.Client{Timeout: 10 * time.Second}

func init() {
	flag.StringVar(&url, "URL", "https://jsonplaceholder.typicode.com/", "The URL to test CRUD operations.")
	flag.Parse()
}

func main() {
	c := client.New(url, httpClient)

	todo, err := c.GetTodo(context.Background())
	if err != nil {
		slog.Error("GET request failed", "err", err)
		return
	}
	slog.Info("todo", "id", todo.ID, "title", todo.Title, "completed", todo.Completed)

	post := client.PostBody{
		UserID: 1,
		ID:     1,
		Title:  "my title",
		Body:   "my body",
	}
	result, err := c.PostPost(context.Background(), post)
	if err != nil {
		slog.Error("POST request failed", "err", err)
		return
	}
	slog.Info("post", "userId", result.UserID, "id", result.ID, "title", result.Title, "body", result.Body)

	if err := c.DeletePost(context.Background()); err != nil {
		slog.Error("DELETE request failed", "err", err)
		return
	}
	slog.Info("delete successful")
}
