// Code generated manually. DO NOT EDIT.
// source: ledger/v1/ledger.proto

package v1

import (
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Transaction struct {
	Id          string
	AccountId   string
	Amount      float64
	Currency    string
	Category    string
	Description string
	OccurredAt  *timestamppb.Timestamp
	CreatedAt   *timestamppb.Timestamp
	UpdatedAt   *timestamppb.Timestamp
}

func (x *Transaction) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Transaction) GetAccountId() string {
	if x != nil {
		return x.AccountId
	}
	return ""
}

func (x *Transaction) GetAmount() float64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *Transaction) GetCurrency() string {
	if x != nil {
		return x.Currency
	}
	return ""
}

func (x *Transaction) GetCategory() string {
	if x != nil {
		return x.Category
	}
	return ""
}

func (x *Transaction) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Transaction) GetOccurredAt() *timestamppb.Timestamp {
	if x != nil {
		return x.OccurredAt
	}
	return nil
}

func (x *Transaction) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Transaction) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

type Budget struct {
	Id        string
	Name      string
	Amount    float64
	Currency  string
	Period    string
	StartDate *timestamppb.Timestamp
	EndDate   *timestamppb.Timestamp
	CreatedAt *timestamppb.Timestamp
	UpdatedAt *timestamppb.Timestamp
}

func (x *Budget) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Budget) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Budget) GetAmount() float64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *Budget) GetCurrency() string {
	if x != nil {
		return x.Currency
	}
	return ""
}

func (x *Budget) GetPeriod() string {
	if x != nil {
		return x.Period
	}
	return ""
}

func (x *Budget) GetStartDate() *timestamppb.Timestamp {
	if x != nil {
		return x.StartDate
	}
	return nil
}

func (x *Budget) GetEndDate() *timestamppb.Timestamp {
	if x != nil {
		return x.EndDate
	}
	return nil
}

func (x *Budget) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Budget) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

type Report struct {
	Id           string
	Name         string
	Period       string
	GeneratedAt  *timestamppb.Timestamp
	TotalIncome  float64
	TotalExpense float64
	Currency     string
}

func (x *Report) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Report) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Report) GetPeriod() string {
	if x != nil {
		return x.Period
	}
	return ""
}

func (x *Report) GetGeneratedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.GeneratedAt
	}
	return nil
}

func (x *Report) GetTotalIncome() float64 {
	if x != nil {
		return x.TotalIncome
	}
	return 0
}

func (x *Report) GetTotalExpense() float64 {
	if x != nil {
		return x.TotalExpense
	}
	return 0
}

func (x *Report) GetCurrency() string {
	if x != nil {
		return x.Currency
	}
	return ""
}

type CreateTransactionRequest struct {
	Transaction *Transaction
}

func (x *CreateTransactionRequest) GetTransaction() *Transaction {
	if x != nil {
		return x.Transaction
	}
	return nil
}

type GetTransactionRequest struct {
	Id string
}

