package handlers

import (
	"net/http"

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

func NewWithdrawalHandler(deps withdrawalDeps) {
	handler := withdrawalHandler{
		withdrawalService: deps.withdrawalService,
	}

	usersGroup := deps.router.Group("/withdrawal")
	{
		usersGroup.POST("/send", handler.Withdrawal)
	}

}

func (w withdrawalHandler) Withdrawal(c *gin.Context) {

	var requestData *models.Transaction

	if err := c.BindJSON(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	w.withdrawalService.Withdrawal(c, requestData)
	c.JSON(http.StatusOK, gin.H{"message": "Withdrawal successful"})

}
