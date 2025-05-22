package dependency

import (
	"go-web-api/config"
	contractRepository "go-web-api/domain/repository"
	infraRepository "go-web-api/infra/persistence/repository"
	"go-web-api/usecase"
)

func GetUserRepository(cfg *config.Config) contractRepository.UserRepository {
	return infraRepository.NewUserRepository(cfg)
}

func GetBookingRepository(cfg *config.Config) contractRepository.BookingRepository {
	return infraRepository.NewBookingRepository(cfg)
}

func GetBookingUsecase(cfg *config.Config) *usecase.BookingUsecase {
	return usecase.NewBookingUsecase(cfg, GetBookingRepository(cfg))
}
