// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"fmt"
	"github.com/h4x4d/go_hsse_hotels/hotel/internal/restapi/handlers"
	"github.com/h4x4d/go_hsse_hotels/pkg/client"
	"log"
	"net/http"
	"os"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/h4x4d/go_hsse_hotels/hotel/internal/restapi/operations"
	"github.com/h4x4d/go_hsse_hotels/hotel/internal/restapi/operations/hotel"
)

//go:generate swagger generate server --target ../../hotel --name HotelsHotel --spec ../api/swagger/hotels.yaml --principal interface{}

func configureFlags(api *operations.HotelsHotelAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.HotelsHotelAPI) http.Handler {
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

	manager, err := client.NewClient()
	if err != nil {
		log.Fatal(err)
	}
	// Applies when the "api_key" header is set
	api.APIKeyAuth = func(token string) (interface{}, error) {
		userId, err := manager.CheckToken(token)
		if err != nil {
			return nil, err
		}
		return userId, nil
	}

	// Set your custom authorizer if needed. Default one is security.Authorized()
	// Expected interface runtime.Authorizer
	//
	// Example:
	// api.APIAuthorizer = security.Authorized()

	// catch cases when error is nil but the first argument returning from function is nil TODO
	// maybe this have been already caught, I dont remember clearly

	// creating database_service and then add it to context of handlers
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"), "db", os.Getenv("POSTGRES_PORT"), os.Getenv("HOTEL_DB_NAME"))
	handler, makeErr := handlers.NewHandler(connStr)
	for makeErr != nil {
		handler, makeErr = handlers.NewHandler(connStr)
	}

	api.HotelCreateHotelHandler = hotel.CreateHotelHandlerFunc(handler.CreateHotelHandler)
	api.HotelGetHotelsHandler = hotel.GetHotelsHandlerFunc(handler.GetHotelsHandler)
	api.HotelGetHotelByIDHandler = hotel.GetHotelByIDHandlerFunc(handler.GetHotelByIDHandler)
	api.HotelUpdateHotelHandler = hotel.UpdateHotelHandlerFunc(handler.UpdateHotelHandler)

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
