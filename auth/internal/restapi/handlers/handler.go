package handlers

import (
    "auth/internal/utils"
)

type Handler struct {
    Client *utils.Client
}

func NewHandler() (*Handler, error) {
    client, err := utils.NewClient()
    if err != nil {
        return nil, err
    }
    return &Handler{client}, nil
}
