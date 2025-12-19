package service

import (
	"bytes"
	"context"
	"encoding/csv"
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"

	"github.com/Deevins/final-task-course-2-go-lang/ledger/internal/model"
	"github.com/Deevins/final-task-course-2-go-lang/ledger/internal/repository"
	"github.com/Deevins/final-task-course-2-go-lang/ledger/internal/storage"
)

const (
	csvHeaderAccountID   = "account_id"
	csvHeaderAmount      = "amount"
	csvHeaderCurrency    = "currency"
	csvHeaderCategory    = "category"
	csvHeaderDescription = "description"
	csvHeaderOccurredAt  = "occurred_at"
)

type LedgerService interface {
	CreateTransaction(ctx context.Context, tx model.Transaction) (model.Transaction, error)
	GetTransaction(ctx context.Context, id string) (model.Transaction, error)
	UpdateTransaction(ctx context.Context, tx model.Transaction) (model.Transaction, error)
	DeleteTransaction(ctx context.Context, id string) error
	ListTransactions(ctx context.Context, accountID string) []model.Transaction

	CreateBudget(ctx context.Context, budget model.Budget) (model.Budget, error)
	GetBudget(ctx context.Context, accountID, id string) (model.Budget, error)
	UpdateBudget(ctx context.Context, budget model.Budget) (model.Budget, error)
	DeleteBudget(ctx context.Context, accountID, id string) error
	ListBudgets(ctx context.Context, accountID string) []model.Budget

	CreateReport(ctx context.Context, report model.Report) (model.Report, error)
	GetReport(ctx context.Context, accountID, id string) (model.Report, error)
	UpdateReport(ctx context.Context, report model.Report) (model.Report, error)
	DeleteReport(ctx context.Context, accountID, id string) error
	ListReports(ctx context.Context, accountID string) []model.Report

	ImportTransactionsCSV(ctx context.Context, accountID string, csvContent []byte, hasHeader bool) (int, error)
	ExportTransactionsCSV(ctx context.Context, accountID string) ([]byte, error)
}

type DefaultLedgerService struct {
	repo  repository.LedgerRepository
	cache *storage.ReportCache
}

func NewLedgerService(repo repository.LedgerRepository, cache *storage.ReportCache) *DefaultLedgerService {
	return &DefaultLedgerService{repo: repo, cache: cache}
}

func (s *DefaultLedgerService) CreateTransaction(ctx context.Context, tx model.Transaction) (model.Transaction, error) {
	now := time.Now().UTC()
	if tx.ID == "" {
		tx.ID = uuid.NewString()
	}
	if tx.OccurredAt.IsZero() {
		tx.OccurredAt = now
	}
	tx.CreatedAt = now
	tx.UpdatedAt = now
	if err := s.ensureBudgetAvailable(ctx, tx); err != nil {
		return model.Transaction{}, err
	}
	return s.repo.CreateTransaction(ctx, tx)
}

func (s *DefaultLedgerService) GetTransaction(ctx context.Context, id string) (model.Transaction, error) {
	return s.repo.GetTransaction(ctx, id)
}

func (s *DefaultLedgerService) UpdateTransaction(ctx context.Context, tx model.Transaction) (model.Transaction, error) {
	current, err := s.repo.GetTransaction(ctx, tx.ID)
	if err != nil {
		return model.Transaction{}, err
	}
	tx.CreatedAt = current.CreatedAt
	tx.UpdatedAt = time.Now().UTC()
	if tx.OccurredAt.IsZero() {
		tx.OccurredAt = current.OccurredAt
	}
	return s.repo.UpdateTransaction(ctx, tx)
}

func (s *DefaultLedgerService) DeleteTransaction(ctx context.Context, id string) error {
	return s.repo.DeleteTransaction(ctx, id)
}

func (s *DefaultLedgerService) ListTransactions(ctx context.Context, accountID string) []model.Transaction {
	items := s.repo.ListTransactions(ctx)
	if accountID == "" {
		return items
	}
	filtered := make([]model.Transaction, 0, len(items))
	for _, tx := range items {
		if tx.AccountID == accountID {
			filtered = append(filtered, tx)
		}
	}
	return filtered
}

