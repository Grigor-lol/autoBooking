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

type Report struct {
}
