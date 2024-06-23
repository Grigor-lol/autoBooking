package service

import (
	"autoBron"
	"errors"
	"time"
)

func (s *Service) CheckAvailability(id int, period *autoBron.AvailabilityPeriod) (bool, error) {
	return s.repo.CheckAvailability(id, period)
}

func (s *Service) CreateBooking(booking *autoBron.BookingRequest) error {
	// Проверка: срок аренды от 1 до 30 дней
	if booking.EndDate.Before(booking.StartDate) || booking.EndDate.Sub(booking.StartDate).Hours()/24 > 30 {
		return errors.New("rental period must be between 1 and 30 days")
	}

	// Проверка: начало и конец аренды - будние дни
	if booking.StartDate.Weekday() == time.Saturday || booking.StartDate.Weekday() == time.Sunday ||
		booking.EndDate.Weekday() == time.Saturday || booking.EndDate.Weekday() == time.Sunday {
		return errors.New("start and end date must be on weekdays")
	}

	// Проверка: доступность автомобиля
	period := &autoBron.AvailabilityPeriod{StartDate: booking.StartDate, EndDate: booking.EndDate}
	available, err := s.repo.CheckAvailability(booking.CarID, period)
	if err != nil {
		return err
	}
	if !available {
		return errors.New("car is not available for the selected period")
	}

	// Проверка: интервал между бронированиями должен быть не менее 3 дней
	hasBuffer, err := s.repo.HasBufferPeriod(booking.CarID, booking.StartDate, booking.EndDate)
	if err != nil {
		return err
	}
	if !hasBuffer {
		return errors.New("there must be a buffer of 3 days between bookings")
	}

	return s.repo.CreateBooking(booking)
}

func (s *Service) GenerateReport() (*autoBron.Report, error) {
	return s.repo.GenerateReport()
}

func (s *Service) CalculateRentalCost(id int, period *autoBron.AvailabilityPeriod) (float32, error) {
	// Проверка: срок аренды от 1 до 30 дней
	if period.EndDate.Before(period.StartDate) || period.EndDate.Sub(period.StartDate).Hours()/24 > 30 {
		return 0, errors.New("rental period must be between 1 and 30 days")
	}

	// Проверка: начало и конец аренды - будние дни
	if period.StartDate.Weekday() == time.Saturday || period.StartDate.Weekday() == time.Sunday ||
		period.EndDate.Weekday() == time.Saturday || period.EndDate.Weekday() == time.Sunday {
		return 0, errors.New("start and end date must be on weekdays")
	}

	// Проверка: доступность автомобиля
	available, err := s.repo.CheckAvailability(id, period)
	if err != nil {
		return 0, err
	}
	if !available {
		return 0, errors.New("car is not available for the selected period")
	}

	// Расчет стоимости аренды
	days := int(period.EndDate.Sub(period.StartDate).Hours()/24) + 1
	cost := float32(0.0)
	baseRate := float32(1000.0)

	for i := 1; i <= days; i++ {
		if i <= 4 {
			cost += baseRate
		} else if i <= 9 {
			cost += baseRate * 0.95
		} else if i <= 17 {
			cost += baseRate * 0.90
		} else if i <= 29 {
			cost += baseRate * 0.85
		} else {
			cost += baseRate * 0.85
		}
	}

	return cost, nil
}
