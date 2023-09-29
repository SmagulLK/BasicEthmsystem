package handlers

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"TestProjectEthereum/internal/config"
	"TestProjectEthereum/internal/controller/http/middleware"
	"TestProjectEthereum/internal/service"
)

// Deps is a http handler dependencies.
type Deps struct {
	Logger   *zap.Logger
	Services *service.Service
}

// NewRouter returns a new http router.
func NewRouter(deps Deps) *gin.Engine {
	router := gin.New()

	if config.Get().IsProd() {
		gin.SetMode(gin.ReleaseMode)
	}

	middleware.ApplyMiddlewares(router, deps.Logger)
	// Add the Gin logger middleware to log request information
	router.Use(gin.Logger())
	router.GET("/", deps.GenerationWallet)

	return router
}

// api := router.Group("/api")
// {

// 	newPetitionHandler({
// 		router:          api,
// 		petitionService: deps.Services.,
// 	})

// }
