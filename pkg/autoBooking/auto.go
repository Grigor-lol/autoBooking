package autoBooking

import (
	"strings"
	"time"
)

const CarCounts = 5

type CustomDate struct {
	time.Time
}

func (c *CustomDate) UnmarshalJSON(b []byte) (err error) {
	layout := "2006-01-02 15:04:05"

	s := strings.Trim(string(b), "\"") // remove quotes
	if s == "null" {
		return
	}
	c.Time, err = time.Parse(layout, s)
	return
}

type BookingRequest struct {
	CarID     int        `json:"car_id" binding:"required"`
	StartDate CustomDate `json:"start_date" binding:"required"`
	EndDate   CustomDate `json:"end_date" binding:"required"`
}

type AvailabilityPeriod struct {
	StartDate CustomDate `json:"start_date" binding:"required"`
	EndDate   CustomDate `json:"end_date" binding:"required"`
}

type CarUsage struct {
	CarID      int     `db:"car_id"`
	CarNumber  string  `db:"car_number"`
	DaysRented float64 `db:"days_rented"`
}

type CarUsageReport struct {
	CarNumber string  `json:"car_number"`
	UsageRate float64 `json:"usage_rate"`
}
