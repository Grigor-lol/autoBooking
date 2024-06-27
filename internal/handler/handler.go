package handler

import (
	"autoBron/pkg/autoBooking"
	"github.com/gin-gonic/gin"
)

type Service interface {
	CheckAvailability(id int, period *autoBooking.AvailabilityPeriod) (bool, error)
	CreateBooking(id int, period *autoBooking.AvailabilityPeriod) error
	GenerateReport(month uint8, year int) ([]autoBooking.CarUsageReport, float64, error)
	CalculateRentalCost(id int, period *autoBooking.AvailabilityPeriod) (float32, error)
}

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.GET("/availability/:car_id", h.CheckAvailability)
	router.GET("/cost/:car_id", h.CalculateRentalCost)
	router.POST("/book/:car_id", h.CreateBooking)
	router.GET("/report", h.GenerateReport)

	return router
}
