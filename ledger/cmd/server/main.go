package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/Deevins/final-task-course-2-go-lang/ledger/internal/app"
	"github.com/Deevins/final-task-course-2-go-lang/ledger/internal/config"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	cfg := config.Load()

	application, err := app.New(ctx, cfg)
	if err != nil {
		log.Fatalf("init app: %v", err)
	}

	if err := application.Run(ctx); err != nil {
		log.Fatalf("server error: %v", err)
	}
}
