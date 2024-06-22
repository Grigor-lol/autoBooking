package handler

import (
	"autoBron"
	"github.com/gin-gonic/gin"
)

type Service interface {
	CheckAvailability(id int) (bool, error)
	CreateBooking(booking *autoBron.BookingRequest) error
	GenerateReport() (*autoBron.Report, error)
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
	router.POST("/book", h.CreateBooking)
	router.GET("/report", h.GenerateReport)

	return router
}
