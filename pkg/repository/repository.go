package repository

import (
	"autoBron"
	"github.com/jmoiron/sqlx"
)

type Database interface {
	CheckAvailability(id int) (bool, error)
	CreateBooking(booking *autoBron.BookingRequest) error
	GenerateReport() (*autoBron.Report, error)
}

type Repository struct {
	Database
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Database: newPostgreSql(db),
	}
}
