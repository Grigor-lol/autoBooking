package postgres

import (
	"autoBron/pkg/autoBooking"
	"errors"
	"fmt"
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

func (r *PostgreDB) CheckAvailability(id int, period *autoBooking.AvailabilityPeriod) (bool, error) {
	var exists bool
	checkCarQuery := `SELECT EXISTS(SELECT 1 FROM cars WHERE id = $1)`
	err := r.db.Get(&exists, checkCarQuery, id)
	if err != nil {
		return false, err
	}
	if !exists {
		return false, errors.New("no cars with this id") // Автомобиль с данным ID не существует
	}

	var count int
	query := `
	SELECT COUNT(*) 
	FROM bookings 
	WHERE car_id = $1 
	AND ($2, $3) OVERLAPS (start_date, end_date)
	`

	err = r.db.Get(&count, query, id, period.StartDate.Time, period.EndDate.Time)
	if err != nil {
		return false, err
	}
	return count == 0, nil
}

func (r *PostgreDB) CreateBooking(id int, period *autoBooking.AvailabilityPeriod) error {
	query := `
	INSERT INTO bookings (car_id, start_date, end_date)
	VALUES ($1, $2, $3)
	`

	_, err := r.db.Exec(query, id, period.StartDate.Time, period.EndDate.Time)
	return err
}

func (r *PostgreDB) GetCarUsageReport(startDate, endDate time.Time) ([]autoBooking.CarUsage, error) {
	query := `
	SELECT 
		c.id as car_id,
		c.car_number as car_number,
		COALESCE(
		SUM(
			EXTRACT(EPOCH FROM (LEAST($2, b.end_date) - GREATEST($1, b.start_date))) / 86400
		  ),0
		) as days_rented
	FROM 
		cars c
	JOIN 
		bookings b ON c.id = b.car_id AND b.start_date <= $2 AND b.end_date >= $1
	GROUP BY 
		c.id, c.car_number
	`
	var reports []autoBooking.CarUsage
	err := r.db.Select(&reports, query, startDate, endDate)
	for _, r := range reports {
		fmt.Println(r.CarID, r.DaysRented)
	}
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
