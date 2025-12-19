package service

import (
	"bytes"
	"context"
	"encoding/csv"
	"fmt"
	"strconv"
	"time"

	"github.com/Deevins/final-task-course-2-go-lang/gateway/internal/model"
	ledgerv1 "github.com/Deevins/final-task-course-2-go-lang/gateway/internal/pb/ledger/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type LedgerGatewayService interface {
	ListTransactions(ctx context.Context, accountID string) ([]model.Transaction, error)
	CreateTransaction(ctx context.Context, req model.CreateTransactionRequest) (*model.Transaction, error)
	ListBudgets(ctx context.Context) ([]model.Budget, error)
	CreateBudget(ctx context.Context, req model.CreateBudgetRequest) (*model.Budget, error)
	ListReports(ctx context.Context) ([]model.Report, error)
	CreateReport(ctx context.Context, req model.CreateReportRequest) (*model.Report, error)
	ImportTransactionsCSV(ctx context.Context, csvContent []byte, hasHeader bool) (int32, error)
	ExportTransactionsCSV(ctx context.Context, accountID string) ([]byte, error)
	ImportTransactionsSheet(ctx context.Context, req model.ImportTransactionsSheetRequest) (int32, error)
}

type ledgerGatewayService struct {
	client ledgerv1.LedgerServiceClient
}

func NewLedgerGatewayService(client ledgerv1.LedgerServiceClient) LedgerGatewayService {
	if client == nil {
		panic("ledger gateway service requires gRPC client")
	}
	return &ledgerGatewayService{client: client}
}

func (s *ledgerGatewayService) ListTransactions(ctx context.Context, accountID string) ([]model.Transaction, error) {
	resp, err := s.client.ListTransactions(ctx, &ledgerv1.ListTransactionsRequest{AccountId: accountID})
	if err != nil {
		return nil, err
	}
	return fromProtoTransactions(resp.GetTransactions()), nil
}

func (s *ledgerGatewayService) CreateTransaction(ctx context.Context, req model.CreateTransactionRequest) (*model.Transaction, error) {
	resp, err := s.client.CreateTransaction(ctx, &ledgerv1.CreateTransactionRequest{
		Transaction: &ledgerv1.Transaction{
			AccountId:   req.AccountID,
			Amount:      req.Amount,
			Currency:    req.Currency,
			Category:    req.Category,
			Description: req.Description,
			OccurredAt:  timestamppb.New(req.OccurredAt),
		},
	})
	if err != nil {
		return nil, err
	}
	return fromProtoTransaction(resp.GetTransaction()), nil
}

func (s *ledgerGatewayService) ListBudgets(ctx context.Context) ([]model.Budget, error) {
	resp, err := s.client.ListBudgets(ctx, &ledgerv1.ListBudgetsRequest{})
	if err != nil {
		return nil, err
	}
	return fromProtoBudgets(resp.GetBudgets()), nil
}

func (s *ledgerGatewayService) CreateBudget(ctx context.Context, req model.CreateBudgetRequest) (*model.Budget, error) {
	resp, err := s.client.CreateBudget(ctx, &ledgerv1.CreateBudgetRequest{
		Budget: &ledgerv1.Budget{
			Name:      req.Name,
			Amount:    req.Amount,
			Currency:  req.Currency,
			Period:    req.Period,
			StartDate: timestamppb.New(req.StartDate),
			EndDate:   timestamppb.New(req.EndDate),
		},
	})
	if err != nil {
		return nil, err
	}
	return fromProtoBudget(resp.GetBudget()), nil
}

func (s *ledgerGatewayService) ListReports(ctx context.Context) ([]model.Report, error) {
	resp, err := s.client.ListReports(ctx, &ledgerv1.ListReportsRequest{})
	if err != nil {
		return nil, err
	}
	return fromProtoReports(resp.GetReports()), nil
}

func (s *ledgerGatewayService) CreateReport(ctx context.Context, req model.CreateReportRequest) (*model.Report, error) {
	resp, err := s.client.CreateReport(ctx, &ledgerv1.CreateReportRequest{
		Report: &ledgerv1.Report{
			Name:        req.Name,
			Period:      req.Period,
			GeneratedAt: timestamppb.New(req.GeneratedAt),
			Currency:    req.Currency,
		},
	})
	if err != nil {
		return nil, err
	}
	return fromProtoReport(resp.GetReport()), nil
}

func (s *ledgerGatewayService) ImportTransactionsCSV(ctx context.Context, csvContent []byte, hasHeader bool) (int32, error) {
	resp, err := s.client.ImportTransactionsCsv(ctx, &ledgerv1.ImportTransactionsCsvRequest{
		CsvContent: csvContent,
		HasHeader:  hasHeader,
	})
	if err != nil {
		return 0, err
	}
	return resp.GetImported(), nil
}

func (s *ledgerGatewayService) ExportTransactionsCSV(ctx context.Context, accountID string) ([]byte, error) {
	resp, err := s.client.ExportTransactionsCsv(ctx, &ledgerv1.ExportTransactionsCsvRequest{AccountId: accountID})
	if err != nil {
		return nil, err
	}
	return resp.GetCsvContent(), nil
}

