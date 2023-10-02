package scanner

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

type WalletDepositService struct {
	client           *ethclient.Client
	hotWalletAddress common.Address
	privateKey       *ecdsa.PrivateKey
	lastBlockNumber  uint64
}

func NewWalletDepositService(client *ethclient.Client, hotWalletAddress common.Address, privateKey *ecdsa.PrivateKey) *WalletDepositService {
	return &WalletDepositService{
		client:           client,
		hotWalletAddress: hotWalletAddress,
		privateKey:       privateKey,
		lastBlockNumber:  0,
	}
}

func (s *WalletDepositService) Start() {
	for {
		// Get the latest block number.
		latestBlockNumber, err := s.client.BlockNumber(context.Background())
		if err != nil {
			fmt.Println(err)
			continue
		}

		// If the latest block number is greater than the last block number processed,
		// then process the new blocks.
		if latestBlockNumber > s.lastBlockNumber {
			for i := s.lastBlockNumber + 1; i <= latestBlockNumber; i++ {
				// Get the block by number.
				block, err := s.client.BlockByNumber(context.Background(), big.NewInt(int64(i)))
				if err != nil {
					fmt.Println(err)
					continue
				}

				// Process the block for transactions related to the watched accounts.
				for _, tx := range block.Transactions() {

					// Check if the transaction is related to one of the watched accounts.
					if tx.To() != nil && *tx.To() == s.hotWalletAddress {
						// Transfer all funds from the transaction sender to the hot wallet.

						// Get the sender's address (From) using tx.From() method.
						//TODO think about common.Hash{}, uint(0)
						// fromAddress, err := s.client.TransactionSender(context.Background(), tx, common.Hash{}, uint(0))
						// if err != nil {
						// 	fmt.Println(err)
						// 	continue
						// }

						// err = s.transferFunds(fromAddress, tx.Value())
						// if err != nil {
						// 	fmt.Println(err)
						// 	continue
						// }

						// Update the balance in the database for the account.
						// err = s.updateBalance(fromAddress, tx.Value())
						// if err != nil {
						// 	fmt.Println(err)
						// 	continue
						// }
					}
				}

				// Update the last block number processed.
				s.lastBlockNumber = i
			}
		}

		// Wait for 1 second before processing the next block.
		time.Sleep(time.Second)
	}
}

func (s *WalletDepositService) transferFunds(fromAddress common.Address, amount *big.Int) error {
	// Get the latest block number.
	latestBlockNumber, err := s.client.BlockNumber(context.Background())
	if err != nil {
		return err
	}

	// Create a transaction to transfer the funds.
	tx := types.NewTransaction(
		latestBlockNumber+1, // Nonce: Use the next available nonce
		s.hotWalletAddress,  // To: Hot wallet address
		amount,              // Value: Amount to send
		uint64(21000),       // Gas limit: This is a typical gas limit for a simple transaction
		big.NewInt(1),       // Gas price: You can adjust this based on your requirements
		nil,                 // Data: Transaction data (nil for simple value transfer)
	)

	// Sign the transaction with the private key.
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(big.NewInt(int64(s.lastBlockNumber+1))), s.privateKey)
	if err != nil {
		return err
	}

	// Send the transaction to the Ethereum network.
	err = s.client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		return err
	}

	return nil
}

func (s *WalletDepositService) updateBalance(fromAddress common.Address, amount *big.Int) error {
	// TODO: Implement this function to update the balance in the database for the account.

	return nil
}
