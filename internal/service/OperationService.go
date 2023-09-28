package service

import (
	"TestProjectEthereum/internal/repository"
	"context"
	"go.uber.org/zap"
	"math/big"
)

type OperationService struct {
	logger *zap.Logger
	repo   *repository.Repository
}

func NewOperationService(repo *repository.Repository, logger *zap.Logger) *OperationService {
	return &OperationService{logger: logger, repo: repo}
}
func (OperationServ *OperationService) UpdateBalance(ctx context.Context, balance big.Int) error {
	return OperationServ.repo.BalanceUpdate(ctx, balance)
}
