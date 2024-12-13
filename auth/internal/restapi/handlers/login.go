package handlers

import (
	"context"
	"github.com/go-openapi/runtime/middleware"
	"github.com/h4x4d/go_hsse_hotels/auth/internal/impl"
	"github.com/h4x4d/go_hsse_hotels/auth/internal/models"
	"github.com/h4x4d/go_hsse_hotels/auth/internal/restapi/operations"
)

func (h *Handler) LoginHandler(api operations.PostLoginParams) middleware.Responder {
	_, span := h.tracer.Start(context.Background(), "login")
	defer span.End()
	token, err := impl.LoginUser(h.Client, api.Body)
	conflict := int64(operations.PostLoginUnauthorizedCode)
	if err != nil {
		return new(operations.PostRegisterConflict).WithPayload(&models.Error{
			ErrorMessage:    err.Error(),
			ErrorStatusCode: &conflict,
		})
	}
	result := new(operations.PostLoginOK).WithPayload(&operations.PostLoginOKBody{
		Token: *token,
	})
	return result
}
