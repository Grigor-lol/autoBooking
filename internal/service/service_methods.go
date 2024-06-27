package service

import (
	"autoBron/pkg/autoBooking"
	"errors"
	"time"
)

func (s *Service) CheckAvailability(id int, period *autoBooking.AvailabilityPeriod) (bool, error) {
	return s.repo.CheckAvailability(id, period)
}

func (s *Service) CreateBooking(id int, period *autoBooking.AvailabilityPeriod) error {
	err := validateRentalPeriod(period)
	if err != nil {
		return err
	}

	available, err := s.repo.CheckAvailability(id, period)
	if err != nil {
		return err
	}
	if !available {
		return errors.New("car is not available for the selected period")
	}

	// Проверка: интервал между бронированиями должен быть не менее 3 дней
	hasBuffer, err := s.repo.HasBufferPeriod(id, period.StartDate.Time, period.EndDate.Time)
	if err != nil {
		return err
	}
	if !hasBuffer {
		return errors.New("there must be a buffer of 3 days between bookings")
	}

	return s.repo.CreateBooking(id, period)
}

func (s *Service) GenerateReport(month uint8, year int) ([]autoBooking.CarUsageReport, float64, error) {
	startDate := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.UTC)
	endDate := startDate.AddDate(0, 1, -1)

	reports, err := s.repo.GetCarUsageReport(startDate, endDate)
	if err != nil {
		return nil, 0, err
	}

	// Определение общего количества рабочих дней в месяце
	workDays := countWorkdays(startDate, endDate)

	var totalUsage float64
	var usageReports []autoBooking.CarUsageReport

	for _, report := range reports {
		usageRate := float64(report.DaysRented) / float64(workDays) * 100
		usageReports = append(usageReports, autoBooking.CarUsageReport{
			CarNumber: report.CarNumber,
			UsageRate: usageRate,
		})
		totalUsage += usageRate
	}

	averageUsage := totalUsage / float64(autoBooking.CarCounts)
	return usageReports, averageUsage, nil
}

func (s *Service) CalculateRentalCost(id int, period *autoBooking.AvailabilityPeriod) (float32, error) {
	err := validateRentalPeriod(period)
	if err != nil {
		return 0, err
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
	days := int(period.EndDate.Sub(period.StartDate.Time).Hours()/24) + 1
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

func countWorkdays(startDate, endDate time.Time) int {
	workDays := 0
	for date := startDate; !date.After(endDate); date = date.AddDate(0, 0, 1) {
		if date.Weekday() != time.Saturday && date.Weekday() != time.Sunday {
			workDays += 1
		}
	}
	return workDays
}

func validateRentalPeriod(period *autoBooking.AvailabilityPeriod) error {
	if period.EndDate.Before(period.StartDate.Time) ||
		period.EndDate.Sub(period.StartDate.Time) > 30*24*time.Hour ||
		period.EndDate.Sub(period.StartDate.Time) < 24*time.Hour {
		return errors.New("rental period must be between 1 and 30 days")
	}

	// Проверка: начало и конец аренды - будние дни
	if period.StartDate.Weekday() == time.Saturday || period.StartDate.Weekday() == time.Sunday ||
		period.EndDate.Weekday() == time.Saturday || period.EndDate.Weekday() == time.Sunday {
		return errors.New("start and end date must be on weekdays")
	}

	return nil
}
