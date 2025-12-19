package service

import (
	"context"
	"math"
	"testing"
	"time"

	"github.com/Deevins/final-task-course-2-go-lang/ledger/internal/model"
	"github.com/Deevins/final-task-course-2-go-lang/ledger/internal/repository"
	"github.com/Deevins/final-task-course-2-go-lang/ledger/internal/storage"
)

func TestBudgetTransactionsAndReport(t *testing.T) {
	tests := []struct {
		name string
		run  func(t *testing.T)
	}{
		{
			name: "budget enforcement and report summary",
			run: func(t *testing.T) {
				ctx := context.Background()
				store := storage.NewInMemoryLedgerStorage()
				repo := repository.NewInMemoryLedgerRepository(store)
				service := NewLedgerService(repo, nil)

				accountID := "account-123"
				category := "Food"
				currency := "USD"
				start := time.Date(2024, time.May, 1, 0, 0, 0, 0, time.UTC)
				end := start.AddDate(0, 1, 0).Add(-time.Nanosecond)

				_, err := service.CreateBudget(ctx, model.Budget{
					AccountID: accountID,
					Name:      category,
					Amount:    100,
					Currency:  currency,
					StartDate: start,
					EndDate:   end,
				})
				if err != nil {
					t.Fatalf("create budget: %v", err)
				}

				_, err = service.CreateTransaction(ctx, model.Transaction{
					AccountID:  accountID,
					Amount:     -40,
					Currency:   currency,
					Category:   category,
					OccurredAt: start.AddDate(0, 0, 2),
				})
				if err != nil {
					t.Fatalf("create transaction 1: %v", err)
				}

				_, err = service.CreateTransaction(ctx, model.Transaction{
					AccountID:  accountID,
					Amount:     -50,
					Currency:   currency,
					Category:   category,
					OccurredAt: start.AddDate(0, 0, 10),
				})
				if err != nil {
					t.Fatalf("create transaction 2: %v", err)
				}

				transactions := service.ListTransactions(ctx, accountID)
				if len(transactions) != 2 {
					t.Fatalf("expected 2 transactions, got %d", len(transactions))
				}
				for _, tx := range transactions {
					if tx.AccountID != accountID {
						t.Fatalf("expected transaction to be linked to account %q, got %q", accountID, tx.AccountID)
					}
				}

				_, err = service.CreateTransaction(ctx, model.Transaction{
					AccountID:  accountID,
					Amount:     -20,
					Currency:   currency,
					Category:   category,
					OccurredAt: start.AddDate(0, 0, 15),
				})
				if err == nil || !IsBudgetExceeded(err) {
					t.Fatalf("expected budget exceeded error, got %v", err)
				}

				report, err := service.CreateReport(ctx, model.Report{
					AccountID: accountID,
					Name:      "May Food Report",
					Period:    "2024-05",
					Currency:  currency,
				})
				if err != nil {
					t.Fatalf("create report: %v", err)
				}

				assertFloatNear(t, report.TotalExpense, 90)
				if len(report.Categories) != 1 {
					t.Fatalf("expected 1 report category, got %d", len(report.Categories))
				}
				categorySummary := report.Categories[0]
				if categorySummary.Category != category {
					t.Fatalf("expected category %q, got %q", category, categorySummary.Category)
				}
				assertFloatNear(t, categorySummary.TotalExpense, 90)
				assertFloatNear(t, categorySummary.BudgetAmount, 100)
				assertFloatNear(t, categorySummary.BudgetUsagePercent, 90)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, tt.run)
	}
}

func assertFloatNear(t *testing.T, got, want float64) {
	t.Helper()
	if math.Abs(got-want) > 0.0001 {
		t.Fatalf("expected %.2f, got %.2f", want, got)
	}
}
