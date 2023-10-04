package models

type Transaction struct {
	Value      string `json:"value"`
	PrivateKey string `json:"private_key"`
	AddressTo  string `json:"address"`
	Hex        string `json:"hex,omitempty"` // this is no need on frontend
}

const (
	GasLimit = 210000
)
