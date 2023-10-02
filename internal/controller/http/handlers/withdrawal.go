package handlers

import (
	"fmt"
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
	fmt.Println("Inside")
	var requestData models.Transaction

	if err := c.BindJSON(&requestData); err != nil {
		fmt.Println("bad req")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Println("request: ", requestData)

	if err := w.withdrawalService.Withdrawal(c, requestData); err != nil {
		fmt.Println("service error: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Withdrawal successful"})

}

// runtime error: invalid memory address or nil pointer dereference
//   /usr/local/go/src/runtime/panic.go:261 (0x44df37)
//           panicmem: panic(memoryError)
//   /usr/local/go/src/runtime/signal_unix.go:861 (0x44df05)
//           sigpanic: panicmem()
//   /go/pkg/mod/go.uber.org/zap@v1.26.0/logger.go:330 (0x6ab99e)
//           (*Logger).check: if lvl < zapcore.DPanicLevel && !log.core.Enabled(lvl) {
//   /go/pkg/mod/go.uber.org/zap@v1.26.0/logger.go:237 (0x6ab737)
//           (*Logger).Debug: if ce := log.check(DebugLevel, msg); ce != nil {
//   /go/bin/app/internal/service/OperationService.go:50 (0x9d4924)
//           (*OperationService).Withdrawal: Op.logger.Debug("publicKey was created: ")
//   /go/bin/app/internal/controller/http/handlers/withdrawal.go:46 (0x9d6a59)
//           withdrawalHandler.Withdrawal: if err := w.withdrawalService.Withdrawal(c, requestData); err != nil {
//   /go/pkg/mod/github.com/gin-gonic/gin@v1.9.1/context.go:174 (0x7f3b9d)
//           (*Context).Next: c.handlers[c.index](c)
//   /go/pkg/mod/github.com/gin-gonic/gin@v1.9.1/logger.go:240 (0x7f3b6c)
//           LoggerWithConfig.func1: c.Next()
//   /go/pkg/mod/github.com/gin-gonic/gin@v1.9.1/context.go:174 (0x7ef8aa)
//           (*Context).Next: c.handlers[c.index](c)
//   /go/bin/app/internal/controller/http/middleware/log.go:22 (0x7fdfeb)
//           ApplyMiddlewares.logMiddleware.func1: c.Next()
//   /go/pkg/mod/github.com/gin-gonic/gin@v1.9.1/context.go:174 (0x7f49f9)
//           (*Context).Next: c.handlers[c.index](c)
//   /go/pkg/mod/github.com/gin-gonic/gin@v1.9.1/recovery.go:102 (0x7f49e7)
//           CustomRecoveryWithWriter.func1: c.Next()
//   /go/pkg/mod/github.com/gin-gonic/gin@v1.9.1/context.go:174 (0x7f2c5a)
//           (*Context).Next: c.handlers[c.index](c)
//   /go/pkg/mod/github.com/gin-gonic/gin@v1.9.1/gin.go:620 (0x7f28ed)
//           (*Engine).handleHTTPRequest: c.Next()
//   /go/pkg/mod/github.com/gin-gonic/gin@v1.9.1/gin.go:576 (0x7f259c)
//           (*Engine).ServeHTTP: engine.handleHTTPRequest(c)
//   /usr/local/go/src/net/http/server.go:2938 (0x68a30d)
//           serverHandler.ServeHTTP: handler.ServeHTTP(rw, req)
//   /usr/local/go/src/net/http/server.go:2009 (0x686bb3)
//           (*conn).serve: serverHandler{c.server}.ServeHTTP(w, w.req)
//   /usr/local/go/src/runtime/asm_amd64.s:1650 (0x46a5a0)
//           goexit: BYTE    $0x90   // NOP
