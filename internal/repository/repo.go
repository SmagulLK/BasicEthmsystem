package repository

import (
	"TestProjectEthereum/models"
	postgres "TestProjectEthereum/pkg/database/postgresql"
	"context"
	"go.uber.org/zap"
	"math/big"
)

type Repository struct {
	Operation
	Generation
	CommonIn
}

func NewRepository(db *postgres.Postgres, logger *zap.Logger) *Repository {
	return &Repository{
		Operation:  NewOperationRepository(db, logger),
		Generation: NewGenRepository(db, logger),
	}
}

type CommonIn interface {
	InsertData(ctx context.Context, user *models.User) error
}
type Operation interface {
	GetUserByAddress(ctx context.Context, address string) (*models.User, error)
	BalanceUpdate(ctx context.Context, value big.Int) error
}
type Generation interface {
}
