package app

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"go.uber.org/zap"

	"TestProjectEthereum/internal/config"
	"TestProjectEthereum/internal/controller/http/handlers"
	"TestProjectEthereum/internal/repository"
	"TestProjectEthereum/internal/service"
	postgres "TestProjectEthereum/pkg/database/postgresql"
	"TestProjectEthereum/pkg/httpserver"
)

// App is a application interface.
type App struct {
	logger     *zap.Logger
	db         *postgres.Postgres
	httpServer httpserver.Server
}

// New creates a new App.
func New(ctx context.Context, logger *zap.Logger) (App, error) {
	cfg := config.Get()

	postgresInstance, err := postgres.New(
		ctx,
		postgres.NewConnectionConfig(
			cfg.Postgres.Host,
			cfg.Postgres.Port,
			cfg.Postgres.DBName,
			cfg.Postgres.User,
			cfg.Postgres.Password,
			cfg.Postgres.SSLMode,
		),
	)
	if err != nil {
		return App{}, fmt.Errorf("can not connect to postgres: %w", err)
	}

	// passwordHasher := hash.NewArgon2Hasher()

	// tokenMaker, err := token.NewJWTMaker(cfg.Token.SecretKey)
	// if err != nil {
	// 	return App{}, fmt.Errorf("failed to create token maker: %w", err)
	// }

	repositories := repository.NewRepository(postgresInstance, logger)
	services, err := service.NewService(repositories, logger, cfg.Ethereum.TestURL)
	if err != nil {
		return App{}, fmt.Errorf("can not connect to Ethereum: %w", err)
	}

	router := handlers.NewRouter(handlers.Deps{
		Logger:   logger,
		Services: services,
	})

	return App{
		logger: logger,
		db:     postgresInstance,
		httpServer: httpserver.New(
			router,
			httpserver.WithHost(cfg.HTTP.Host),
			httpserver.WithPort(cfg.HTTP.Port),
			httpserver.WithMaxHeaderBytes(cfg.HTTP.MaxHeaderBytes),
			httpserver.WithReadTimeout(cfg.HTTP.ReadTimeout),
			httpserver.WithWriteTimeout(cfg.HTTP.WriteTimeout),
		),
	}, nil
}

// Run runs the application.
func (a App) Run(ctx context.Context) error {
	eChan := make(chan error)
	interrupt := make(chan os.Signal, 1)

	a.logger.Info("Http server is starting")

	go func() {
		if err := a.httpServer.Start(); err != nil {
			eChan <- fmt.Errorf("failed to listen and serve: %w", err)
		}
	}()

	signal.Notify(interrupt, syscall.SIGINT, syscall.SIGTERM)
	select {
	case err := <-eChan:
		return fmt.Errorf("conduit started failed: %w", err)
	case <-interrupt:
	}

	const httpShutdownTimeout = 5 * time.Second
	if err := a.httpServer.Stop(ctx, httpShutdownTimeout); err != nil {
		return fmt.Errorf("failed to stop http server: %w", err)
	}

	return nil
}
