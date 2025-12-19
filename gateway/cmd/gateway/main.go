package main

import (
	"context"
	"log"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/Deevins/final-task-course-2-go-lang/gateway/internal/config"
	"github.com/Deevins/final-task-course-2-go-lang/gateway/internal/handler"
	"github.com/Deevins/final-task-course-2-go-lang/gateway/internal/middleware"
	authv1 "github.com/Deevins/final-task-course-2-go-lang/gateway/internal/pb/auth/v1"
	ledgerv1 "github.com/Deevins/final-task-course-2-go-lang/gateway/internal/pb/ledger/v1"
	"github.com/Deevins/final-task-course-2-go-lang/gateway/internal/server/httpserver"
	"github.com/Deevins/final-task-course-2-go-lang/gateway/internal/service"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	cfg := config.Load()

	log.Printf("Connecting to auth gRPC backend at %s", cfg.GRPC.AuthAddress)
	authConn, err := dialGRPC(cfg.GRPC.AuthAddress)
	if err != nil {
		log.Fatalf("failed to connect to auth gRPC backend: %v", err)
	}
	defer authConn.Close()
	log.Printf("Connected to auth gRPC backend successfully")

	log.Printf("Connecting to ledger gRPC backend at %s", cfg.GRPC.LedgerAddress)
	ledgerConn, err := dialGRPC(cfg.GRPC.LedgerAddress)
	if err != nil {
		log.Fatalf("failed to connect to ledger gRPC backend: %v", err)
	}
	defer ledgerConn.Close()
	log.Printf("Connected to ledger gRPC backend successfully")

	authService := service.NewAuthGatewayService(authv1.NewAuthServiceClient(authConn))
	ledgerService := service.NewLedgerGatewayService(ledgerv1.NewLedgerServiceClient(ledgerConn))
	authHandler := handler.NewAuthHandler(authService)
	ledgerHandler := handler.NewLedgerHandler(ledgerService)

	engine := gin.New()
	engine.Use(gin.Logger(), gin.Recovery())
	api := engine.Group("/api")
	authHandler.Register(api)
	ledgerHandler.Register(api, middleware.JWTAuth(authService))

	server := httpserver.New(cfg.HTTP, engine)

	go func() {
		log.Printf("HTTP server listening on %s", cfg.HTTP.Address)
		if err := server.Start(); err != nil {
			log.Fatalf("http server stopped: %v", err)
		}
	}()

	<-ctx.Done()

	shutdownCtx, cancel := context.WithTimeout(context.Background(), cfg.HTTP.ShutdownTimeout)
	defer cancel()

	if err := server.Shutdown(shutdownCtx); err != nil {
		log.Printf("graceful shutdown failed: %v", err)
	}
}

func dialGRPC(address string) (*grpc.ClientConn, error) {
	conn, err := grpc.NewClient(
		address,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, err
	}
	return conn, nil
}
