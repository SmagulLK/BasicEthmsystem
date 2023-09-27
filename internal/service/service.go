package service

import "TestProjectEthereum/internal/repository"

type Service struct {
	OperationServiceIn
}

func NewService(repository *repository.Repository) *Service {
	return &Service{
		OperationServiceIn: NewOperationService(repository),
	}
}

type OperationServiceIn interface {
	UpdateBalance(balance int32) error
}
