package autoBron

import "time"

type BookingRequest struct {
	CarID     int       `json:"car_id" binding:"required"`
	StartDate time.Time `json:"start_date" binding:"required"`
	EndDate   time.Time `json:"end_date" binding:"required"`
}

type AvailabilityPeriod struct {
	StartDate time.Time `json:"start_date" binding:"required"`
	EndDate   time.Time `json:"end_date" binding:"required"`
}

type CarUsage struct {
	CarID      int    `db:"car_id"`
	CarNumber  string `db:"car_number"`
	DaysRented int    `db:"days_rented"`
}

type CarUsageReport struct {
	CarNumber string  `json:"car_number"`
	UsageRate float64 `json:"usage_rate"`
}