func (x *GetTransactionRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type UpdateTransactionRequest struct {
	Transaction *Transaction
}

func (x *UpdateTransactionRequest) GetTransaction() *Transaction {
	if x != nil {
		return x.Transaction
	}
	return nil
}

type DeleteTransactionRequest struct {
	Id string
}

func (x *DeleteTransactionRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type ListTransactionsRequest struct {
	AccountId string
	Limit     int32
	PageToken string
}

func (x *ListTransactionsRequest) GetAccountId() string {
	if x != nil {
		return x.AccountId
	}
	return ""
}

func (x *ListTransactionsRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *ListTransactionsRequest) GetPageToken() string {
	if x != nil {
		return x.PageToken
	}
	return ""
}

type ListTransactionsResponse struct {
	Transactions  []*Transaction
	NextPageToken string
}

func (x *ListTransactionsResponse) GetTransactions() []*Transaction {
	if x != nil {
		return x.Transactions
	}
	return nil
}

func (x *ListTransactionsResponse) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

type TransactionResponse struct {
	Transaction *Transaction
}

func (x *TransactionResponse) GetTransaction() *Transaction {
	if x != nil {
		return x.Transaction
	}
	return nil
}

type DeleteResponse struct {
	Deleted bool
}

func (x *DeleteResponse) GetDeleted() bool {
	if x != nil {
		return x.Deleted
	}
	return false
}

type CreateBudgetRequest struct {
	Budget *Budget
}

func (x *CreateBudgetRequest) GetBudget() *Budget {
	if x != nil {
		return x.Budget
	}
	return nil
}

type GetBudgetRequest struct {
	Id string
}

func (x *GetBudgetRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type UpdateBudgetRequest struct {
	Budget *Budget
}

func (x *UpdateBudgetRequest) GetBudget() *Budget {
	if x != nil {
		return x.Budget
	}
	return nil
}

type DeleteBudgetRequest struct {
	Id string
}

func (x *DeleteBudgetRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type ListBudgetsRequest struct {
	Limit     int32
	PageToken string
}

func (x *ListBudgetsRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *ListBudgetsRequest) GetPageToken() string {
	if x != nil {
		return x.PageToken
	}
	return ""
}

type ListBudgetsResponse struct {
	Budgets       []*Budget
	NextPageToken string
}

func (x *ListBudgetsResponse) GetBudgets() []*Budget {
	if x != nil {
		return x.Budgets
	}
	return nil
}

func (x *ListBudgetsResponse) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

type BudgetResponse struct {
	Budget *Budget
}

func (x *BudgetResponse) GetBudget() *Budget {
	if x != nil {
		return x.Budget
	}
	return nil
}

type CreateReportRequest struct {
	Report *Report
}

func (x *CreateReportRequest) GetReport() *Report {
	if x != nil {
		return x.Report
	}
	return nil
}

type GetReportRequest struct {
	Id string
}

func (x *GetReportRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type UpdateReportRequest struct {
	Report *Report
}

func (x *UpdateReportRequest) GetReport() *Report {
	if x != nil {
		return x.Report
	}
	return nil
}

type DeleteReportRequest struct {
	Id string
}

func (x *DeleteReportRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type ListReportsRequest struct {
	Limit     int32
	PageToken string
}

func (x *ListReportsRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *ListReportsRequest) GetPageToken() string {
	if x != nil {
		return x.PageToken
	}
	return ""
}

type ListReportsResponse struct {
	Reports       []*Report
	NextPageToken string
}

func (x *ListReportsResponse) GetReports() []*Report {
	if x != nil {
		return x.Reports
	}
	return nil
}

func (x *ListReportsResponse) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

type ReportResponse struct {
	Report *Report
}

func (x *ReportResponse) GetReport() *Report {
	if x != nil {
		return x.Report
	}
	return nil
}

type ImportTransactionsCsvRequest struct {
	CsvContent []byte
	HasHeader  bool
}

func (x *ImportTransactionsCsvRequest) GetCsvContent() []byte {
	if x != nil {
		return x.CsvContent
	}
	return nil
}

func (x *ImportTransactionsCsvRequest) GetHasHeader() bool {
	if x != nil {
		return x.HasHeader
	}
	return false
}

type ImportTransactionsCsvResponse struct {
	Imported int32
}

func (x *ImportTransactionsCsvResponse) GetImported() int32 {
	if x != nil {
		return x.Imported
	}
	return 0
}

type ExportTransactionsCsvRequest struct {
	AccountId string
}

func (x *ExportTransactionsCsvRequest) GetAccountId() string {
	if x != nil {
		return x.AccountId
	}
	return ""
}

type ExportTransactionsCsvResponse struct {
	CsvContent []byte
}

func (x *ExportTransactionsCsvResponse) GetCsvContent() []byte {
	if x != nil {
		return x.CsvContent
	}
	return nil
}
