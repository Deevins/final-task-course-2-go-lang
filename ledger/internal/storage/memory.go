package storage

import (
	"errors"
	"sync"

	"github.com/Deevins/final-task-course-2-go-lang/ledger/internal/model"
)

var ErrNotFound = errors.New("not found")

type InMemoryLedgerStorage struct {
	mu           sync.RWMutex
	transactions map[string]model.Transaction
	budgets      map[string]model.Budget
	reports      map[string]model.Report
}

func NewInMemoryLedgerStorage() *InMemoryLedgerStorage {
	return &InMemoryLedgerStorage{
		transactions: make(map[string]model.Transaction),
		budgets:      make(map[string]model.Budget),
		reports:      make(map[string]model.Report),
	}
}

func (s *InMemoryLedgerStorage) CreateTransaction(tx model.Transaction) model.Transaction {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.transactions[tx.ID] = tx
	return tx
}

func (s *InMemoryLedgerStorage) GetTransaction(id string) (model.Transaction, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	tx, ok := s.transactions[id]
	if !ok {
		return model.Transaction{}, ErrNotFound
	}
	return tx, nil
}

func (s *InMemoryLedgerStorage) UpdateTransaction(tx model.Transaction) (model.Transaction, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, ok := s.transactions[tx.ID]; !ok {
		return model.Transaction{}, ErrNotFound
	}
	s.transactions[tx.ID] = tx
	return tx, nil
}

func (s *InMemoryLedgerStorage) DeleteTransaction(id string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, ok := s.transactions[id]; !ok {
		return ErrNotFound
	}
	delete(s.transactions, id)
	return nil
}

func (s *InMemoryLedgerStorage) ListTransactions() []model.Transaction {
	s.mu.RLock()
	defer s.mu.RUnlock()
	items := make([]model.Transaction, 0, len(s.transactions))
	for _, tx := range s.transactions {
		items = append(items, tx)
	}
	return items
}

func (s *InMemoryLedgerStorage) CreateBudget(budget model.Budget) model.Budget {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.budgets[budget.ID] = budget
	return budget
}

func (s *InMemoryLedgerStorage) GetBudget(id string) (model.Budget, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	budget, ok := s.budgets[id]
	if !ok {
		return model.Budget{}, ErrNotFound
	}
	return budget, nil
}

func (s *InMemoryLedgerStorage) UpdateBudget(budget model.Budget) (model.Budget, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, ok := s.budgets[budget.ID]; !ok {
		return model.Budget{}, ErrNotFound
	}
	s.budgets[budget.ID] = budget
	return budget, nil
}

func (s *InMemoryLedgerStorage) DeleteBudget(id string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, ok := s.budgets[id]; !ok {
		return ErrNotFound
	}
	delete(s.budgets, id)
	return nil
}

func (s *InMemoryLedgerStorage) ListBudgets() []model.Budget {
	s.mu.RLock()
	defer s.mu.RUnlock()
	items := make([]model.Budget, 0, len(s.budgets))
	for _, budget := range s.budgets {
		items = append(items, budget)
	}
	return items
}

func (s *InMemoryLedgerStorage) CreateReport(report model.Report) model.Report {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.reports[report.ID] = report
	return report
}

func (s *InMemoryLedgerStorage) GetReport(id string) (model.Report, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	report, ok := s.reports[id]
	if !ok {
		return model.Report{}, ErrNotFound
	}
	return report, nil
}

func (s *InMemoryLedgerStorage) UpdateReport(report model.Report) (model.Report, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, ok := s.reports[report.ID]; !ok {
		return model.Report{}, ErrNotFound
	}
	s.reports[report.ID] = report
	return report, nil
}

func (s *InMemoryLedgerStorage) DeleteReport(id string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, ok := s.reports[id]; !ok {
		return ErrNotFound
	}
	delete(s.reports, id)
	return nil
}

func (s *InMemoryLedgerStorage) ListReports() []model.Report {
	s.mu.RLock()
	defer s.mu.RUnlock()
	items := make([]model.Report, 0, len(s.reports))
	for _, report := range s.reports {
		items = append(items, report)
	}
	return items
}
