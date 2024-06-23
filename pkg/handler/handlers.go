package handler

import (
	"autoBron"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) CheckAvailability(c *gin.Context) {
	carID, err := strconv.Atoi(c.Param("car_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Invalid car ID: %s", err.Error())})
		return
	}

	body := new(autoBron.AvailabilityPeriod)
	err = c.ShouldBindJSON(body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Invalid request body: %s", err.Error())})
		return
	}

	available, err := h.service.CheckAvailability(carID, body)
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

func (h *Handler) CalculateRentalCost(c *gin.Context) {
	carID, err := strconv.Atoi(c.Param("car_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Invalid car ID: %s", err.Error())})
		return
	}

	req := new(autoBron.AvailabilityPeriod)
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	cost, err := h.service.CalculateRentalCost(carID, req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"cost": cost})
}
