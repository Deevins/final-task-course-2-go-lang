package service

import (
	"bytes"
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
	CreateTransaction(tx model.Transaction) (model.Transaction, error)
	GetTransaction(id string) (model.Transaction, error)
	UpdateTransaction(tx model.Transaction) (model.Transaction, error)
	DeleteTransaction(id string) error
	ListTransactions(accountID string) []model.Transaction

	CreateBudget(budget model.Budget) (model.Budget, error)
	GetBudget(id string) (model.Budget, error)
	UpdateBudget(budget model.Budget) (model.Budget, error)
	DeleteBudget(id string) error
	ListBudgets() []model.Budget

	CreateReport(report model.Report) (model.Report, error)
	GetReport(id string) (model.Report, error)
	UpdateReport(report model.Report) (model.Report, error)
	DeleteReport(id string) error
	ListReports() []model.Report

	ImportTransactionsCSV(csvContent []byte, hasHeader bool) (int, error)
	ExportTransactionsCSV(accountID string) ([]byte, error)
	ExportReportSheet(reportID string) (model.ReportSheet, error)
}

type DefaultLedgerService struct {
	repo repository.LedgerRepository
}

func NewLedgerService(repo repository.LedgerRepository) *DefaultLedgerService {
	return &DefaultLedgerService{repo: repo}
}

func (s *DefaultLedgerService) CreateTransaction(tx model.Transaction) (model.Transaction, error) {
	now := time.Now().UTC()
	if tx.ID == "" {
		tx.ID = uuid.NewString()
	}
	if tx.OccurredAt.IsZero() {
		tx.OccurredAt = now
	}
	tx.CreatedAt = now
	tx.UpdatedAt = now
	if err := s.ensureBudgetAvailable(tx); err != nil {
		return model.Transaction{}, err
	}
	return s.repo.CreateTransaction(tx)
}

func (s *DefaultLedgerService) GetTransaction(id string) (model.Transaction, error) {
	return s.repo.GetTransaction(id)
}

func (s *DefaultLedgerService) UpdateTransaction(tx model.Transaction) (model.Transaction, error) {
	current, err := s.repo.GetTransaction(tx.ID)
	if err != nil {
		return model.Transaction{}, err
	}
	tx.CreatedAt = current.CreatedAt
	tx.UpdatedAt = time.Now().UTC()
	if tx.OccurredAt.IsZero() {
		tx.OccurredAt = current.OccurredAt
	}
	return s.repo.UpdateTransaction(tx)
}

func (s *DefaultLedgerService) DeleteTransaction(id string) error {
	return s.repo.DeleteTransaction(id)
}

