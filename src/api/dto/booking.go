package dto

type BookingRequest struct {
	EventID int    `json:"event_id" binding:"required"`
	Seat    string `json:"seat" binding:"required"`
}

type BookingStatusResponse struct {
	BookingID string `json:"booking_id"`
	Status    string `json:"status"`
}

type BookTicketRequest struct {
	
}
