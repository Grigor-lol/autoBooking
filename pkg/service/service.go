package service

import (
	"autoBron"
	"time"
)

type Repository interface {
	CheckAvailability(id int, period *autoBron.AvailabilityPeriod) (bool, error)
	CreateBooking(booking *autoBron.BookingRequest) error
	GetCarUsageReport(startDate, endDate time.Time) ([]autoBron.CarUsage, error)

	HasBufferPeriod(carID int, startDate, endDate time.Time) (bool, error)
}

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{
		repo: repo,
	}
}
