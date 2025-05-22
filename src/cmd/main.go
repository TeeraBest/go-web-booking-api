package main

import (
	"context"
	"go-web-api/api"
	"go-web-api/config"
	"go-web-api/dependency"
	"go-web-api/infra/cache"
	database "go-web-api/infra/persistence/database"
	"go-web-api/infra/persistence/migration"
	"go-web-api/pkg/logging"
)

// @securityDefinitions.apikey AuthBearer
// @in header
// @name Authorization
func main() {

	cfg := config.GetConfig()
	logger := logging.NewLogger(cfg)

	err := cache.InitRedis(cfg)
	defer cache.CloseRedis()
	if err != nil {
		logger.Fatal(logging.Redis, logging.Startup, err.Error(), nil)
	}

	err = database.InitDb(cfg)
	defer database.CloseDb()
	if err != nil {
		logger.Fatal(logging.Postgres, logging.Startup, err.Error(), nil)
	}
	migration.Up1()

	go func() {
		bookingUsecase := dependency.GetBookingUsecase(cfg)
		ctx := context.Background()
		for {
			booking, err := bookingUsecase.ProcessNext(ctx)
			if err == nil && booking != nil {
				// Create the ticket in the background
				err := bookingUsecase.CreateTicketForBooking(ctx, booking.ID)
				if err == nil {
					_ = bookingUsecase.SetStatus(ctx, booking.ID, "confirmed")
				}
			}
		}
	}()

	api.InitServer(cfg)
}