func (s *ledgerGatewayService) ImportTransactionsSheet(ctx context.Context, req model.ImportTransactionsSheetRequest) (int32, error) {
	csvContent, err := buildSheetCSV(req.Rows)
	if err != nil {
		return 0, err
	}
	resp, err := s.client.ImportTransactionsCsv(ctx, &ledgerv1.ImportTransactionsCsvRequest{
		CsvContent: csvContent,
		HasHeader:  true,
	})
	if err != nil {
		return 0, err
	}
	return resp.GetImported(), nil
}

func buildSheetCSV(rows []model.SheetTransactionRow) ([]byte, error) {
	if len(rows) == 0 {
		return nil, fmt.Errorf("rows are required")
	}
	var buf bytes.Buffer
	writer := csv.NewWriter(&buf)
	if err := writer.Write([]string{
		"account_id",
		"amount",
		"currency",
		"category",
		"description",
		"occurred_at",
	}); err != nil {
		return nil, fmt.Errorf("write header: %w", err)
	}
	for _, row := range rows {
		record := []string{
			row.AccountID,
			strconv.FormatFloat(row.Amount, 'f', -1, 64),
			row.Currency,
			row.Category,
			row.Description,
			row.OccurredAt.UTC().Format(time.RFC3339),
		}
		if err := writer.Write(record); err != nil {
			return nil, fmt.Errorf("write row: %w", err)
		}
	}
	writer.Flush()
	if err := writer.Error(); err != nil {
		return nil, fmt.Errorf("flush csv: %w", err)
	}
	return buf.Bytes(), nil
}

func fromProtoTransactions(items []*ledgerv1.Transaction) []model.Transaction {
	out := make([]model.Transaction, 0, len(items))
	for _, item := range items {
		if item == nil {
			continue
		}
		if tx := fromProtoTransaction(item); tx != nil {
			out = append(out, *tx)
		}
	}
	return out
}

func fromProtoTransaction(item *ledgerv1.Transaction) *model.Transaction {
	if item == nil {
		return nil
	}
	return &model.Transaction{
		ID:          item.GetId(),
		AccountID:   item.GetAccountId(),
		Amount:      item.GetAmount(),
		Currency:    item.GetCurrency(),
		Category:    item.GetCategory(),
		Description: item.GetDescription(),
		OccurredAt:  toTime(item.GetOccurredAt()),
		CreatedAt:   toTime(item.GetCreatedAt()),
		UpdatedAt:   toTime(item.GetUpdatedAt()),
	}
}

func fromProtoBudgets(items []*ledgerv1.Budget) []model.Budget {
	out := make([]model.Budget, 0, len(items))
	for _, item := range items {
		if item == nil {
			continue
		}
		out = append(out, model.Budget{
			ID:        item.GetId(),
			Name:      item.GetName(),
			Amount:    item.GetAmount(),
			Currency:  item.GetCurrency(),
			Period:    item.GetPeriod(),
			StartDate: toTime(item.GetStartDate()),
			EndDate:   toTime(item.GetEndDate()),
			CreatedAt: toTime(item.GetCreatedAt()),
			UpdatedAt: toTime(item.GetUpdatedAt()),
		})
	}
	return out
}

func fromProtoBudget(item *ledgerv1.Budget) *model.Budget {
	if item == nil {
		return nil
	}
	return &model.Budget{
		ID:        item.GetId(),
		Name:      item.GetName(),
		Amount:    item.GetAmount(),
		Currency:  item.GetCurrency(),
		Period:    item.GetPeriod(),
		StartDate: toTime(item.GetStartDate()),
		EndDate:   toTime(item.GetEndDate()),
		CreatedAt: toTime(item.GetCreatedAt()),
		UpdatedAt: toTime(item.GetUpdatedAt()),
	}
}

func fromProtoReports(items []*ledgerv1.Report) []model.Report {
	out := make([]model.Report, 0, len(items))
	for _, item := range items {
		if item == nil {
			continue
		}
		out = append(out, model.Report{
			ID:           item.GetId(),
			Name:         item.GetName(),
			Period:       item.GetPeriod(),
			GeneratedAt:  toTime(item.GetGeneratedAt()),
			TotalIncome:  item.GetTotalIncome(),
			TotalExpense: item.GetTotalExpense(),
			Currency:     item.GetCurrency(),
		})
	}
	return out
}

func fromProtoReport(item *ledgerv1.Report) *model.Report {
	if item == nil {
		return nil
	}
	return &model.Report{
		ID:           item.GetId(),
		Name:         item.GetName(),
		Period:       item.GetPeriod(),
		GeneratedAt:  toTime(item.GetGeneratedAt()),
		TotalIncome:  item.GetTotalIncome(),
		TotalExpense: item.GetTotalExpense(),
		Currency:     item.GetCurrency(),
	}
}

func toTime(ts *timestamppb.Timestamp) time.Time {
	if ts == nil {
		return time.Time{}
	}
	return ts.AsTime()
}
