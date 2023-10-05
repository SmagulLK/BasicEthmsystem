package handlers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"TestProjectEthereum/internal/service"
	"TestProjectEthereum/models"
)

// type WithdrawalService interface {
// 	Withdrawal(ctx context.Context, tr models.Transaction) error
// }

type withdrawalDeps struct {
	router            *gin.RouterGroup
	withdrawalService service.OperationService
}

type withdrawalHandler struct {
	withdrawalService service.OperationService
}

func NewWithdrawalHandler(deps withdrawalDeps) {
	handler := withdrawalHandler{
		withdrawalService: deps.withdrawalService,
	}

	usersGroup := deps.router.Group("/withdraw")
	{
		usersGroup.POST("/send", handler.Withdrawal)
	}

}

func (w withdrawalHandler) Withdrawal(c *gin.Context) {
	fmt.Println("Inside")
	var requestData models.Transaction

	if err := c.BindJSON(&requestData); err != nil {
		fmt.Println("bad req")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Println("request: ", requestData)

	if err := w.withdrawalService.Withdrawal(context.Background(), requestData); err != nil {
		fmt.Println("service error: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Withdrawal successful"})

}
