package service

import (
	"go-crud-evo/internal/repository"
)

type NumberService interface {
	ProcessNumber(number int) ([]int, error)
}

type DefaultNumberService struct {
	repo repository.NumberRepository
}

func NewDefaultNumberService(repo repository.NumberRepository) *DefaultNumberService {
	return &DefaultNumberService{repo: repo}
}

func (s *DefaultNumberService) ProcessNumber(number int) ([]int, error) {
	// Business logic: Save the number
	if err := s.repo.Save(number); err != nil {
		return nil, err
	}

	// Business logic: Retrieve all numbers (sorting is handled by repo/DB in this case)
	return s.repo.GetAll()
}
