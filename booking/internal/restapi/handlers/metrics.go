package handlers

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/h4x4d/go_hsse_hotels/booking/internal/restapi/operations/instruments"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func MetricsHandler(p instruments.GetMetricsParams) middleware.Responder {
	return NewCustomResponder(p.HTTPRequest, promhttp.Handler())
}
