package repository

import (
	"TestProjectEthereum/models"
	postgres "TestProjectEthereum/pkg/database/postgresql"

	"go.uber.org/zap"
)

type Repository struct {
	Operation
}

func NewRepository(db *postgres.Postgres, logger *zap.Logger) *Repository {
	return &Repository{
		Operation: NewOperationRepository(db),
	}
}

type Operation interface {
	GetUserbyId(id int) (models.User, error)
	BalanceUpdate(value int32) error
}
