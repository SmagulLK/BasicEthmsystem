package repository

import (
	"TestProjectEthereum/models"
	postgres "TestProjectEthereum/pkg/database/postgresql"
	"context"
	"go.uber.org/zap"
)

type Common struct {
	db     *postgres.Postgres
	logger *zap.Logger
}

//	func NewCommonRepository(db *postgres.Postgres, logger *zap.Logger) *Common {
//		return &Common{db: db, logger: logger}
//	}
func (Op *Common) InsertData(ctx context.Context, user *models.User) error {
	statement, arguments, err := Op.db.Builder.Insert("User").Columns("private_key", "public_key", "address").Values(user.PrivateKey, user.PublicKey, user.Address).ToSql()
	if err != nil {
		Op.logger.Error(err.Error())
	}
	Op.db.Pool.Exec(ctx, statement, arguments)
	return nil
}
