package service

import (
	"context"
	"math/big"

	"go.uber.org/zap"

	"TestProjectEthereum/internal/repository"
	ethereum "TestProjectEthereum/pkg/blockchain/etherium"
)

type Service struct {
	OperationServiceIn
	GenerationIn
	OperationService
}

func NewService(repository *repository.Repository, logger *zap.Logger, etherumURL string) (*Service, error) {

	etheriumInstance, err := ethereum.NewEthereumClient(etherumURL)
	if err != nil {
		return &Service{}, err
	}
	return &Service{
		OperationServiceIn: NewOperationService(repository, logger, etheriumInstance),
		GenerationIn:       NewGenService(repository, logger),
		OperationService:   *NewOperationService(repository, logger, etheriumInstance),
	}, nil
}

type OperationServiceIn interface {
	UpdateBalance(ctx context.Context, balance big.Int) error
}
type GenerationIn interface {
	Generate(ctx context.Context) (string, string, string, error)
}
