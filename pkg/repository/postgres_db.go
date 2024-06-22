package repository

import (
	"autoBron"
	"github.com/jmoiron/sqlx"
)

type PostgreDB struct {
	db *sqlx.DB
}

func newPostgreSql(db *sqlx.DB) *PostgreDB {
	return &PostgreDB{
		db: db,
	}
}

func (r *PostgreDB) CheckAvailability(id int) (bool, error) {
	return true, nil
}

func (r *PostgreDB) CreateBooking(booking *autoBron.BookingRequest) error {
	return nil
}

func (r *PostgreDB) GenerateReport() (*autoBron.Report, error) {
	return nil, nil
}
