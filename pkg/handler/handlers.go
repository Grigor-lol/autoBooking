package handler

import (
	"autoBron"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
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

	logrus.Info(req.StartDate)

	err := h.service.CreateBooking(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "Booking created"})
}

func (h *Handler) GenerateReport(c *gin.Context) {
	month, err := strconv.Atoi(c.Param("month"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Invalid month parameter: %s", err.Error())})
		return
	}

	if month < 1 || month > 12 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid month, should be from 1 to 12"})
		return
	}

	year, err := strconv.Atoi(c.Param("year"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Invalid year parameter: %s", err.Error())})
		return
	}

	if year < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid year"})
		return
	}

	usageReports, averageUsage, err := h.service.GenerateReport(uint8(month), year)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"reports":       usageReports,
		"average_usage": averageUsage,
	})
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
