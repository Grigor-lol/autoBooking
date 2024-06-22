package handler

import (
	"autoBron"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) CheckAvailability(c *gin.Context) {
	carID, err := strconv.Atoi(c.Param("car_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid car ID"})
		return
	}

	available, err := h.service.CheckAvailability(carID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"available": available})
}

func (h *Handler) CreateBooking(c *gin.Context) {
	req := new(autoBron.BookingRequest)
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.service.CreateBooking(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "Booking created"})
}

func (h *Handler) GenerateReport(c *gin.Context) {
	report, err := h.service.GenerateReport()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, report)
}
