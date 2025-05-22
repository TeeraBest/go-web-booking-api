package handler

import (
	"context"
	"go-web-api/api/dto"
	"go-web-api/api/helper"
	"go-web-api/config"
	"go-web-api/dependency"
	"go-web-api/usecase"
	usecasedto "go-web-api/usecase/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

type BookingHandler struct {
	usecase *usecase.BookingUsecase
}

func NewBookingHandler(cfg *config.Config) *BookingHandler {
	usecase := usecase.NewBookingUsecase(cfg, dependency.GetBookingRepository(cfg))
	return &BookingHandler{usecase: usecase}
}

// BookTicket godoc
// @Summary Book a ticket (async)
// @Description Enqueue a booking request
// @Tags Booking
// @Accept  json
// @Produce  json
// @Param Request body dto.BookingRequest true "BookingRequest"
// @Success 202 {object} helper.BaseHttpResponse "Accepted"
// @Failure 400 {object} helper.BaseHttpResponse "Failed"
// @Router /v1/bookings [post]
func (h *BookingHandler) BookTicket(c *gin.Context) {
	var req dto.BookingRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, helper.GenerateBaseResponseWithValidationError(nil, false, helper.ValidationError, err))
		return
	}
	userID := c.GetString("user_id")
	createReq := usecasedto.CreateBooking{
		UserID:  userID,
		EventID: req.EventID,
		Seat:    req.Seat,
	}
	booking, err := h.usecase.Create(context.Background(), createReq)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, helper.GenerateBaseResponseWithError(nil, false, helper.InternalError, err))
		return
	}
	c.JSON(http.StatusAccepted, helper.GenerateBaseResponse(gin.H{"booking_id": booking.ID, "status": booking.Status}, true, helper.Success))
}

// GetBookingStatus godoc
// @Summary Get booking status
// @Description Get status of a booking
// @Tags Booking
// @Accept  json
// @Produce  json
// @Param id path string true "Booking ID"
// @Success 200 {object} dto.BookingStatusResponse
// @Router /v1/bookings/{id}/status [get]
func (h *BookingHandler) GetBookingStatus(c *gin.Context) {
	id := c.Param("id")
	status, err := h.usecase.GetStatus(context.Background(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	}
	c.JSON(http.StatusOK, dto.BookingStatusResponse{BookingID: id, Status: status})
}
