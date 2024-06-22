package autoBron

type BookingRequest struct {
	CarID     int   `json:"car_id" binding:"required"`
	StartDate int64 `json:"start_date" binding:"required"` // Unix timestamp in seconds
	EndDate   int64 `json:"end_date" binding:"required"`   // Unix timestamp in seconds
}

type Report struct {
}
