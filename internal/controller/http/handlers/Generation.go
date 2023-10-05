package handlers

import (
	"github.com/gin-gonic/gin"
)

func (h *Deps) GenerationWallet(ctx *gin.Context) {

	addr, pvkStr, pubstr, err := h.Services.GenerationIn.Generate()
	if err != nil {
		ctx.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(200, gin.H{
		"address":    addr,
		"privateKey": pvkStr,
		"publicKey":  pubstr,
	})
}
