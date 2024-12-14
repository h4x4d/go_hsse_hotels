package handlers

import (
	"context"
	"github.com/go-openapi/runtime/middleware"
	"github.com/h4x4d/go_hsse_hotels/auth/internal/impl"
	"github.com/h4x4d/go_hsse_hotels/auth/internal/models"
	"github.com/h4x4d/go_hsse_hotels/auth/internal/restapi/operations"
	"log/slog"
)

func (h *Handler) RegisterHandler(api operations.PostRegisterParams) middleware.Responder {
	// Tracing
	_, span := h.tracer.Start(context.Background(), "register")
	defer span.End()

	token, err := impl.CreateUser(h.Client, api.Body)
	if err != nil {
		// Logging
		slog.Error(
			"failed register new user",
			slog.String("method", "POST"),
			slog.Group("user-properties",
				slog.String("login", *api.Body.Login),
				slog.String("email", *api.Body.Email),
				slog.Int("telegram-id", int(*api.Body.TelegramID)),
			),
			slog.Int("status_code", operations.PostRegisterConflictCode),
			slog.String("error", err.Error()),
		)

		conflict := int64(operations.PostRegisterConflictCode)
		return new(operations.PostRegisterConflict).WithPayload(&models.Error{
			ErrorMessage:    err.Error(),
			ErrorStatusCode: &conflict,
		})
	}
	// Logging
	slog.Info(
		"register new user",
		slog.String("method", "POST"),
		slog.Group("user-properties",
			slog.String("login", *api.Body.Login),
			slog.String("email", *api.Body.Email),
			slog.Int("telegram-id", int(*api.Body.TelegramID)),
		),
		slog.Int("status_code", operations.PostRegisterOKCode),
	)

	result := new(operations.PostRegisterOK).WithPayload(&operations.PostRegisterOKBody{
		Token: *token,
	})
	return result
}
