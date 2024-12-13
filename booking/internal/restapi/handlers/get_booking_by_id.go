package handlers

import (
	"context"
	"fmt"
	"github.com/go-openapi/runtime/middleware"
	"github.com/h4x4d/go_hsse_hotels/booking/internal/models"
	"github.com/h4x4d/go_hsse_hotels/booking/internal/restapi/operations/customer"
	"github.com/h4x4d/go_hsse_hotels/booking/internal/utils"
	"google.golang.org/grpc/metadata"
)

func (handler *Handler) GetBookingByID(params customer.GetBookingByIDParams, user *models.User) (responder middleware.Responder) {
	defer utils.CatchPanic(&responder)

	ctx, span := handler.tracer.Start(context.Background(), "get booking by id")
	defer span.End()
	traceId := fmt.Sprintf("%s", span.SpanContext().TraceID())
	ctx = metadata.AppendToOutgoingContext(ctx, "x-trace-id", traceId)

	owner, err := handler.Database.CheckOwnership(ctx, params.BookingID, user)
	if err != nil {
		return utils.HandleInternalError(err)
	}
	if !owner {
		errCode := int64(customer.GetBookingByIDForbiddenCode)
		result := new(customer.GetBookingByIDForbidden)
		result.SetPayload(&models.Error{
			ErrorMessage:    "You don't have permission to get this booking",
			ErrorStatusCode: &errCode,
		})
		return result
	}
	booking, errGet := handler.Database.GetByID(params.BookingID)
	if errGet != nil {
		return utils.HandleInternalError(errGet)
	}

	result := new(customer.GetBookingByIDOK)
	result.SetPayload(booking)
	return result
}
