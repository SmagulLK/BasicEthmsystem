package models

import "crypto"

type User struct {
	balance    float64
	userID     int
	publicKey  crypto.PublicKey
	privateKey crypto.PrivateKey
}