func (s *DefaultLedgerService) CreateBudget(ctx context.Context, budget model.Budget) (model.Budget, error) {
	now := time.Now().UTC()
	if budget.ID == "" {
		budget.ID = uuid.NewString()
	}
	budget.CreatedAt = now
	budget.UpdatedAt = now
	return s.repo.CreateBudget(ctx, budget)
}

func (s *DefaultLedgerService) GetBudget(ctx context.Context, accountID, id string) (model.Budget, error) {
	return s.repo.GetBudget(ctx, accountID, id)
}

func (s *DefaultLedgerService) UpdateBudget(ctx context.Context, budget model.Budget) (model.Budget, error) {
	current, err := s.repo.GetBudget(ctx, budget.AccountID, budget.ID)
	if err != nil {
		return model.Budget{}, err
	}
	budget.CreatedAt = current.CreatedAt
	budget.UpdatedAt = time.Now().UTC()
	if budget.Month == "" {
		budget.Month = current.Month
	}
	return s.repo.UpdateBudget(ctx, budget)
}

func (s *DefaultLedgerService) DeleteBudget(ctx context.Context, accountID, id string) error {
	return s.repo.DeleteBudget(ctx, accountID, id)
}

func (s *DefaultLedgerService) ListBudgets(ctx context.Context, accountID string) []model.Budget {
	return s.repo.ListBudgets(ctx, accountID)
}

func (s *DefaultLedgerService) CreateReport(ctx context.Context, report model.Report) (model.Report, error) {
	now := time.Now().UTC()
	if report.ID == "" {
		report.ID = uuid.NewString()
	}
	if report.GeneratedAt.IsZero() {
		report.GeneratedAt = now
	}
	start, end, err := parsePeriod(report.Period)
	if err != nil {
		return model.Report{}, err
	}
	transactions := s.ListTransactions(ctx, report.AccountID)
	budgets := s.ListBudgets(ctx, report.AccountID)
	reportTotals := buildReportSummary(transactions, budgets, start, end, report.Currency)
	report.TotalIncome = reportTotals.TotalIncome
	report.TotalExpense = reportTotals.TotalExpense
	report.Currency = reportTotals.Currency
	report.Categories = reportTotals.Categories
	created, err := s.repo.CreateReport(ctx, report)
	if err != nil {
		return model.Report{}, err
	}
	s.cacheReport(ctx, created)
	return created, nil
}

func (s *DefaultLedgerService) GetReport(ctx context.Context, accountID, id string) (model.Report, error) {
	if s.cache != nil {
		cached, err := s.cache.GetReport(ctx, id)
		if err == nil {
			if cached.AccountID != "" && cached.AccountID != accountID {
				return model.Report{}, storage.ErrNotFound
			}
			return cached, nil
		}
		if err != storage.ErrNotFound {
			return model.Report{}, err
		}
	}
	report, err := s.repo.GetReport(ctx, accountID, id)
	if err != nil {
		return model.Report{}, err
	}
	s.cacheReport(ctx, report)
	return report, nil
}

func (s *DefaultLedgerService) UpdateReport(ctx context.Context, report model.Report) (model.Report, error) {
	current, err := s.repo.GetReport(ctx, report.AccountID, report.ID)
	if err != nil {
		return model.Report{}, err
	}
	if report.GeneratedAt.IsZero() {
		report.GeneratedAt = time.Now().UTC()
	}
	if report.Period == "" {
		report.Period = current.Period
	}
	if report.Currency == "" {
		report.Currency = current.Currency
	}
	start, end, err := parsePeriod(report.Period)
	if err != nil {
		return model.Report{}, err
	}
	transactions := s.ListTransactions(ctx, report.AccountID)
	budgets := s.ListBudgets(ctx, report.AccountID)
	reportTotals := buildReportSummary(transactions, budgets, start, end, report.Currency)
	report.TotalIncome = reportTotals.TotalIncome
	report.TotalExpense = reportTotals.TotalExpense
	report.Currency = reportTotals.Currency
	report.Categories = reportTotals.Categories
	updated, err := s.repo.UpdateReport(ctx, report)
	if err != nil {
		return model.Report{}, err
	}
	s.cacheReport(ctx, updated)
	return updated, nil
}

