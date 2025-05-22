package repository

import (
	"context"
	"errors"
	"go-web-api/config"
	"go-web-api/domain/filter"
	model "go-web-api/domain/model"
	database "go-web-api/infra/persistence/database"
	"time"

	"gorm.io/gorm"
)

type PostgresBookingRepository struct {
	*BaseRepository[model.Booking]
}

func NewBookingRepository(cfg *config.Config) *PostgresBookingRepository {
	var preloads []database.PreloadEntity = []database.PreloadEntity{}
	return &PostgresBookingRepository{BaseRepository: NewBaseRepository[model.Booking](cfg, preloads)}
}

// You can add custom methods for booking here if needed

// Create a new booking
func (r *PostgresBookingRepository) Create(ctx context.Context, entity model.Booking) (model.Booking, error) {
	err := r.BaseRepository.database.WithContext(ctx).Create(&entity).Error
	return entity, err
}

// Update a booking by id
func (r *PostgresBookingRepository) Update(ctx context.Context, id int, entity map[string]interface{}) (model.Booking, error) {
	var booking model.Booking
	err := r.BaseRepository.database.WithContext(ctx).Model(&booking).Where("id = ?", id).Updates(entity).Error
	if err != nil {
		return booking, err
	}
	err = r.BaseRepository.database.WithContext(ctx).First(&booking, "id = ?", id).Error
	return booking, err
}

// Delete a booking by id
func (r *PostgresBookingRepository) Delete(ctx context.Context, id int) error {
	return r.BaseRepository.database.WithContext(ctx).Delete(&model.Booking{}, "id = ?", id).Error
}

// Get a booking by string id
func (r *PostgresBookingRepository) GetById(ctx context.Context, id string) (model.Booking, error) {
	var booking model.Booking
	err := r.BaseRepository.database.WithContext(ctx).First(&booking, "id = ?", id).Error
	return booking, err
}

// Get bookings by filter
func (r *PostgresBookingRepository) GetByFilter(ctx context.Context, req filter.PaginationInputWithFilter) (int64, *[]model.Booking, error) {
	var bookings []model.Booking
	var count int64
	db := r.BaseRepository.database.WithContext(ctx).Model(&model.Booking{})
	// You can add filter logic here if needed
	err := db.Count(&count).Find(&bookings).Error
	return count, &bookings, err
}

// Update a single field by string id
func (r *PostgresBookingRepository) UpdateField(ctx context.Context, id string, field string, value interface{}) error {
	return r.BaseRepository.database.WithContext(ctx).Model(&model.Booking{}).Where("id = ?", id).Update(field, value).Error
}

// FindPending finds the first booking with status 'pending'
func (r *PostgresBookingRepository) FindPending(ctx context.Context, booking *model.Booking) error {
	for {
		err := r.BaseRepository.database.WithContext(ctx).Where("status = ?", "pending").First(booking).Error
		if err == nil {
			return nil
		}
		if errors.Is(err, gorm.ErrRecordNotFound) {
			time.Sleep(time.Second)
			continue
		}
		return err
	}
}

// CreateTicket creates a new ticket for a booking
func (r *PostgresBookingRepository) CreateTicket(ctx context.Context, ticket model.Ticket) error {
	return r.BaseRepository.database.WithContext(ctx).Create(&ticket).Error
}
