package repository

import (
	"go.uber.org/zap"

	"TestProjectEthereum/models"
	postgres "TestProjectEthereum/pkg/database/postgresql"
)

type Repository struct {
	Operation
}

func NewRepository(db *postgres.Postgres, logger *zap.Logger) *Repository {
	return &Repository{
		Operation: NewOperationRepository(db, logger),
	}
}

type Operation interface {
	GetUserbyId(id int) (models.User, error)
	BalanceUpdate(value int32) error
}
