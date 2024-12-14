package handlers

import (
	"context"
	"github.com/go-openapi/runtime/middleware"
	"github.com/h4x4d/go_hsse_hotels/auth/internal/impl"
	"github.com/h4x4d/go_hsse_hotels/auth/internal/models"
	"github.com/h4x4d/go_hsse_hotels/auth/internal/restapi/operations"
	"log/slog"
)

func (h *Handler) ChangePasswordHandler(api operations.PostChangePasswordParams) middleware.Responder {
	// Tracing
	_, span := h.tracer.Start(context.Background(), "change_password")
	defer span.End()

	token, err := impl.ChangePasswordUser(h.Client, api.Body)
	conflict := int64(operations.PostChangePasswordBadRequestCode)
	if err != nil {
		// Logging
		slog.Error(
			"failed to change password",
			slog.String("method", "POST"),
			slog.Group("user-properties",
				slog.String("login", *api.Body.Login),
			),
			slog.Int("status_code", operations.PostChangePasswordBadRequestCode),
			slog.String("error", err.Error()),
		)

		return new(operations.PostChangePasswordUnauthorized).WithPayload(&models.Error{
			ErrorMessage:    err.Error(),
			ErrorStatusCode: &conflict,
		})
	}
	// Logging
	slog.Info(
		"user changed password",
		slog.String("method", "POST"),
		slog.Group("user-properties",
			slog.String("login", *api.Body.Login),
		),
		slog.Int("status_code", operations.PostChangePasswordOKCode),
	)

	result := new(operations.PostChangePasswordOK).WithPayload(&operations.PostChangePasswordOKBody{
		Token: *token,
	})
	return result
}
