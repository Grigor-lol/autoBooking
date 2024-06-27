package repository

import (
	"autoBron/internal/repository/postgres"
	"autoBron/pkg/autoBooking"
	"github.com/jmoiron/sqlx"
	"time"
)

type Repository interface {
	CheckAvailability(id int, period *autoBooking.AvailabilityPeriod) (bool, error)
	CreateBooking(id int, period *autoBooking.AvailabilityPeriod) error
	GetCarUsageReport(startDate, endDate time.Time) ([]autoBooking.CarUsage, error)

	HasBufferPeriod(carID int, startDate, endDate time.Time) (bool, error)
}

func NewRepository(db *sqlx.DB) Repository {
	return postgres.NewPostgreSql(db)
}
