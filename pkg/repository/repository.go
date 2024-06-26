package repository

import (
	"autoBron"
	"autoBron/pkg/repository/postgres"
	"github.com/jmoiron/sqlx"
	"time"
)

type Repository interface {
	CheckAvailability(id int, period *autoBron.AvailabilityPeriod) (bool, error)
	CreateBooking(booking *autoBron.BookingRequest) error
	GetCarUsageReport(startDate, endDate time.Time) ([]autoBron.CarUsage, error)

	HasBufferPeriod(carID int, startDate, endDate time.Time) (bool, error)
}

func NewRepository(db *sqlx.DB) Repository {
	return postgres.NewPostgreSql(db)
}
