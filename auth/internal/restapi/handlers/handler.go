package handlers

import (
	"auth/internal/config"
)

type Handler struct {
	Client *config.Client
}

func NewHandler() (*Handler, error) {
	client, err := config.NewClient()
	if err != nil {
		return nil, err
	}
	return &Handler{client}, nil
}
