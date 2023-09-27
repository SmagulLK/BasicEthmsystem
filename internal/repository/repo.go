package repository

import (
	"TestProjectEthereum/models"
	"github.com/jmoiron/sqlx"
)

type Repository struct {
	Operation
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Operation: NewOperationRepository(db),
	}
}

type Operation interface {
	GetUserbyId(id int) (models.User, error)
	BalanceUpdate(value float64) error
}