package repository

import (
	"go.uber.org/zap"

	"TestProjectEthereum/models"
	postgres "TestProjectEthereum/pkg/database/postgresql"
)

type OperationRepository struct {
	db     *postgres.Postgres
	logger *zap.Logger
}

func NewOperationRepository(db *postgres.Postgres, logger *zap.Logger) *OperationRepository {
	return &OperationRepository{db: db, logger: logger}
}
func (Op *OperationRepository) GetUserbyId(id int) (models.User, error) {
	return models.User{}, nil
}
func (Op *OperationRepository) BalanceUpdate(balance int32) error {
	return nil
}
