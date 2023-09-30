package service

import (
	"context"
	"crypto/ecdsa"
	"math/big"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"go.uber.org/zap"

	"TestProjectEthereum/internal/repository"
	"TestProjectEthereum/models"
)

type OperationService struct {
	logger         *zap.Logger
	repo           *repository.Repository
	ethereumClient *ethclient.Client
}

func NewOperationService(repo *repository.Repository, logger *zap.Logger, ethereumClient *ethclient.Client) *OperationService {
	return &OperationService{logger: logger, repo: repo, ethereumClient: ethereumClient}
}
func (OperationServ *OperationService) UpdateBalance(ctx context.Context, balance big.Int) error {
	return OperationServ.repo.BalanceUpdate(ctx, balance)
}

func (Op *OperationService) Withdrawal(ctx context.Context, tr *models.Transaction) error {

	//"fad9c8855b740a0b7ed4c221dbad0f33a83a49cad6b3fe8d5817ac83d38b6a19"
	privateKey, err := crypto.HexToECDSA(tr.PrivateKey)
	if err != nil {
		Op.logger.Error("error creating private key:", zap.Error(err))
		return err
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		Op.logger.Error("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
		return err
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := Op.ethereumClient.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		Op.logger.Error(err.Error())
		return err
	}

	// Convert the string value to int64
	valueInt64, err := strconv.ParseInt(tr.Value, 10, 64)
	if err != nil {
		Op.logger.Error("Invalid 'value' field", zap.Error(err))
		return err
	}

	//1000000000000000000
	value := big.NewInt(valueInt64)     // in wei (1 eth)
	gasLimit := uint64(models.GasLimit) // in units
	gasPrice, err := Op.ethereumClient.SuggestGasPrice(context.Background())
	if err != nil {
		Op.logger.Error(err.Error())
		return err
	}
	//0x4592d8f8d7b001e72cb26a73e4fa1806a51ac79d
	toAddress := common.HexToAddress(tr.AddressTo)
	var data []byte
	tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, data)

	chainID, err := Op.ethereumClient.NetworkID(context.Background())
	if err != nil {
		Op.logger.Error(err.Error())
		return err
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		Op.logger.Error(err.Error())
		return err
	}

	err = Op.ethereumClient.SendTransaction(context.Background(), signedTx)
	if err != nil {
		Op.logger.Error(err.Error())
		return err
	}
	Op.logger.Debug("tx sent: ", zap.String("HEX", signedTx.Hash().Hex()))
	tr.Hex = signedTx.Hash().Hex()

	err = Op.repo.Withdrawal(ctx, tr)
	if err != nil {
		Op.logger.Error(err.Error())
		return err
	}

	Op.logger.Info("tx has been inserted to db: ", zap.String("HEX", signedTx.Hash().Hex()))

	return nil
	//return address, pvkStr, pubStr, nil
}
