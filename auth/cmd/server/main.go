package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/Deevins/final-task-course-2-go-lang/auth/internal/app"
	"github.com/Deevins/final-task-course-2-go-lang/auth/internal/config"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("load config: %v", err)
	}

	application, err := app.New(ctx, cfg)
	if err != nil {
		log.Fatalf("init app: %v", err)
	}

	if err := application.Run(ctx); err != nil {
		log.Fatalf("server error: %v", err)
	}
}
