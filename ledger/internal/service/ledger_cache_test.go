package service

import (
	"context"
	"testing"
	"time"

	"github.com/Deevins/final-task-course-2-go-lang/ledger/internal/cache"
	"github.com/Deevins/final-task-course-2-go-lang/ledger/internal/model"
	"github.com/Deevins/final-task-course-2-go-lang/ledger/internal/repository"
	"github.com/Deevins/final-task-course-2-go-lang/ledger/internal/storage"
)

func TestGetReportSummaryCachesResults(t *testing.T) {
	ctx := context.Background()
	store := storage.NewInMemoryLedgerStorage()
	repo := repository.NewInMemoryLedgerRepository(store)
	summaryCache := cache.NewReportSummaryCacheMock(t)
	summaryCache.GetSummaryFunc = func(ctx context.Context, key string) (model.ReportSummary, error) {
		return model.ReportSummary{}, cache.ErrNotFound
	}
	summaryCache.SetSummaryFunc = func(ctx context.Context, key string, summary model.ReportSummary) error {
		return nil
	}
	service := NewLedgerService(repo, nil, summaryCache, nil)

	accountID := "account-1"
	category := "Food"
	currency := "USD"
	start := time.Date(2024, time.March, 1, 0, 0, 0, 0, time.UTC)
	end := time.Date(2024, time.March, 31, 23, 59, 59, 0, time.UTC)

	_, err := service.CreateBudget(ctx, model.Budget{
		AccountID: accountID,
		Name:      category,
		Amount:    200,
		Currency:  currency,
		Period:    "monthly",
		Month:     start,
	})
	if err != nil {
		t.Fatalf("create budget: %v", err)
	}

	_, err = service.CreateTransaction(ctx, model.Transaction{
		AccountID:  accountID,
		Amount:     -50,
		Currency:   currency,
		Category:   category,
		OccurredAt: start.AddDate(0, 0, 2),
	})
	if err != nil {
		t.Fatalf("create transaction: %v", err)
	}

	summary, err := service.GetReportSummary(ctx, accountID, start, end)
	if err != nil {
		t.Fatalf("get summary: %v", err)
	}
	if summary.TotalExpense != 50 {
		t.Fatalf("expected total expense 50, got %.2f", summary.TotalExpense)
	}
	if summaryCache.SetSummaryCalls() != 1 {
		t.Fatalf("expected summary cache set once, got %d", summaryCache.SetSummaryCalls())
	}
}

func TestListBudgetsUsesCacheAndInvalidatesOnWrite(t *testing.T) {
	ctx := context.Background()
	store := storage.NewInMemoryLedgerStorage()
	repo := repository.NewInMemoryLedgerRepository(store)
	budgetCache := cache.NewBudgetListCacheMock(t)
	budgetCache.GetBudgetsFunc = func(ctx context.Context) ([]model.Budget, error) {
		return nil, cache.ErrNotFound
	}
	budgetCache.SetBudgetsFunc = func(ctx context.Context, budgets []model.Budget) error {
		return nil
	}
	budgetCache.DeleteBudgetsFunc = func(ctx context.Context) error {
		return nil
	}
	service := NewLedgerService(repo, nil, nil, budgetCache)

	accountID := "account-2"
	budget := model.Budget{
		AccountID: accountID,
		Name:      "Rent",
		Amount:    1000,
		Currency:  "USD",
		Period:    "monthly",
		Month:     time.Date(2024, time.April, 1, 0, 0, 0, 0, time.UTC),
	}

	_, err := service.CreateBudget(ctx, budget)
	if err != nil {
		t.Fatalf("create budget: %v", err)
	}
	if budgetCache.DeleteBudgetsCalls() != 1 {
		t.Fatalf("expected budget cache invalidation, got %d", budgetCache.DeleteBudgetsCalls())
	}

	items := service.ListBudgets(ctx, accountID)
	if len(items) != 1 {
		t.Fatalf("expected 1 budget, got %d", len(items))
	}
	if budgetCache.SetBudgetsCalls() != 1 {
		t.Fatalf("expected budget cache set once, got %d", budgetCache.SetBudgetsCalls())
	}

	budgetCache.GetBudgetsFunc = func(ctx context.Context) ([]model.Budget, error) {
		return items, nil
	}
	items = service.ListBudgets(ctx, accountID)
	if len(items) != 1 {
		t.Fatalf("expected 1 cached budget, got %d", len(items))
	}
	if budgetCache.GetBudgetsCalls() == 0 {
		t.Fatalf("expected budget cache get to be called")
	}
}
