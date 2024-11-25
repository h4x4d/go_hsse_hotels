// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"github.com/h4x4d/go_hsse_hotels/booking/internal/restapi/operations"
	customer2 "github.com/h4x4d/go_hsse_hotels/booking/internal/restapi/operations/customer"
	"github.com/h4x4d/go_hsse_hotels/booking/internal/restapi/operations/hotelier"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
)

//go:generate swagger generate server --target ../../booking --name HotelsBooking --spec ../docs/swagger/booking.yaml --principal interface{}

func configureFlags(api *operations.HotelsBookingAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.HotelsBookingAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.UseSwaggerUI()
	// To continue using redoc as your UI, uncomment the following line
	// api.UseRedoc()

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	// Applies when the "api_key" header is set
	api.APIKeyAuth = func(token string) (interface{}, error) {
		return "OK", nil
	}

	// Set your custom authorizer if needed. Default one is security.Authorized()
	// Expected interface runtime.Authorizer
	//
	// Example:
	// api.APIAuthorizer = security.Authorized()

	if api.CustomerCreateBookingHandler == nil {
		api.CustomerCreateBookingHandler = customer2.CreateBookingHandlerFunc(func(params customer2.CreateBookingParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation customer.CreateBooking has not yet been implemented")
		})
	}
	if api.HotelierGetBookingHandler == nil {
		api.HotelierGetBookingHandler = hotelier.GetBookingHandlerFunc(func(params hotelier.GetBookingParams) middleware.Responder {
			return middleware.NotImplemented("operation customer.GetBooking has not yet been implemented")
		})
	}
	if api.CustomerGetBookingByIDHandler == nil {
		api.CustomerGetBookingByIDHandler = customer2.GetBookingByIDHandlerFunc(func(params customer2.GetBookingByIDParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation customer.GetBookingByID has not yet been implemented")
		})
	}
	if api.CustomerUpdateBookingHandler == nil {
		api.CustomerUpdateBookingHandler = customer2.UpdateBookingHandlerFunc(func(params customer2.UpdateBookingParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation customer.UpdateBooking has not yet been implemented")
		})
	}

	api.PreServerShutdown = func() {}

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix".
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation.
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics.
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
