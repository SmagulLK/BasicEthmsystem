package models

import (
	"github.com/jackc/pgx/v5/pgtype"
	"math/big"
)

type User struct {
	Balance    pgtype.Numeric `json:"balance" binding:"required"`
	UserID     int            `json:"user_id,omitempty" `
	PublicKey  string         `json:"public_key" binding:"required"`
	PrivateKey string         `json:"private_key" binding:"required"`
	Address    string         `json:"address" binding:"required"`
}
type HotWallet struct {
	PublicKey  string  `json:"public_key" binding:"required"`
	PrivateKey string  `json:"private_key" binding:"required"`
	Address    string  `json:"address" binding:"required"`
	Balance    big.Int `json:"balance" binding:"required"`
}
