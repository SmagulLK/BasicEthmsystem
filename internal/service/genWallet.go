package service

import (
	"context"
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

func (Gen *GenerationService) Generate(ctx context.Context) (string, string, string, error) {
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
	var user = models.User{
		*new(big.Int),
		utils.BeginId,
		pubStr,
		pvkStr,
		address,
	}

	err = Gen.repo.CommonIn.InsertData(ctx, &user)
	if err != nil {
		Gen.logger.Error(err.Error())
		return "", "", "", err
	}
	// return nil
	return address, pvkStr, pubStr, nil
}
