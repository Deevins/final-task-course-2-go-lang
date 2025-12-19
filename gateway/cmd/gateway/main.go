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
	budgetv1 "github.com/Deevins/final-task-course-2-go-lang/gateway/internal/pb/budget/v1"
	"github.com/Deevins/final-task-course-2-go-lang/gateway/internal/server/httpserver"
	"github.com/Deevins/final-task-course-2-go-lang/gateway/internal/service"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	cfg := config.Load()

	log.Printf("Connecting to gRPC backend at %s", cfg.GRPC.Address)
	conn, err := dialGRPC(cfg.GRPC.Address)
	if err != nil {
		log.Fatalf("failed to connect to gRPC backend: %v", err)
	}
	defer conn.Close()
	log.Printf("Connected to gRPC backend successfully")

	budgetService := service.NewBudgetGatewayService(budgetv1.NewBudgetServiceClient(conn))
	budgetHandler := handler.NewBudgetHandler(budgetService)

	engine := gin.New()
	engine.Use(gin.Logger(), gin.Recovery())
	api := engine.Group("/api")
	budgetHandler.Register(api)

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
