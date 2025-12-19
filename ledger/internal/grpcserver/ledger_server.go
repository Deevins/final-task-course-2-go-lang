package grpcserver

import (
	"context"
	"time"

	"github.com/Deevins/final-task-course-2-go-lang/ledger/internal/model"
	pb "github.com/Deevins/final-task-course-2-go-lang/ledger/internal/pb/ledger/v1"
	"github.com/Deevins/final-task-course-2-go-lang/ledger/internal/service"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type LedgerServer struct {
	pb.UnimplementedLedgerServiceServer
	ledgerService service.LedgerService
}

var _ pb.LedgerServiceServer = (*LedgerServer)(nil)

func NewLedgerServer(svc service.LedgerService) *LedgerServer {
	return &LedgerServer{ledgerService: svc}
}

func (s *LedgerServer) CreateTransaction(ctx context.Context, req *pb.CreateTransactionRequest) (*pb.TransactionResponse, error) {
	if req.GetTransaction() == nil {
		return nil, status.Error(codes.InvalidArgument, "transaction is required")
	}

	created, err := s.ledgerService.CreateTransaction(ctx, toModelTransaction(req.GetTransaction()))
	if err != nil {
		if service.IsValidationError(err) {
			return nil, status.Errorf(codes.InvalidArgument, "create transaction: %v", err)
		}
		if service.IsBudgetExceeded(err) {
			return nil, status.Errorf(codes.FailedPrecondition, "create transaction: %v", err)
		}
		return nil, status.Errorf(codes.Internal, "create transaction: %v", err)
	}

	return &pb.TransactionResponse{Transaction: toProtoTransaction(created)}, nil
}

func (s *LedgerServer) GetTransaction(ctx context.Context, req *pb.GetTransactionRequest) (*pb.TransactionResponse, error) {
	if req.GetId() == "" {
		return nil, status.Error(codes.InvalidArgument, "id is required")
	}

	transaction, err := s.ledgerService.GetTransaction(ctx, req.GetId())
	if err != nil {
		if service.IsNotFound(err) {
			return nil, status.Error(codes.NotFound, "transaction not found")
		}
		return nil, status.Errorf(codes.Internal, "get transaction: %v", err)
	}

	return &pb.TransactionResponse{Transaction: toProtoTransaction(transaction)}, nil
}

func (s *LedgerServer) UpdateTransaction(ctx context.Context, req *pb.UpdateTransactionRequest) (*pb.TransactionResponse, error) {
	if req.GetTransaction() == nil {
		return nil, status.Error(codes.InvalidArgument, "transaction is required")
	}

	updated, err := s.ledgerService.UpdateTransaction(ctx, toModelTransaction(req.GetTransaction()))
	if err != nil {
		if service.IsNotFound(err) {
			return nil, status.Error(codes.NotFound, "transaction not found")
		}
		if service.IsValidationError(err) {
			return nil, status.Errorf(codes.InvalidArgument, "update transaction: %v", err)
		}
		return nil, status.Errorf(codes.Internal, "update transaction: %v", err)
	}

	return &pb.TransactionResponse{Transaction: toProtoTransaction(updated)}, nil
}

func (s *LedgerServer) DeleteTransaction(ctx context.Context, req *pb.DeleteTransactionRequest) (*pb.DeleteResponse, error) {
	if req.GetId() == "" {
		return nil, status.Error(codes.InvalidArgument, "id is required")
	}

	if err := s.ledgerService.DeleteTransaction(ctx, req.GetId()); err != nil {
		if service.IsNotFound(err) {
			return &pb.DeleteResponse{Deleted: false}, nil
		}
		if service.IsValidationError(err) {
			return nil, status.Errorf(codes.InvalidArgument, "delete transaction: %v", err)
		}
		return nil, status.Errorf(codes.Internal, "delete transaction: %v", err)
	}

	return &pb.DeleteResponse{Deleted: true}, nil
}

func (s *LedgerServer) ListTransactions(ctx context.Context, req *pb.ListTransactionsRequest) (*pb.ListTransactionsResponse, error) {
	items := s.ledgerService.ListTransactions(ctx, req.GetAccountId())
	resp := &pb.ListTransactionsResponse{}
	resp.Transactions = make([]*pb.Transaction, 0, len(items))
	for _, tx := range items {
		resp.Transactions = append(resp.Transactions, toProtoTransaction(tx))
	}
	return resp, nil
}

func (s *LedgerServer) CreateBudget(ctx context.Context, req *pb.CreateBudgetRequest) (*pb.BudgetResponse, error) {
	if req.GetBudget() == nil {
		return nil, status.Error(codes.InvalidArgument, "budget is required")
	}

	created, err := s.ledgerService.CreateBudget(ctx, toModelBudget(req.GetBudget()))
	if err != nil {
		if service.IsValidationError(err) {
			return nil, status.Errorf(codes.InvalidArgument, "create budget: %v", err)
		}
		return nil, status.Errorf(codes.Internal, "create budget: %v", err)
	}

	return &pb.BudgetResponse{Budget: toProtoBudget(created)}, nil
}

func (s *LedgerServer) GetBudget(ctx context.Context, req *pb.GetBudgetRequest) (*pb.BudgetResponse, error) {
	if req.GetId() == "" {
		return nil, status.Error(codes.InvalidArgument, "id is required")
	}
	if req.GetAccountId() == "" {
		return nil, status.Error(codes.InvalidArgument, "account_id is required")
	}

	budget, err := s.ledgerService.GetBudget(ctx, req.GetAccountId(), req.GetId())
	if err != nil {
		if service.IsNotFound(err) {
			return nil, status.Error(codes.NotFound, "budget not found")
		}
		if service.IsValidationError(err) {
			return nil, status.Errorf(codes.InvalidArgument, "get budget: %v", err)
		}
		return nil, status.Errorf(codes.Internal, "get budget: %v", err)
	}

	return &pb.BudgetResponse{Budget: toProtoBudget(budget)}, nil
}

func (s *LedgerServer) UpdateBudget(ctx context.Context, req *pb.UpdateBudgetRequest) (*pb.BudgetResponse, error) {
	if req.GetBudget() == nil {
		return nil, status.Error(codes.InvalidArgument, "budget is required")
	}

	updated, err := s.ledgerService.UpdateBudget(ctx, toModelBudget(req.GetBudget()))
	if err != nil {
		if service.IsNotFound(err) {
			return nil, status.Error(codes.NotFound, "budget not found")
		}
		if service.IsValidationError(err) {
			return nil, status.Errorf(codes.InvalidArgument, "update budget: %v", err)
		}
		return nil, status.Errorf(codes.Internal, "update budget: %v", err)
	}

	return &pb.BudgetResponse{Budget: toProtoBudget(updated)}, nil
}

func (s *LedgerServer) DeleteBudget(ctx context.Context, req *pb.DeleteBudgetRequest) (*pb.DeleteResponse, error) {
	if req.GetId() == "" {
		return nil, status.Error(codes.InvalidArgument, "id is required")
	}
	if req.GetAccountId() == "" {
		return nil, status.Error(codes.InvalidArgument, "account_id is required")
	}

	if err := s.ledgerService.DeleteBudget(ctx, req.GetAccountId(), req.GetId()); err != nil {
		if service.IsNotFound(err) {
			return &pb.DeleteResponse{Deleted: false}, nil
		}
		if service.IsValidationError(err) {
			return nil, status.Errorf(codes.InvalidArgument, "delete budget: %v", err)
		}
		return nil, status.Errorf(codes.Internal, "delete budget: %v", err)
	}

	return &pb.DeleteResponse{Deleted: true}, nil
}

func (s *LedgerServer) ListBudgets(ctx context.Context, req *pb.ListBudgetsRequest) (*pb.ListBudgetsResponse, error) {
	if req.GetAccountId() == "" {
		return nil, status.Error(codes.InvalidArgument, "account_id is required")
	}
	items := s.ledgerService.ListBudgets(ctx, req.GetAccountId())
	resp := &pb.ListBudgetsResponse{}
	resp.Budgets = make([]*pb.Budget, 0, len(items))
	for _, budget := range items {
		resp.Budgets = append(resp.Budgets, toProtoBudget(budget))
	}
	return resp, nil
}

func (s *LedgerServer) CreateReport(ctx context.Context, req *pb.CreateReportRequest) (*pb.ReportResponse, error) {
	if req.GetReport() == nil {
		return nil, status.Error(codes.InvalidArgument, "report is required")
	}

	created, err := s.ledgerService.CreateReport(ctx, toModelReport(req.GetReport()))
	if err != nil {
		if service.IsValidationError(err) {
			return nil, status.Errorf(codes.InvalidArgument, "create report: %v", err)
		}
		return nil, status.Errorf(codes.Internal, "create report: %v", err)
	}

	return &pb.ReportResponse{Report: toProtoReport(created)}, nil
}

func (s *LedgerServer) GetReport(ctx context.Context, req *pb.GetReportRequest) (*pb.ReportResponse, error) {
	if req.GetId() == "" {
		return nil, status.Error(codes.InvalidArgument, "id is required")
	}
	if req.GetAccountId() == "" {
		return nil, status.Error(codes.InvalidArgument, "account_id is required")
	}

	report, err := s.ledgerService.GetReport(ctx, req.GetAccountId(), req.GetId())
	if err != nil {
		if service.IsNotFound(err) {
			return nil, status.Error(codes.NotFound, "report not found")
		}
		if service.IsValidationError(err) {
			return nil, status.Errorf(codes.InvalidArgument, "get report: %v", err)
		}
		return nil, status.Errorf(codes.Internal, "get report: %v", err)
	}

	return &pb.ReportResponse{Report: toProtoReport(report)}, nil
}

func (s *LedgerServer) UpdateReport(ctx context.Context, req *pb.UpdateReportRequest) (*pb.ReportResponse, error) {
	if req.GetReport() == nil {
		return nil, status.Error(codes.InvalidArgument, "report is required")
	}

	updated, err := s.ledgerService.UpdateReport(ctx, toModelReport(req.GetReport()))
	if err != nil {
		if service.IsNotFound(err) {
			return nil, status.Error(codes.NotFound, "report not found")
		}
		if service.IsValidationError(err) {
			return nil, status.Errorf(codes.InvalidArgument, "update report: %v", err)
		}
		return nil, status.Errorf(codes.Internal, "update report: %v", err)
	}

	return &pb.ReportResponse{Report: toProtoReport(updated)}, nil
}

func (s *LedgerServer) DeleteReport(ctx context.Context, req *pb.DeleteReportRequest) (*pb.DeleteResponse, error) {
	if req.GetId() == "" {
		return nil, status.Error(codes.InvalidArgument, "id is required")
	}
	if req.GetAccountId() == "" {
		return nil, status.Error(codes.InvalidArgument, "account_id is required")
	}

	if err := s.ledgerService.DeleteReport(ctx, req.GetAccountId(), req.GetId()); err != nil {
		if service.IsNotFound(err) {
			return &pb.DeleteResponse{Deleted: false}, nil
		}
		if service.IsValidationError(err) {
			return nil, status.Errorf(codes.InvalidArgument, "delete report: %v", err)
		}
		return nil, status.Errorf(codes.Internal, "delete report: %v", err)
	}

	return &pb.DeleteResponse{Deleted: true}, nil
}

func (s *LedgerServer) ListReports(ctx context.Context, req *pb.ListReportsRequest) (*pb.ListReportsResponse, error) {
	if req.GetAccountId() == "" {
		return nil, status.Error(codes.InvalidArgument, "account_id is required")
	}
	items := s.ledgerService.ListReports(ctx, req.GetAccountId())
	resp := &pb.ListReportsResponse{}
	resp.Reports = make([]*pb.Report, 0, len(items))
	for _, report := range items {
		resp.Reports = append(resp.Reports, toProtoReport(report))
	}
	return resp, nil
}

func (s *LedgerServer) ImportTransactionsCsv(ctx context.Context, req *pb.ImportTransactionsCsvRequest) (*pb.ImportTransactionsCsvResponse, error) {
	if req.GetAccountId() == "" {
		return nil, status.Error(codes.InvalidArgument, "account_id is required")
	}
	if len(req.GetCsvContent()) == 0 {
		return nil, status.Error(codes.InvalidArgument, "csv_content is required")
	}

	count, err := s.ledgerService.ImportTransactionsCSV(ctx, req.GetAccountId(), req.GetCsvContent(), req.GetHasHeader())
	if err != nil {
		if service.IsValidationError(err) {
			return nil, status.Errorf(codes.InvalidArgument, "import csv: %v", err)
		}
		return nil, status.Errorf(codes.Internal, "import csv: %v", err)
	}

	return &pb.ImportTransactionsCsvResponse{Imported: int32(count)}, nil
}

func (s *LedgerServer) ExportTransactionsCsv(ctx context.Context, req *pb.ExportTransactionsCsvRequest) (*pb.ExportTransactionsCsvResponse, error) {
	csvContent, err := s.ledgerService.ExportTransactionsCSV(ctx, req.GetAccountId())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "export csv: %v", err)
	}

	return &pb.ExportTransactionsCsvResponse{CsvContent: csvContent}, nil
}

