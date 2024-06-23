package handler

import (
	"autoBron"
	"github.com/gin-gonic/gin"
)

type Service interface {
	CheckAvailability(id int, period *autoBron.AvailabilityPeriod) (bool, error)
	CreateBooking(booking *autoBron.BookingRequest) error
	GenerateReport() (*autoBron.Report, error)
	CalculateRentalCost(id int, period *autoBron.AvailabilityPeriod) (float32, error)
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
	router.POST("/book", h.CreateBooking)
	router.GET("/report", h.GenerateReport)

	return router
}
