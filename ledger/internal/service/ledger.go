package service

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"strconv"
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
	return s.repo.CreateTransaction(tx), nil
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
	return s.repo.CreateBudget(budget), nil
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
	return s.repo.CreateReport(report), nil
}

func (s *DefaultLedgerService) GetReport(id string) (model.Report, error) {
	return s.repo.GetReport(id)
}

func (s *DefaultLedgerService) UpdateReport(report model.Report) (model.Report, error) {
	_, err := s.repo.GetReport(report.ID)
	if err != nil {
		return model.Report{}, err
	}
	if report.GeneratedAt.IsZero() {
		report.GeneratedAt = time.Now().UTC()
	}
	return s.repo.UpdateReport(report)
}

func (s *DefaultLedgerService) DeleteReport(id string) error {
	return s.repo.DeleteReport(id)
}

func (s *DefaultLedgerService) ListReports() []model.Report {
	return s.repo.ListReports()
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

		amount, err := strconv.ParseFloat(row[1], 64)
		if err != nil {
			return count, fmt.Errorf("row %d: parse amount: %w", i+1, err)
		}

		occurredAt, err := time.Parse(time.RFC3339, row[5])
		if err != nil {
			return count, fmt.Errorf("row %d: parse occurred_at: %w", i+1, err)
		}

		_, err = s.CreateTransaction(model.Transaction{
			AccountID:   row[0],
			Amount:      amount,
			Currency:    row[2],
			Category:    row[3],
			Description: row[4],
			OccurredAt:  occurredAt,
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
		record := []string{
			tx.AccountID,
			strconv.FormatFloat(tx.Amount, 'f', -1, 64),
			tx.Currency,
			tx.Category,
			tx.Description,
			tx.OccurredAt.Format(time.RFC3339),
		}
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
