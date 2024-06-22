package service

import "autoBron"

type Repository interface {
	CheckAvailability(id int) (bool, error)
	CreateBooking(booking *autoBron.BookingRequest) error
	GenerateReport() (*autoBron.Report, error)
}

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{
		repo: repo,
	}
}
