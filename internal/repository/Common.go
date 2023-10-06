package repository

import (
	"context"
	"fmt"
	"time"

	"go.uber.org/zap"

	"TestProjectEthereum/models"
	postgres "TestProjectEthereum/pkg/database/postgresql"
)

type Common struct {
	db     *postgres.Postgres
	logger *zap.Logger
}

func NewCommonRepository(db *postgres.Postgres, logger *zap.Logger) *Common {
	return &Common{db: db, logger: logger}
}
func (Op *Common) InsertData(user models.User) error {
	Op.logger.Info("inside InsertData")

	statement, arguments, err := Op.db.Builder.Insert("users").Columns("private_key", "public_key", "addres", "balance").Values(user.PrivateKey, user.PublicKey, user.Address, user.Balance.Int).ToSql()
	if err != nil {
		Op.logger.Error(err.Error())
	}
	Op.logger.Info(statement)
	fmt.Println(arguments)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	_, err = Op.db.Pool.Exec(ctx, statement, arguments...)
	if err != nil {
		Op.logger.Error(err.Error())
	}
	return nil
}
