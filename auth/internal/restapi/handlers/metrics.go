package handlers

import (
	"auth/internal/restapi/operations"
	"github.com/go-openapi/runtime/middleware"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func MetricsHandler(p operations.GetMetricsParams) middleware.Responder {
	return NewCustomResponder(p.HTTPRequest, promhttp.Handler())
}
