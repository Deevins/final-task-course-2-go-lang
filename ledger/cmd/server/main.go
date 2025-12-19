package main

import (
	"context"
	"errors"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"

	"github.com/Deevins/final-task-course-2-go-lang/ledger/internal/config"
	"github.com/Deevins/final-task-course-2-go-lang/ledger/internal/grpcserver"
	httpHandler "github.com/Deevins/final-task-course-2-go-lang/ledger/internal/handler/http"
	pb "github.com/Deevins/final-task-course-2-go-lang/ledger/internal/pb/ledger/v1"
	"github.com/Deevins/final-task-course-2-go-lang/ledger/internal/repository"
	"github.com/Deevins/final-task-course-2-go-lang/ledger/internal/service"
	"github.com/Deevins/final-task-course-2-go-lang/ledger/internal/storage"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	cfg := config.Load()

	r := gin.Default()

	db, err := storage.NewPostgresPool(ctx, cfg.PostgresDSN)
	if err != nil {
		log.Fatalf("connect postgres: %v", err)
	}
	defer func() {
		db.Close()
	}()

	redisClient, err := storage.NewRedisClient(ctx, cfg.RedisAddr, cfg.RedisPass, cfg.RedisDB)
	if err != nil {
		log.Fatalf("connect redis: %v", err)
	}
	defer func() {
		if err := redisClient.Close(); err != nil {
			log.Printf("close redis: %v", err)
		}
	}()

	repo := repository.NewPostgresLedgerRepository(db)
	reportCache := storage.NewReportCache(redisClient, 5*time.Minute)
	ledgerService := service.NewLedgerService(repo, reportCache)
	validatedService := service.NewValidationService(ledgerService)

	healthHandler := httpHandler.NewHealthHandler()
	api := r.Group("/api/v1")
	healthHandler.RegisterRoutes(api)

	httpPort := cfg.HTTPPort
	grpcPort := cfg.GRPCPort

	httpServer := &http.Server{
		Addr:    ":" + httpPort,
		Handler: r,
	}

	lis, err := net.Listen("tcp", ":"+grpcPort)
	if err != nil {
		log.Fatalf("listen gRPC: %v", err)
	}

	grpcSrv := grpc.NewServer()
	pb.RegisterLedgerServiceServer(grpcSrv, grpcserver.NewLedgerServer(validatedService))

	g, gctx := errgroup.WithContext(ctx)

	g.Go(func() error {
		log.Printf("HTTP server listening on :%s", httpPort)
		if err := httpServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			return err
		}
		return nil
	})

	g.Go(func() error {
		log.Printf("gRPC server listening on :%s", grpcPort)
		if err := grpcSrv.Serve(lis); err != nil {
			return err
		}
		return nil
	})

	g.Go(func() error {
		<-gctx.Done()
		shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := httpServer.Shutdown(shutdownCtx); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Printf("HTTP shutdown error: %v", err)
		}
		grpcSrv.GracefulStop()
		return nil
	})

	if err := g.Wait(); err != nil && !errors.Is(err, context.Canceled) {
		log.Fatalf("server error: %v", err)
	}
}
