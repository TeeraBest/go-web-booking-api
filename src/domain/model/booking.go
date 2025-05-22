package model

// Booking represents a ticket booking request
// Add more fields as needed for your business logic

type Booking struct {
	ID      string `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	UserID  string `json:"user_id"`
	EventID int    `json:"event_id"`
	Seat    string `json:"seat"`
	Status  string `json:"status"` // e.g. pending, confirmed, failed
}
