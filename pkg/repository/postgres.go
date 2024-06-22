package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Config struct {
	Username string
	Password string
}

const (
	createBookingsTable = `
	CREATE TABLE IF NOT EXISTS bookings (
		id SERIAL PRIMARY KEY,
		car_id INT NOT NULL,
		start_date TIMESTAMP NOT NULL,
		end_date TIMESTAMP NOT NULL
	);
	`

	createCarsTable = `
	CREATE TABLE IF NOT EXISTS cars (
		id SERIAL PRIMARY KEY,
		car_number VARCHAR(20) NOT NULL
	);
	`
)

func NewPostgresDB(cfg Config) (*sqlx.DB, error) {
	connStr := fmt.Sprintf("user=%s password=%s dbname=postgres sslmode=disable", cfg.Username, cfg.Password)
	db, err := sqlx.Connect("postgres", connStr)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	err = createTables(db)
	return db, err
}

func createTables(conn *sqlx.DB) error {
	tx, err := conn.Begin()
	if err != nil {
		return err
	}

	_, err = tx.Exec(createCarsTable)
	if err != nil {
		tx.Rollback()
		return err
	}

	_, err = tx.Exec(createBookingsTable)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}
