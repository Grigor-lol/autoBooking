package postgres

import (
	"autoBron"
	"github.com/jmoiron/sqlx"
	"time"
)

type PostgreDB struct {
	db *sqlx.DB
}

func NewPostgreSql(db *sqlx.DB) *PostgreDB {
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

func (r *PostgreDB) GetCarUsageReport(startDate, endDate time.Time) ([]autoBron.CarUsage, error) {
	query := `
	SELECT 
		c.id as car_id,
		c.car_number as car_number,
		SUM(
			LEAST($2, b.end_date) - GREATEST($1, b.start_date) + 1
		) as days_rented
	FROM 
		cars c
	LEFT JOIN 
		bookings b ON c.id = b.car_id AND b.start_date <= $2 AND b.end_date >= $1
	GROUP BY 
		c.id, c.car_number
	`
	var reports []autoBron.CarUsage
	err := r.db.Select(&reports, query, startDate, endDate)
	return reports, err
}

func (r *PostgreDB) HasBufferPeriod(carID int, startDate, endDate time.Time) (bool, error) {
	var count int

	query := `
	SELECT COUNT(*) 
	FROM bookings 
	WHERE car_id = $1 
	AND (
		(end_date >= ($2::timestamp - interval '3 days' )) AND 
		(start_date <= ($3::timestamp + interval '3 days'))
	)
	`

	err := r.db.Get(&count, query, carID, startDate, endDate)
	if err != nil {
		return false, err
	}
	return count == 0, nil
}
