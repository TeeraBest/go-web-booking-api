package model

type Ticket struct {
	ID        string `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	BookingID string `gorm:"type:uuid;not null" json:"booking_id"`
	UserID    string `json:"user_id"`
	EventID   int    `json:"event_id"`
	Seat      string `json:"seat"`
	Status    string `json:"status"`
}
