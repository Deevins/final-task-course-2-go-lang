package app

import (
	"context"
	"errors"
	"log"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"

	"github.com/Deevins/final-task-course-2-go-lang/gateway/internal/config"
	"github.com/Deevins/final-task-course-2-go-lang/gateway/internal/grpcclient"
	"github.com/Deevins/final-task-course-2-go-lang/gateway/internal/handler"
	"github.com/Deevins/final-task-course-2-go-lang/gateway/internal/middleware"
	authv1 "github.com/Deevins/final-task-course-2-go-lang/gateway/internal/pb/auth/v1"
	ledgerv1 "github.com/Deevins/final-task-course-2-go-lang/gateway/internal/pb/ledger/v1"
	"github.com/Deevins/final-task-course-2-go-lang/gateway/internal/router"
	"github.com/Deevins/final-task-course-2-go-lang/gateway/internal/server/httpserver"
	"github.com/Deevins/final-task-course-2-go-lang/gateway/internal/service"
)

type App struct {
	server     *httpserver.Server
	authConn   *grpc.ClientConn
	ledgerConn *grpc.ClientConn
	address    string
}

func New(cfg config.Config) (*App, error) {
	log.Printf("Connecting to auth gRPC backend at %s", cfg.GRPC.AuthAddress)
	authConn, err := grpcclient.Dial(cfg.GRPC.AuthAddress)
	if err != nil {
		return nil, err
	}
	log.Printf("Connected to auth gRPC backend successfully")

	log.Printf("Connecting to ledger gRPC backend at %s", cfg.GRPC.LedgerAddress)
	ledgerConn, err := grpcclient.Dial(cfg.GRPC.LedgerAddress)
	if err != nil {
		authConn.Close()
		return nil, err
	}
	log.Printf("Connected to ledger gRPC backend successfully")

	authService := service.NewAuthGatewayService(authv1.NewAuthServiceClient(authConn))
	ledgerService := service.NewLedgerGatewayService(ledgerv1.NewLedgerServiceClient(ledgerConn))
	authHandler := handler.NewAuthHandler(authService)
	ledgerHandler := handler.NewLedgerHandler(ledgerService)

	engine := gin.New()
	engine.Use(gin.Logger(), gin.Recovery())
	router.Register(engine, authHandler, ledgerHandler, middleware.JWTAuth(authService))

	server := httpserver.New(cfg.HTTP, engine)

	return &App{
		server:     server,
		authConn:   authConn,
		ledgerConn: ledgerConn,
		address:    cfg.HTTP.Address,
	}, nil
}

func (a *App) Run(ctx context.Context) error {
	defer a.closeConnections()

	errCh := make(chan error, 1)
	go func() {
		log.Printf("HTTP server listening on %s", a.address)
		errCh <- a.server.Start()
	}()

	select {
	case <-ctx.Done():
		if err := a.server.Shutdown(ctx); err != nil {
			log.Printf("graceful shutdown failed: %v", err)
		}
		return nil
	case err := <-errCh:
		if err != nil && !errors.Is(err, context.Canceled) {
			return err
		}
		return nil
	}
}

func (a *App) closeConnections() {
	if a.authConn != nil {
		a.authConn.Close()
	}
	if a.ledgerConn != nil {
		a.ledgerConn.Close()
	}
}
