package repository

import (
	"autoBron"
	"github.com/jmoiron/sqlx"
	"time"
)

type Database interface {
	CheckAvailability(id int, period *autoBron.AvailabilityPeriod) (bool, error)
	CreateBooking(booking *autoBron.BookingRequest) error
	GenerateReport() (*autoBron.Report, error)

	HasBufferPeriod(carID int, startDate, endDate time.Time) (bool, error)
}

type Repository struct {
	Database
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Database: newPostgreSql(db),
	}
}