func (s *DefaultLedgerService) ListTransactions(accountID string) []model.Transaction {
	items := s.repo.ListTransactions()
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

func (s *DefaultLedgerService) CreateBudget(budget model.Budget) (model.Budget, error) {
	now := time.Now().UTC()
	if budget.ID == "" {
		budget.ID = uuid.NewString()
	}
	budget.CreatedAt = now
	budget.UpdatedAt = now
	return s.repo.CreateBudget(budget)
}

func (s *DefaultLedgerService) GetBudget(id string) (model.Budget, error) {
	return s.repo.GetBudget(id)
}

func (s *DefaultLedgerService) UpdateBudget(budget model.Budget) (model.Budget, error) {
	current, err := s.repo.GetBudget(budget.ID)
	if err != nil {
		return model.Budget{}, err
	}
	budget.CreatedAt = current.CreatedAt
	budget.UpdatedAt = time.Now().UTC()
	if budget.StartDate.IsZero() {
		budget.StartDate = current.StartDate
	}
	if budget.EndDate.IsZero() {
		budget.EndDate = current.EndDate
	}
	return s.repo.UpdateBudget(budget)
}

func (s *DefaultLedgerService) DeleteBudget(id string) error {
	return s.repo.DeleteBudget(id)
}

func (s *DefaultLedgerService) ListBudgets() []model.Budget {
	return s.repo.ListBudgets()
}

func (s *DefaultLedgerService) CreateReport(report model.Report) (model.Report, error) {
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
	transactions := s.repo.ListTransactions()
	budgets := s.repo.ListBudgets()
	reportTotals := buildReportSummary(transactions, budgets, start, end, report.Currency)
	report.TotalIncome = reportTotals.TotalIncome
	report.TotalExpense = reportTotals.TotalExpense
	report.Currency = reportTotals.Currency
	report.Categories = reportTotals.Categories
	return s.repo.CreateReport(report)
}

func (s *DefaultLedgerService) GetReport(id string) (model.Report, error) {
	return s.repo.GetReport(id)
}

func (s *DefaultLedgerService) UpdateReport(report model.Report) (model.Report, error) {
	current, err := s.repo.GetReport(report.ID)
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
	transactions := s.repo.ListTransactions()
	budgets := s.repo.ListBudgets()
	reportTotals := buildReportSummary(transactions, budgets, start, end, report.Currency)
	report.TotalIncome = reportTotals.TotalIncome
	report.TotalExpense = reportTotals.TotalExpense
	report.Currency = reportTotals.Currency
	report.Categories = reportTotals.Categories
	return s.repo.UpdateReport(report)
}

func (s *DefaultLedgerService) DeleteReport(id string) error {
	return s.repo.DeleteReport(id)
}

func (s *DefaultLedgerService) ListReports() []model.Report {
	return s.repo.ListReports()
}

func (s *DefaultLedgerService) ExportReportSheet(reportID string) (model.ReportSheet, error) {
	report, err := s.repo.GetReport(reportID)
	if err != nil {
		return model.ReportSheet{}, err
	}

	categories := report.Categories
	summary := model.ReportSheetSummary{
		ReportID:     report.ID,
		Name:         report.Name,
		Period:       report.Period,
		GeneratedAt:  report.GeneratedAt,
		TotalIncome:  report.TotalIncome,
		TotalExpense: report.TotalExpense,
		Currency:     report.Currency,
	}

	if len(categories) == 0 || (summary.TotalIncome == 0 && summary.TotalExpense == 0) {
		start, end, err := parsePeriod(report.Period)
		if err != nil {
			return model.ReportSheet{}, err
		}
		transactions := s.repo.ListTransactions()
		budgets := s.repo.ListBudgets()
		reportTotals := buildReportSummary(transactions, budgets, start, end, report.Currency)
		summary.TotalIncome = reportTotals.TotalIncome
		summary.TotalExpense = reportTotals.TotalExpense
		if summary.Currency == "" {
			summary.Currency = reportTotals.Currency
		}
		categories = reportTotals.Categories
	}

	sheetCategories := make([]model.ReportSheetCategory, 0, len(categories))
	for _, category := range categories {
		sheetCategories = append(sheetCategories, model.ReportSheetCategory{
			Category:           category.Category,
			TotalExpense:       category.TotalExpense,
			BudgetAmount:       category.BudgetAmount,
			BudgetUsagePercent: category.BudgetUsagePercent,
		})
	}

	return model.ReportSheet{
		Summary:    summary,
		Categories: sheetCategories,
	}, nil
}

func (s *DefaultLedgerService) ImportTransactionsCSV(csvContent []byte, hasHeader bool) (int, error) {
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
		if len(row) < 6 {
			return count, fmt.Errorf("row %d: expected 6 columns", i+1)
		}

		record, err := parseCSVRecord(row)
		if err != nil {
			return count, fmt.Errorf("row %d: %w", i+1, err)
		}

		_, err = s.CreateTransaction(model.Transaction{
			AccountID:   record.AccountID,
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

func (s *DefaultLedgerService) ExportTransactionsCSV(accountID string) ([]byte, error) {
	transactions := s.ListTransactions(accountID)
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

func (s *DefaultLedgerService) ensureBudgetAvailable(tx model.Transaction) error {
	if tx.Amount >= 0 {
		return nil
	}
	budgets := s.repo.ListBudgets()
	if len(budgets) == 0 {
		return nil
	}
	expense := -tx.Amount
	transactions := s.repo.ListTransactions()
	for _, budget := range budgets {
		if budget.Currency != tx.Currency {
			continue
		}
		if budget.Name != tx.Category {
			continue
		}
		if !withinPeriod(tx.OccurredAt, budget.StartDate, budget.EndDate) {
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
			if !withinPeriod(existing.OccurredAt, budget.StartDate, budget.EndDate) {
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
		percent := 0.0
		if budgetAmount > 0 {
			percent = (total / budgetAmount) * 100
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
	for _, budget := range budgets {
		if budget.Name != category {
			continue
		}
		if currency != "" && budget.Currency != currency {
			continue
		}
		if !periodsOverlap(budget.StartDate, budget.EndDate, start, end) {
			continue
		}
		return budget.Amount
	}
	return 0
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
		AccountID:   row[0],
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
