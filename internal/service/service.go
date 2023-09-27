package service

import (
	"TestProjectEthereum/internal/repository"

	"go.uber.org/zap"
)

type Service struct {
	OperationServiceIn
}

func NewService(repository *repository.Repository, logger *zap.Logger) *Service {
	return &Service{
		OperationServiceIn: NewOperationService(repository),
	}
}

type OperationServiceIn interface {
	UpdateBalance(balance int32) error
}
