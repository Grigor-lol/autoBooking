package repository

import (
	"autoBron"
	"github.com/jmoiron/sqlx"
	"time"
)

type PostgreDB struct {
	db *sqlx.DB
}

func newPostgreSql(db *sqlx.DB) *PostgreDB {
	return &PostgreDB{
		db: db,
	}
}

func (r *PostgreDB) CheckAvailability(id int, period *autoBron.AvailabilityPeriod) (bool, error) {
	var count int

	query := `
	SELECT COUNT(*) 
	FROM bookings 
	WHERE car_id = $1 
	AND ($2, $3) OVERLAPS (start_date, end_date)
	`

	err := r.db.Get(&count, query, id, period.StartDate, period.EndDate)
	if err != nil {
		return false, err
	}
	return count == 0, nil
}

func (r *PostgreDB) CreateBooking(booking *autoBron.BookingRequest) error {
	query := `
	INSERT INTO bookings (car_id, start_date, end_date)
	VALUES ($1, $2, $3)
	`

	_, err := r.db.Exec(query, booking.CarID, booking.StartDate, booking.EndDate)
	return err
}

func (r *PostgreDB) GenerateReport() (*autoBron.Report, error) {
	return nil, nil
}

func (r *PostgreDB) HasBufferPeriod(carID int, startDate, endDate time.Time) (bool, error) {
	var count int

	query := `
	SELECT COUNT(*) 
	FROM bookings 
	WHERE car_id = $1 
	AND (
		(start_date <= $2 AND end_date < $2 - interval '3 days') OR 
		(start_date > $3 + interval '3 days' AND end_date >= $3)
	)
	`

	err := r.db.Get(&count, query, carID, startDate, endDate)
	if err != nil {
		return false, err
	}
	return count == 0, nil
}
