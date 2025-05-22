package repository

import (
	"context"
	"go-web-api/domain/filter"
	"go-web-api/domain/model"
)

type BookingRepository interface {
	Create(ctx context.Context, entity model.Booking) (model.Booking, error)
	Update(ctx context.Context, id int, entity map[string]interface{}) (model.Booking, error)
	Delete(ctx context.Context, id int) error
	GetById(ctx context.Context, id string) (model.Booking, error)
	GetByFilter(ctx context.Context, req filter.PaginationInputWithFilter) (int64, *[]model.Booking, error)
	UpdateField(ctx context.Context, id string, field string, value interface{}) error
	FindPending(ctx context.Context, booking *model.Booking) error
	CreateTicket(ctx context.Context, ticket model.Ticket) error
}
