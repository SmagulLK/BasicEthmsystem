package service

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"go.uber.org/zap"

	"TestProjectEthereum/internal/repository"
	"TestProjectEthereum/models"
	"TestProjectEthereum/pkg/utils"
)

type GenerationService struct {
	logger *zap.Logger
	repo   *repository.Repository
}

func NewGenService(repo *repository.Repository, logger *zap.Logger) *GenerationService {
	return &GenerationService{logger: logger, repo: repo}
}

func (Gen *GenerationService) Generate() (string, string, string, error) {
	//pvk is private key original
	//pvkStr is string
	pvk, err := crypto.GenerateKey()
	if err != nil {
		Gen.logger.Error(err.Error())
		return "", "", "", err
	}
	pvkStr := hexutil.Encode(crypto.FromECDSA(pvk))
	pubStr := hexutil.Encode(crypto.FromECDSAPub(&pvk.PublicKey))

	address := crypto.PubkeyToAddress(pvk.PublicKey).Hex()
	zero := *big.NewInt(0)
	var user = models.User{
		Balance:    zero,
		UserID:     utils.BeginId,
		PublicKey:  pubStr,
		PrivateKey: pvkStr,
		Address:    address,
	}
	fmt.Println("user: ", user)
	//Gen.logger.Info("user: ",zap.zap(user))
	// ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	// defer cancel()
	err = Gen.repo.InsertData(user)
	if err != nil {
		Gen.logger.Info("ERR INSIDE GEN INSER DATA")
		Gen.logger.Error(err.Error())
		return "", "", "", err
	}
	// return nil
	return address, pvkStr, pubStr, nil
}
