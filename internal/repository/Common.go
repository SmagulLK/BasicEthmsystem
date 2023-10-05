package repository

import (
	"context"

	"go.uber.org/zap"

	"TestProjectEthereum/models"
	postgres "TestProjectEthereum/pkg/database/postgresql"
)

type Common struct {
	db     *postgres.Postgres
	logger *zap.Logger
}

//	func NewCommonRepository(db *postgres.Postgres, logger *zap.Logger) *Common {
//		return &Common{db: db, logger: logger}
//	}
func (Op *Common) InsertData(ctx context.Context, user *models.User) error {
	statement, arguments, err := Op.db.Builder.Insert("users").Columns("private_key", "public_key", "addres").Values(user.PrivateKey, user.PublicKey, user.Address).ToSql()
	if err != nil {
		Op.logger.Error(err.Error())
	}
	Op.db.Pool.Exec(ctx, statement, arguments)
	return nil
}
