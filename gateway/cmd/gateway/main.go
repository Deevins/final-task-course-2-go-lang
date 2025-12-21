package main

import (
	"context"
	"log"
	"os/signal"
	"syscall"

	"github.com/Deevins/final-task-course-2-go-lang/gateway/internal/app"
	"github.com/Deevins/final-task-course-2-go-lang/gateway/internal/config"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	cfg := config.Load()

	application, err := app.New(cfg)
	if err != nil {
		log.Fatalf("init app: %v", err)
	}

	if err := application.Run(ctx); err != nil {
		log.Fatalf("server error: %v", err)
	}
}
