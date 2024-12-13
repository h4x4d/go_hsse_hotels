package handlers

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/h4x4d/go_hsse_hotels/auth/internal/restapi/operations"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func MetricsHandler(p operations.GetMetricsParams) middleware.Responder {
	return NewCustomResponder(p.HTTPRequest, promhttp.Handler())
}
