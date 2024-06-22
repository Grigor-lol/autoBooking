package service

import "autoBron"

func (s *Service) CheckAvailability(id int) (bool, error) {
	return s.repo.CheckAvailability(id)
}

func (s *Service) CreateBooking(booking *autoBron.BookingRequest) error {
	return s.repo.CreateBooking(booking)
}

func (s *Service) GenerateReport() (*autoBron.Report, error) {
	return s.repo.GenerateReport()
}
