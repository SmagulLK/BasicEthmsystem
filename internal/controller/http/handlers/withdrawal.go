package handlers

import (
	"math/big"
	"net/http"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/gin-gonic/gin"

	"TestProjectEthereum/internal/service"
	"TestProjectEthereum/models"
)

type withdrawalDeps struct {
	router            *gin.RouterGroup
	withdrawalService service.OperationService
}

type withdrawalHandler struct {
	withdrawalService service.OperationService
}

func newPetitionHandler(deps withdrawalDeps) {
	handler := withdrawalHandler{
		withdrawalService: deps.withdrawalService,
	}

	usersGroup := deps.router.Group("/withdrawal")
	{

		usersGroup.POST("/send", handler.Withdrawal)
	}

}

func (w withdrawalHandler) Withdrawal(c *gin.Context) {

	var requestData models.TransactionFromFrontend

	if err := c.BindJSON(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Convert the string value to *big.Int
	value := new(big.Int)
	value, success := value.SetString(requestData.Value, 10)
	if !success {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid 'value' field"})
		return
	}

	// Convert the string private key to *ecdsa.PrivateKey
	privateKey, err := crypto.HexToECDSA(requestData.PrivateKey)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid 'private_key' field"})
		return
	}

	// Convert the string address to common.Address
	addressTo := common.HexToAddress(requestData.AddressTo)

	w.withdrawalService.Withdrawal(c, models.Transaction{
		Value:      value,
		PrivateKey: privateKey,
		AddressTo:  addressTo,
	})
	c.JSON(http.StatusOK, gin.H{"message": "Withdrawal successful"})

}
