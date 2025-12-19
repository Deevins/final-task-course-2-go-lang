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

	"github.com/Deevins/final-task-course-2-go-lang/auth/internal/config"
	"github.com/Deevins/final-task-course-2-go-lang/auth/internal/grpcserver"
	httpHandler "github.com/Deevins/final-task-course-2-go-lang/auth/internal/handler/http"
	pb "github.com/Deevins/final-task-course-2-go-lang/auth/internal/pb/auth/v1"
	"github.com/Deevins/final-task-course-2-go-lang/auth/internal/repository"
	"github.com/Deevins/final-task-course-2-go-lang/auth/internal/service"
	"github.com/Deevins/final-task-course-2-go-lang/auth/internal/storage"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	r := gin.Default()

	jwtConfig, err := config.LoadJWTConfig()
	if err != nil {
		log.Fatalf("load JWT config: %v", err)
	}

	store := storage.NewInMemoryAuthStorage()
	repo := repository.NewInMemoryAuthRepository(store)
	authService := service.NewAuthService(repo, jwtConfig)

	healthHandler := httpHandler.NewHealthHandler()
	api := r.Group("/api/v1")
	healthHandler.RegisterRoutes(api)

	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == "" {
		httpPort = "8082"
	}
	grpcPort := os.Getenv("GRPC_PORT")
	if grpcPort == "" {
		grpcPort = "9092"
	}

	httpServer := &http.Server{
		Addr:    ":" + httpPort,
		Handler: r,
	}

	lis, err := net.Listen("tcp", ":"+grpcPort)
	if err != nil {
		log.Fatalf("listen gRPC: %v", err)
	}

	grpcSrv := grpc.NewServer()
	pb.RegisterAuthServiceServer(grpcSrv, grpcserver.NewAuthServer(authService))

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