func toModelTransaction(tx *pb.Transaction) model.Transaction {
	return model.Transaction{
		ID:          tx.GetId(),
		AccountID:   tx.GetAccountId(),
		Amount:      tx.GetAmount(),
		Currency:    tx.GetCurrency(),
		Category:    tx.GetCategory(),
		Description: tx.GetDescription(),
		OccurredAt:  toTime(tx.GetOccurredAt()),
		CreatedAt:   toTime(tx.GetCreatedAt()),
		UpdatedAt:   toTime(tx.GetUpdatedAt()),
	}
}

func toProtoTransaction(tx model.Transaction) *pb.Transaction {
	return &pb.Transaction{
		Id:          tx.ID,
		AccountId:   tx.AccountID,
		Amount:      tx.Amount,
		Currency:    tx.Currency,
		Category:    tx.Category,
		Description: tx.Description,
		OccurredAt:  timestamppb.New(tx.OccurredAt),
		CreatedAt:   timestamppb.New(tx.CreatedAt),
		UpdatedAt:   timestamppb.New(tx.UpdatedAt),
	}
}

func toModelBudget(budget *pb.Budget) model.Budget {
	return model.Budget{
		ID:        budget.GetId(),
		AccountID: budget.GetAccountId(),
		Name:      budget.GetName(),
		Amount:    budget.GetAmount(),
		Currency:  budget.GetCurrency(),
		Period:    budget.GetPeriod(),
		StartDate: toTime(budget.GetStartDate()),
		EndDate:   toTime(budget.GetEndDate()),
		CreatedAt: toTime(budget.GetCreatedAt()),
		UpdatedAt: toTime(budget.GetUpdatedAt()),
	}
}

