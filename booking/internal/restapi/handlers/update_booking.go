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

func (handler *Handler) UpdateBooking(params customer.UpdateBookingParams, user *models.User) (responder middleware.Responder) {
	defer utils.CatchPanic(&responder)
	ctx, span := handler.tracer.Start(context.Background(), "update booking")
	traceId := fmt.Sprintf("%s", span.SpanContext().TraceID())
	ctx = metadata.AppendToOutgoingContext(ctx, "x-trace-id", traceId)

	defer span.End()
	owner, err := handler.Database.CheckOwnership(ctx, params.BookingID, user)
	if err != nil {
		return utils.HandleInternalError(err)
	}
	if !owner {
		errCode := int64(customer.GetBookingByIDForbiddenCode)
		result := new(customer.GetBookingByIDForbidden)
		result.SetPayload(&models.Error{
			ErrorMessage:    "You don't have permission to update this booking",
			ErrorStatusCode: &errCode,
		})
		return result
	}
	booking, errUpdate := handler.Database.Update(ctx, params.BookingID, params.Object)
	if errUpdate != nil {
		return utils.HandleInternalError(errUpdate)
	}

	result := new(customer.UpdateBookingOK)
	result.SetPayload(booking)
	return result
}
