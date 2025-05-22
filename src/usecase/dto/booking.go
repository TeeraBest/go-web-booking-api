package dto

type CreateBooking struct {
	EventID int    `json:"event_id" binding:"required"`
	Seat    string `json:"seat" binding:"required"`
	UserID  string `json:"user_id"`
}

type UpdateBooking struct {
	Status string `json:"status"`
}

type Booking struct {
	ID      string `json:"id"`
	UserID  string `json:"user_id"`
	EventID int    `json:"event_id"`
	Seat    string `json:"seat"`
	Status  string `json:"status"`
}