func (s *DefaultLedgerService) DeleteReport(ctx context.Context, accountID, id string) error {
	if err := s.repo.DeleteReport(ctx, accountID, id); err != nil {
		return err
	}
	s.invalidateReportCache(ctx, id)
	return nil
}

func (s *DefaultLedgerService) ListReports(ctx context.Context, accountID string) []model.Report {
	return s.repo.ListReports(ctx, accountID)
}

func (s *DefaultLedgerService) ImportTransactionsCSV(ctx context.Context, accountID string, csvContent []byte, hasHeader bool) (int, error) {
	reader := csv.NewReader(bytes.NewReader(csvContent))
	records, err := reader.ReadAll()
	if err != nil {
		return 0, fmt.Errorf("read csv: %w", err)
	}

	start := 0
	if hasHeader && len(records) > 0 {
		start = 1
	}

	count := 0
	for i := start; i < len(records); i++ {
		row := records[i]
		if len(row) != 6 {
			return count, fmt.Errorf("row %d: expected 6 columns", i+1)
		}

		record, err := parseCSVRecord(row)
		if err != nil {
			return count, fmt.Errorf("row %d: %w", i+1, err)
		}

		_, err = s.CreateTransaction(ctx, model.Transaction{
			AccountID:   accountID,
			Amount:      record.Amount,
			Currency:    record.Currency,
			Category:    record.Category,
			Description: record.Description,
			OccurredAt:  record.OccurredAt,
		})
		if err != nil {
			return count, err
		}
		count++
	}
	return count, nil
}

func (s *DefaultLedgerService) ExportTransactionsCSV(ctx context.Context, accountID string) ([]byte, error) {
	transactions := s.ListTransactions(ctx, accountID)
	buf := &bytes.Buffer{}
	writer := csv.NewWriter(buf)

	header := []string{csvHeaderAccountID, csvHeaderAmount, csvHeaderCurrency, csvHeaderCategory, csvHeaderDescription, csvHeaderOccurredAt}
	if err := writer.Write(header); err != nil {
		return nil, fmt.Errorf("write header: %w", err)
	}

	for _, tx := range transactions {
		record := csvRecordFromTransaction(tx)
		if err := writer.Write(record); err != nil {
			return nil, fmt.Errorf("write record: %w", err)
		}
	}

	writer.Flush()
	if err := writer.Error(); err != nil {
		return nil, fmt.Errorf("flush csv: %w", err)
	}

	return buf.Bytes(), nil
}

func IsNotFound(err error) bool {
	return err == storage.ErrNotFound
}

func (s *DefaultLedgerService) ensureBudgetAvailable(ctx context.Context, tx model.Transaction) error {
	if tx.Amount >= 0 {
		return nil
	}
	budgets := s.repo.ListBudgets(ctx, tx.AccountID)
	if len(budgets) == 0 {
		return nil
	}
	expense := -tx.Amount
	transactions := s.ListTransactions(ctx, tx.AccountID)
	for _, budget := range budgets {
		if budget.Currency != tx.Currency {
			continue
		}
		if budget.Name != tx.Category {
			continue
		}
		budgetStart, budgetEnd := monthRange(budget.Month)
		if !withinPeriod(tx.OccurredAt, budgetStart, budgetEnd) {
			continue
		}
		total := 0.0
		for _, existing := range transactions {
			if existing.Amount >= 0 {
				continue
			}
			if existing.Currency != tx.Currency || existing.Category != tx.Category {
				continue
			}
			if !withinPeriod(existing.OccurredAt, budgetStart, budgetEnd) {
				continue
			}
			total += -existing.Amount
		}
		if total+expense > budget.Amount {
			return fmt.Errorf("%w: %s budget exceeded", ErrBudgetExceeded, budget.Name)
		}
	}
	return nil
}

