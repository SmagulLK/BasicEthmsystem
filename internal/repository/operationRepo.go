package repository

import (
	"context"
	"fmt"
	"math/big"

	"go.uber.org/zap"

	"TestProjectEthereum/models"
	postgres "TestProjectEthereum/pkg/database/postgresql"
)

type OperationRepository struct {
	Common
	db     *postgres.Postgres
	logger *zap.Logger
}

func NewOperationRepository(db *postgres.Postgres, logger *zap.Logger) *OperationRepository {
	return &OperationRepository{db: db, logger: logger}
}
func (Op *OperationRepository) GetUserByAddress(ctx context.Context, address string) (*models.User, error) {
	statements, arguments, err := Op.db.Builder.Select("users").Where("addres", address).Suffix("RETURN account_id, public_key, private_key, balance, addres").ToSql()
	if err != nil {
		Op.logger.Error(err.Error())
	}
	var user models.User
	Op.db.Pool.QueryRow(ctx, statements, arguments).Scan(user.UserID, user.Balance, user.PrivateKey, user.Balance, user.Address)

	return &user, nil
}
func (Op *OperationRepository) BalanceUpdate(ctx context.Context, balance big.Int) error {
	statement, arguments, err := Op.db.Builder.Update("users").Set("balance", balance).ToSql()
	if err != nil {
		Op.logger.Error(err.Error())
	}
	Op.db.Pool.Exec(ctx, statement, arguments)
	return nil
}

func (Op *OperationRepository) Withdrawal(ctx context.Context, tr *models.Transaction) error {

	sql, args, err := Op.db.Builder.Insert("transactions").Columns("amount", "private_key", "adress_to", "hex").
		Values(tr.ValueBigInt, tr.PrivateKey, tr.AddressTo, tr.Hex).ToSql()
	if err != nil {
		Op.logger.Error(err.Error())
	}

	Op.logger.Info(sql)
	fmt.Println(args)

	result, err := Op.db.Pool.Exec(ctx, sql, args)
	if err != nil {
		Op.logger.Error(err.Error())
	}
	rowsAffected := result.RowsAffected()
	if rowsAffected != 1 {
		Op.logger.Error("expected 1 row to be affected")
		return fmt.Errorf("expected 1 row to be affected, but %d rows were affected", rowsAffected)
	}
	return nil
}
