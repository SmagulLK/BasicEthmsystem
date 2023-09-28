package repository

import (
	"TestProjectEthereum/models"
	"github.com/jmoiron/sqlx"
)

type OperationRepository struct {
	db *sqlx.DB
}

func NewOperationRepository(db *sqlx.DB) *OperationRepository {
	return &OperationRepository{db: db}
}
func (Op *OperationRepository) GetUserbyId(id int) (*models.User, error) {

	return &models.User{}, nil
}
func (Op *OperationRepository) BalanceUpdate(balance int32) error {
	return nil
}

func (Op *OperationRepository) InsertData(user *models.User) error {

	return nil
}
