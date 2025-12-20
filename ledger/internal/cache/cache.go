package cache

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"

	"github.com/Deevins/final-task-course-2-go-lang/ledger/internal/model"
)

const budgetListCacheKey = "budgets:all"

var ErrNotFound = errors.New("cache: not found")

//go:generate go run github.com/gojuno/minimock/v3/cmd/minimock -i github.com/Deevins/final-task-course-2-go-lang/ledger/internal/cache.ReportSummaryCache -o ./cache_minimock.go -n ReportSummaryCacheMock
type ReportSummaryCache interface {
	GetSummary(ctx context.Context, key string) (model.ReportSummary, error)
	SetSummary(ctx context.Context, key string, summary model.ReportSummary) error
}

//go:generate go run github.com/gojuno/minimock/v3/cmd/minimock -i github.com/Deevins/final-task-course-2-go-lang/ledger/internal/cache.BudgetListCache -o ./cache_minimock.go -n BudgetListCacheMock
type BudgetListCache interface {
	GetBudgets(ctx context.Context) ([]model.Budget, error)
	SetBudgets(ctx context.Context, budgets []model.Budget) error
	DeleteBudgets(ctx context.Context) error
}

type RedisReportSummaryCache struct {
	client *redis.Client
	ttl    time.Duration
}

func NewReportSummaryCache(client *redis.Client, ttl time.Duration) *RedisReportSummaryCache {
	if client == nil {
		return nil
	}
	return &RedisReportSummaryCache{client: client, ttl: ttl}
}

func (c *RedisReportSummaryCache) GetSummary(ctx context.Context, key string) (model.ReportSummary, error) {
	if key == "" {
		return model.ReportSummary{}, ErrNotFound
	}
	value, err := c.client.Get(ctx, key).Result()
	if err != nil {
		if err == redis.Nil {
			return model.ReportSummary{}, ErrNotFound
		}
		return model.ReportSummary{}, fmt.Errorf("get report summary cache: %w", err)
	}
	var summary model.ReportSummary
	if err := json.Unmarshal([]byte(value), &summary); err != nil {
		return model.ReportSummary{}, fmt.Errorf("decode report summary cache: %w", err)
	}
	return summary, nil
}

func (c *RedisReportSummaryCache) SetSummary(ctx context.Context, key string, summary model.ReportSummary) error {
	if key == "" {
		return nil
	}
	payload, err := json.Marshal(summary)
	if err != nil {
		return fmt.Errorf("encode report summary cache: %w", err)
	}
	if err := c.client.Set(ctx, key, payload, c.ttl).Err(); err != nil {
		return fmt.Errorf("set report summary cache: %w", err)
	}
	return nil
}

type RedisBudgetListCache struct {
	client *redis.Client
	ttl    time.Duration
}

func NewBudgetListCache(client *redis.Client, ttl time.Duration) *RedisBudgetListCache {
	if client == nil {
		return nil
	}
	return &RedisBudgetListCache{client: client, ttl: ttl}
}

func (c *RedisBudgetListCache) GetBudgets(ctx context.Context) ([]model.Budget, error) {
	value, err := c.client.Get(ctx, budgetListCacheKey).Result()
	if err != nil {
		if err == redis.Nil {
			return nil, ErrNotFound
		}
		return nil, fmt.Errorf("get budget list cache: %w", err)
	}
	var budgets []model.Budget
	if err := json.Unmarshal([]byte(value), &budgets); err != nil {
		return nil, fmt.Errorf("decode budget list cache: %w", err)
	}
	return budgets, nil
}

func (c *RedisBudgetListCache) SetBudgets(ctx context.Context, budgets []model.Budget) error {
	payload, err := json.Marshal(budgets)
	if err != nil {
		return fmt.Errorf("encode budget list cache: %w", err)
	}
	if err := c.client.Set(ctx, budgetListCacheKey, payload, c.ttl).Err(); err != nil {
		return fmt.Errorf("set budget list cache: %w", err)
	}
	return nil
}

func (c *RedisBudgetListCache) DeleteBudgets(ctx context.Context) error {
	if err := c.client.Del(ctx, budgetListCacheKey).Err(); err != nil {
		return fmt.Errorf("delete budget list cache: %w", err)
	}
	return nil
}
