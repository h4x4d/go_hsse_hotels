// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"github.com/h4x4d/go_hsse_hotels/hotel/internal/restapi/operations"
	hotel2 "github.com/h4x4d/go_hsse_hotels/hotel/internal/restapi/operations/hotel"
	room2 "github.com/h4x4d/go_hsse_hotels/hotel/internal/restapi/operations/room"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
)

//go:generate swagger generate server --target ../../hotel --name HotelsHotel --spec ../docs/swagger/hotels.yaml --principal interface{}

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
	if api.APIKeyAuth == nil {
		api.APIKeyAuth = func(token string) (interface{}, error) {
			return nil, errors.NotImplemented("api key auth (api_key) api_key from header param [api_key] has not yet been implemented")
		}
	}

	// Set your custom authorizer if needed. Default one is security.Authorized()
	// Expected interface runtime.Authorizer
	//
	// Example:
	// api.APIAuthorizer = security.Authorized()

	if api.HotelCreateHotelHandler == nil {
		api.HotelCreateHotelHandler = hotel2.CreateHotelHandlerFunc(func(params hotel2.CreateHotelParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation hotel.CreateHotel has not yet been implemented")
		})
	}
	if api.RoomCreateRoomHandler == nil {
		api.RoomCreateRoomHandler = room2.CreateRoomHandlerFunc(func(params room2.CreateRoomParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation room.CreateRoom has not yet been implemented")
		})
	}
	if api.HotelDeleteHotelByIDHandler == nil {
		api.HotelDeleteHotelByIDHandler = hotel2.DeleteHotelByIDHandlerFunc(func(params hotel2.DeleteHotelByIDParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation hotel.DeleteHotelByID has not yet been implemented")
		})
	}
	if api.RoomDeleteRoomByIDHandler == nil {
		api.RoomDeleteRoomByIDHandler = room2.DeleteRoomByIDHandlerFunc(func(params room2.DeleteRoomByIDParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation room.DeleteRoomByID has not yet been implemented")
		})
	}
	if api.HotelGetHotelByIDHandler == nil {
		api.HotelGetHotelByIDHandler = hotel2.GetHotelByIDHandlerFunc(func(params hotel2.GetHotelByIDParams) middleware.Responder {
			return middleware.NotImplemented("operation hotel.GetHotelByID has not yet been implemented")
		})
	}
	if api.HotelGetHotelsHandler == nil {
		api.HotelGetHotelsHandler = hotel2.GetHotelsHandlerFunc(func(params hotel2.GetHotelsParams) middleware.Responder {
			return middleware.NotImplemented("operation hotel.GetHotels has not yet been implemented")
		})
	}
	if api.RoomGetRoomByIDHandler == nil {
		api.RoomGetRoomByIDHandler = room2.GetRoomByIDHandlerFunc(func(params room2.GetRoomByIDParams) middleware.Responder {
			return middleware.NotImplemented("operation room.GetRoomByID has not yet been implemented")
		})
	}
	if api.RoomGetRoomsHandler == nil {
		api.RoomGetRoomsHandler = room2.GetRoomsHandlerFunc(func(params room2.GetRoomsParams) middleware.Responder {
			return middleware.NotImplemented("operation room.GetRooms has not yet been implemented")
		})
	}
	if api.HotelUpdateHotelHandler == nil {
		api.HotelUpdateHotelHandler = hotel2.UpdateHotelHandlerFunc(func(params hotel2.UpdateHotelParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation hotel.UpdateHotel has not yet been implemented")
		})
	}
	if api.RoomUpdateRoomHandler == nil {
		api.RoomUpdateRoomHandler = room2.UpdateRoomHandlerFunc(func(params room2.UpdateRoomParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation room.UpdateRoom has not yet been implemented")
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
