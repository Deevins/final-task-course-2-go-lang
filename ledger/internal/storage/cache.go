package storage

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"

	"github.com/Deevins/final-task-course-2-go-lang/ledger/internal/model"
)

const reportCacheKeyPrefix = "ledger:report:"

type ReportCache struct {
	client *redis.Client
	ttl    time.Duration
}

func NewReportCache(client *redis.Client, ttl time.Duration) *ReportCache {
	if client == nil {
		return nil
	}
	return &ReportCache{client: client, ttl: ttl}
}

func (c *ReportCache) GetReport(ctx context.Context, id string) (model.Report, error) {
	if id == "" {
		return model.Report{}, ErrNotFound
	}
	value, err := c.client.Get(ctx, reportCacheKey(id)).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return model.Report{}, ErrNotFound
		}
		return model.Report{}, fmt.Errorf("get report cache: %w", err)
	}
	var report model.Report
	if err := json.Unmarshal([]byte(value), &report); err != nil {
		return model.Report{}, fmt.Errorf("decode report cache: %w", err)
	}
	return report, nil
}

func (c *ReportCache) SetReport(ctx context.Context, report model.Report) error {
	if report.ID == "" {
		return nil
	}
	payload, err := json.Marshal(report)
	if err != nil {
		return fmt.Errorf("encode report cache: %w", err)
	}
	if err := c.client.Set(ctx, reportCacheKey(report.ID), payload, c.ttl).Err(); err != nil {
		return fmt.Errorf("set report cache: %w", err)
	}
	return nil
}

func (c *ReportCache) DeleteReport(ctx context.Context, id string) error {
	if id == "" {
		return nil
	}
	if err := c.client.Del(ctx, reportCacheKey(id)).Err(); err != nil {
		return fmt.Errorf("delete report cache: %w", err)
	}
	return nil
}

func reportCacheKey(id string) string {
	return reportCacheKeyPrefix + id
}
