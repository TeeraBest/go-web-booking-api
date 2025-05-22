package usecase

import (
	"context"
	"go-web-api/config"
	"go-web-api/domain/model"
	"go-web-api/domain/repository"
	"go-web-api/pkg/logging"
	"go-web-api/usecase/dto"
)

type BookingUsecase struct {
	logger     logging.Logger
	cfg        *config.Config
	repository repository.BookingRepository
}

func NewBookingUsecase(cfg *config.Config, repository repository.BookingRepository) *BookingUsecase {
	logger := logging.NewLogger(cfg)
	return &BookingUsecase{
		cfg:        cfg,
		repository: repository,
		logger:     logger,
	}
}

// Create a new booking and enqueue for async processing
func (u *BookingUsecase) Create(ctx context.Context, req dto.CreateBooking) (*dto.Booking, error) {
	booking := model.Booking{
		UserID:  req.UserID,
		EventID: req.EventID,
		Seat:    req.Seat,
		Status:  "pending",
	}
	created, err := u.repository.Create(ctx, booking)
	if err != nil {
		u.logger.Error(logging.General, "BookingCreate", err.Error(), nil)
		return nil, err
	}
	return &dto.Booking{
		ID:      created.ID,
		UserID:  created.UserID,
		EventID: created.EventID,
		Seat:    created.Seat,
		Status:  created.Status,
	}, nil
}

// Process the next booking in the queue (simulate async processing, e.g., by polling DB for pending bookings)
func (u *BookingUsecase) ProcessNext(ctx context.Context) (*dto.Booking, error) {
	// Example: get the first pending booking
	var booking model.Booking
	err := u.repository.FindPending(ctx, &booking)
	if err != nil {
		return nil, err
	}
	return &dto.Booking{
		ID:      booking.ID,
		UserID:  booking.UserID,
		EventID: booking.EventID,
		Seat:    booking.Seat,
		Status:  booking.Status,
	}, nil
}

// Set the status of a booking
func (u *BookingUsecase) SetStatus(ctx context.Context, bookingID, status string) error {
	return u.repository.UpdateField(ctx, bookingID, "status", status)
}

// Get the status of a booking
func (u *BookingUsecase) GetStatus(ctx context.Context, bookingID string) (string, error) {
	booking, err := u.repository.GetById(ctx, bookingID)
	if err != nil {
		return "", err
	}
	return booking.Status, nil
}

// CreateTicketForBooking creates a ticket for a booking and returns error if any
func (u *BookingUsecase) CreateTicketForBooking(ctx context.Context, bookingID string) error {
	booking, err := u.repository.GetById(ctx, bookingID)
	if err != nil {
		u.logger.Error(logging.General, "BookingFetchForTicket", err.Error(), nil)
		return err
	}
	ticket := model.Ticket{
		BookingID: booking.ID,
		UserID:    booking.UserID,
		EventID:   booking.EventID,
		Seat:      booking.Seat,
		Status:    "issued",
	}
	if err := u.repository.CreateTicket(ctx, ticket); err != nil {
		u.logger.Error(logging.General, "TicketCreate", err.Error(), nil)
		return err
	}
	return nil
}