func toProtoBudget(budget model.Budget) *pb.Budget {
	return &pb.Budget{
		Id:        budget.ID,
		AccountId: budget.AccountID,
		Name:      budget.Name,
		Amount:    budget.Amount,
		Currency:  budget.Currency,
		Period:    budget.Period,
		StartDate: timestamppb.New(budget.StartDate),
		EndDate:   timestamppb.New(budget.EndDate),
		CreatedAt: timestamppb.New(budget.CreatedAt),
		UpdatedAt: timestamppb.New(budget.UpdatedAt),
	}
}

func toModelReport(report *pb.Report) model.Report {
	return model.Report{
		ID:           report.GetId(),
		AccountID:    report.GetAccountId(),
		Name:         report.GetName(),
		Period:       report.GetPeriod(),
		GeneratedAt:  toTime(report.GetGeneratedAt()),
		TotalIncome:  report.GetTotalIncome(),
		TotalExpense: report.GetTotalExpense(),
		Currency:     report.GetCurrency(),
	}
}

func toProtoReport(report model.Report) *pb.Report {
	return &pb.Report{
		Id:           report.ID,
		AccountId:    report.AccountID,
		Name:         report.Name,
		Period:       report.Period,
		GeneratedAt:  timestamppb.New(report.GeneratedAt),
		TotalIncome:  report.TotalIncome,
		TotalExpense: report.TotalExpense,
		Currency:     report.Currency,
	}
}

func toTime(ts *timestamppb.Timestamp) time.Time {
	if ts == nil {
		return time.Time{}
	}
	return ts.AsTime()
}
