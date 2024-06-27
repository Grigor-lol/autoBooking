package service

import (
	"autoBron/pkg/autoBooking"
	"time"
)

type Repository interface {
	CheckAvailability(id int, period *autoBooking.AvailabilityPeriod) (bool, error)
	CreateBooking(id int, period *autoBooking.AvailabilityPeriod) error
	GetCarUsageReport(startDate, endDate time.Time) ([]autoBooking.CarUsage, error)

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