func (s *DefaultLedgerService) cacheReport(ctx context.Context, report model.Report) {
	if s.cache == nil {
		return
	}
	_ = s.cache.SetReport(ctx, report)
}

func (s *DefaultLedgerService) invalidateReportCache(ctx context.Context, id string) {
	if s.cache == nil {
		return
	}
	_ = s.cache.DeleteReport(ctx, id)
}

type reportSummary struct {
	TotalIncome  float64
	TotalExpense float64
	Currency     string
	Categories   []model.ReportCategory
}

func buildReportSummary(transactions []model.Transaction, budgets []model.Budget, start, end time.Time, currency string) reportSummary {
	categoryTotals := map[string]float64{}
	totalIncome := 0.0
	totalExpense := 0.0
	resolvedCurrency := currency

	for _, tx := range transactions {
		if tx.OccurredAt.Before(start) || tx.OccurredAt.After(end) {
			continue
		}
		if currency != "" && tx.Currency != currency {
			continue
		}
		if resolvedCurrency == "" {
			resolvedCurrency = tx.Currency
		}
		if tx.Amount >= 0 {
			totalIncome += tx.Amount
			continue
		}
		expense := -tx.Amount
		totalExpense += expense
		categoryTotals[tx.Category] += expense
	}

	categories := make([]string, 0, len(categoryTotals))
	for category := range categoryTotals {
		categories = append(categories, category)
	}
	sort.Strings(categories)

	results := make([]model.ReportCategory, 0, len(categories))
	for _, category := range categories {
		total := categoryTotals[category]
		budgetAmount := findBudgetAmount(budgets, category, resolvedCurrency, start, end)
		var percent *float64
		if budgetAmount > 0 {
			value := (total / budgetAmount) * 100
			percent = &value
		}
		results = append(results, model.ReportCategory{
			Category:           category,
			TotalExpense:       total,
			BudgetAmount:       budgetAmount,
			BudgetUsagePercent: percent,
		})
	}

	return reportSummary{
		TotalIncome:  totalIncome,
		TotalExpense: totalExpense,
		Currency:     resolvedCurrency,
		Categories:   results,
	}
}

func findBudgetAmount(budgets []model.Budget, category, currency string, start, end time.Time) float64 {
	// Business logic: budget amounts are prorated by overlapping calendar days
	// between the report period and each monthly budget (inclusive boundaries).
	startDate := dateOnly(start)
	endDate := dateOnly(end)
	if endDate.Before(startDate) {
		return 0
	}
	total := 0.0
	for _, budget := range budgets {
		if budget.Name != category {
			continue
		}
		if currency != "" && budget.Currency != currency {
			continue
		}
		budgetStart, budgetEnd := monthDateRange(budget.Month)
		if endDate.Before(budgetStart) || budgetEnd.Before(startDate) {
			continue
		}
		// Business logic: compute inclusive overlap days and prorate by days in month.
		overlapStart := maxDate(startDate, budgetStart)
		overlapEnd := minDate(endDate, budgetEnd)
		overlapDays := int(overlapEnd.Sub(overlapStart).Hours()/24) + 1
		monthDays := budgetEnd.Day()
		if monthDays <= 0 || overlapDays <= 0 {
			continue
		}
		total += budget.Amount * (float64(overlapDays) / float64(monthDays))
	}
	return total
}

func periodsOverlap(startA, endA, startB, endB time.Time) bool {
	if startA.IsZero() || endA.IsZero() {
		return true
	}
	if endB.Before(startA) || endA.Before(startB) {
		return false
	}
	return true
}

func withinPeriod(target, start, end time.Time) bool {
	if start.IsZero() && end.IsZero() {
		return true
	}
	if !start.IsZero() && target.Before(start) {
		return false
	}
	if !end.IsZero() && target.After(end) {
		return false
	}
	return true
}

