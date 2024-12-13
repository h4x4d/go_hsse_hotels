package handlers

import (
	"context"
	"github.com/go-openapi/runtime/middleware"
	"github.com/h4x4d/go_hsse_hotels/auth/internal/impl"
	"github.com/h4x4d/go_hsse_hotels/auth/internal/models"
	"github.com/h4x4d/go_hsse_hotels/auth/internal/restapi/operations"
)

func (h *Handler) RegisterHandler(api operations.PostRegisterParams) middleware.Responder {
	_, span := h.tracer.Start(context.Background(), "register")
	defer span.End()
	token, err := impl.CreateUser(h.Client, api.Body)
	if err != nil {
		conflict := int64(operations.PostRegisterConflictCode)
		return new(operations.PostRegisterConflict).WithPayload(&models.Error{
			ErrorMessage:    err.Error(),
			ErrorStatusCode: &conflict,
		})
	}
	result := new(operations.PostRegisterOK).WithPayload(&operations.PostRegisterOKBody{
		Token: *token,
	})
	return result
}
