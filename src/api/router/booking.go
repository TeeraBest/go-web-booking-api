package router

import (
	"go-web-api/api/handler"
	"go-web-api/config"

	"github.com/gin-gonic/gin"
)

// Booking godoc
func Booking(r *gin.RouterGroup, cfg *config.Config) {
	h := handler.NewBookingHandler(cfg)

	r.POST("/bookings", h.BookTicket)
	r.GET("/bookings/:id/status", h.GetBookingStatus)
}
