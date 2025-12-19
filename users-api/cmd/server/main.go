package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"golang.org/x/sync/errgroup"

	_ "github.com/Deevins/final-task-course-2-go-lang/users-api/internal/docs"
	"github.com/Deevins/final-task-course-2-go-lang/users-api/internal/handler"
	"github.com/Deevins/final-task-course-2-go-lang/users-api/internal/repository"
	"github.com/Deevins/final-task-course-2-go-lang/users-api/internal/service"
	"github.com/Deevins/final-task-course-2-go-lang/users-api/internal/ui"
)

// @title           Users API
// @version         1.0
// @description     Пример API пользователей на Gin
// @BasePath        /api/v1
// @schemes         http
func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	r := gin.Default()

	userRepo := repository.NewInMemoryUserRepository(nil)
	userSvc := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userSvc)

	api := r.Group("/api/v1")
	userHandler.Register(api)

	// Swagger UI (gin-swagger): /swagger/index.html и /swagger/doc.json
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// ReDoc (через embed)
	r.GET("/redoc", func(c *gin.Context) {
		b, err := ui.ReDocHTML()
		if err != nil {
			c.String(http.StatusInternalServerError, "redoc template error: %v", err)
			return
		}
		c.Data(http.StatusOK, "text/html; charset=utf-8", b)
	})

	httpServer := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	g, gctx := errgroup.WithContext(ctx)

	g.Go(func() error {
		log.Printf("HTTP server listening on %s", httpServer.Addr)
		if err := httpServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
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
		return nil
	})

	if err := g.Wait(); err != nil && !errors.Is(err, context.Canceled) {
		log.Fatalf("server error: %v", err)
	}
}
