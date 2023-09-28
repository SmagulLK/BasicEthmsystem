package service

import (
	"TestProjectEthereum/internal/repository"
	"context"
	"math/big"

	"go.uber.org/zap"
)

type Service struct {
	OperationServiceIn
	GenerationIn
}

func NewService(repository *repository.Repository, logger *zap.Logger) *Service {
	return &Service{
		OperationServiceIn: NewOperationService(repository, logger),
		GenerationIn:       NewGenService(repository, logger),
	}
}

type OperationServiceIn interface {
	UpdateBalance(ctx context.Context, balance big.Int) error
}
type GenerationIn interface {
	Generate(ctx context.Context) error
}
