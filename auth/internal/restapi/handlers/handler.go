package handlers

import (
	"auth/internal/client"
)

type Handler struct {
	Client *client.Client
}

func NewHandler() (*Handler, error) {
	client, err := client.NewClient()
	if err != nil {
		return nil, err
	}
	return &Handler{client}, nil
}
