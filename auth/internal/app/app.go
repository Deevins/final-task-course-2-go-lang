package app

import (
	"context"
	"errors"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"

	"github.com/Deevins/final-task-course-2-go-lang/auth/internal/config"
	"github.com/Deevins/final-task-course-2-go-lang/auth/internal/grpcserver"
	httpHandler "github.com/Deevins/final-task-course-2-go-lang/auth/internal/handler/http"
	pb "github.com/Deevins/final-task-course-2-go-lang/auth/internal/pb/auth/v1"
	"github.com/Deevins/final-task-course-2-go-lang/auth/internal/repository"
	"github.com/Deevins/final-task-course-2-go-lang/auth/internal/router"
	"github.com/Deevins/final-task-course-2-go-lang/auth/internal/service"
	"github.com/Deevins/final-task-course-2-go-lang/auth/internal/storage"
)

type App struct {
	httpServer      *http.Server
	grpcServer      *grpc.Server
	grpcListener    net.Listener
	db              *pgxpool.Pool
	shutdownTimeout time.Duration
	httpPort        string
	grpcPort        string
}

func New(ctx context.Context, cfg config.Config) (*App, error) {
	engine := gin.Default()

	db, err := storage.NewPostgresPool(ctx, cfg.PostgresDSN)
	if err != nil {
		return nil, err
	}

	repo := repository.NewPostgresAuthRepository(db)
	authService := service.NewAuthService(repo, cfg.JWT)

	healthHandler := httpHandler.NewHealthHandler()
	router.Register(engine, healthHandler)

	httpServer := &http.Server{
		Addr:    ":" + cfg.HTTPPort,
		Handler: engine,
	}

	lis, err := net.Listen("tcp", ":"+cfg.GRPCPort)
	if err != nil {
		db.Close()
		return nil, err
	}

	grpcSrv := grpc.NewServer()
	pb.RegisterAuthServiceServer(grpcSrv, grpcserver.NewAuthServer(authService))

	return &App{
		httpServer:      httpServer,
		grpcServer:      grpcSrv,
		grpcListener:    lis,
		db:              db,
		shutdownTimeout: 5 * time.Second,
		httpPort:        cfg.HTTPPort,
		grpcPort:        cfg.GRPCPort,
	}, nil
}

func (a *App) Run(ctx context.Context) error {
	g, gctx := errgroup.WithContext(ctx)

	g.Go(func() error {
		log.Printf("HTTP server listening on :%s", a.httpPort)
		if err := a.httpServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			return err
		}
		return nil
	})

	g.Go(func() error {
		log.Printf("gRPC server listening on :%s", a.grpcPort)
		if err := a.grpcServer.Serve(a.grpcListener); err != nil {
			return err
		}
		return nil
	})

	g.Go(func() error {
		<-gctx.Done()
		shutdownCtx, cancel := context.WithTimeout(ctx, a.shutdownTimeout)
		defer cancel()

		if err := a.httpServer.Shutdown(shutdownCtx); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Printf("HTTP shutdown error: %v", err)
		}
		a.grpcServer.GracefulStop()
		if a.db != nil {
			a.db.Close()
		}
		return nil
	})

	if err := g.Wait(); err != nil && !errors.Is(err, context.Canceled) {
		return err
	}
	return nil
}
