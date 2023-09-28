package repository

import (
	postgres "TestProjectEthereum/pkg/database/postgresql"
	"go.uber.org/zap"
)

type GenRepository struct {
	Common
	db     *postgres.Postgres
	logger *zap.Logger
}

func NewGenRepository(db *postgres.Postgres, logger *zap.Logger) *GenRepository {
	return &GenRepository{db: db, logger: logger}
}