func parsePeriod(value string) (time.Time, time.Time, error) {
	value = strings.TrimSpace(value)
	if value == "" {
		return time.Time{}, time.Time{}, fmt.Errorf("%w: period is required", ErrValidation)
	}
	if !strings.Contains(value, "/") {
		start, err := time.Parse("2006-01", value)
		if err != nil {
			return time.Time{}, time.Time{}, fmt.Errorf("%w: period must be in start/end or YYYY-MM format", ErrValidation)
		}
		end := start.AddDate(0, 1, 0).Add(-time.Nanosecond)
		return start, end, nil
	}
	parts := strings.Split(value, "/")
	if len(parts) != 2 {
		return time.Time{}, time.Time{}, fmt.Errorf("%w: period must be in start/end format", ErrValidation)
	}
	start, err := parseTimestamp(parts[0])
	if err != nil {
		return time.Time{}, time.Time{}, fmt.Errorf("%w: invalid period start", ErrValidation)
	}
	end, err := parseTimestamp(parts[1])
	if err != nil {
		return time.Time{}, time.Time{}, fmt.Errorf("%w: invalid period end", ErrValidation)
	}
	if end.Before(start) {
		return time.Time{}, time.Time{}, fmt.Errorf("%w: period end before start", ErrValidation)
	}
	return start, end, nil
}

func monthRange(month time.Time) (time.Time, time.Time) {
	month = month.UTC()
	start := time.Date(month.Year(), month.Month(), 1, 0, 0, 0, 0, time.UTC)
	end := start.AddDate(0, 1, 0).Add(-time.Nanosecond)
	return start, end
}

func monthDateRange(month time.Time) (time.Time, time.Time) {
	month = month.UTC()
	start := time.Date(month.Year(), month.Month(), 1, 0, 0, 0, 0, time.UTC)
	end := time.Date(month.Year(), month.Month()+1, 0, 0, 0, 0, 0, time.UTC)
	return start, end
}

func dateOnly(value time.Time) time.Time {
	value = value.UTC()
	return time.Date(value.Year(), value.Month(), value.Day(), 0, 0, 0, 0, time.UTC)
}

func maxDate(a, b time.Time) time.Time {
	if a.After(b) {
		return a
	}
	return b
}

func minDate(a, b time.Time) time.Time {
	if a.Before(b) {
		return a
	}
	return b
}

func parseTimestamp(value string) (time.Time, error) {
	value = strings.TrimSpace(value)
	if value == "" {
		return time.Time{}, fmt.Errorf("empty timestamp")
	}
	parsed, err := time.Parse(time.RFC3339, value)
	if err == nil {
		return parsed, nil
	}
	return time.Parse("2006-01-02", value)
}

func parseCSVRecord(row []string) (model.TransactionCSVRow, error) {
	amount, err := strconv.ParseFloat(row[1], 64)
	if err != nil {
		return model.TransactionCSVRow{}, fmt.Errorf("parse amount: %w", err)
	}

	occurredAt, err := time.Parse(time.RFC3339, row[5])
	if err != nil {
		return model.TransactionCSVRow{}, fmt.Errorf("parse occurred_at: %w", err)
	}

	return model.TransactionCSVRow{
		Amount:      amount,
		Currency:    row[2],
		Category:    row[3],
		Description: row[4],
		OccurredAt:  occurredAt,
	}, nil
}

func csvRecordFromTransaction(tx model.Transaction) []string {
	record := model.TransactionCSVRow{
		AccountID:   tx.AccountID,
		Amount:      tx.Amount,
		Currency:    tx.Currency,
		Category:    tx.Category,
		Description: tx.Description,
		OccurredAt:  tx.OccurredAt,
	}
	return []string{
		record.AccountID,
		strconv.FormatFloat(record.Amount, 'f', -1, 64),
		record.Currency,
		record.Category,
		record.Description,
		record.OccurredAt.Format(time.RFC3339),
	}
}
