package utils

import (
	"github.com/go-openapi/runtime/middleware"
	"net/http"
)

func CatchPanic(responder *middleware.Responder) {
	if err := recover(); err != nil {
		*responder = middleware.Error(http.StatusInternalServerError, err)
	}
}
