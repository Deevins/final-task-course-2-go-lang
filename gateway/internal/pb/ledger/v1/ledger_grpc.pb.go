// Code generated manually. DO NOT EDIT.
// source: ledger/v1/ledger.proto

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

const _ = grpc.SupportPackageIsVersion9

const (
	LedgerService_CreateTransaction_FullMethodName     = "/ledger.v1.LedgerService/CreateTransaction"
	LedgerService_GetTransaction_FullMethodName        = "/ledger.v1.LedgerService/GetTransaction"
	LedgerService_UpdateTransaction_FullMethodName     = "/ledger.v1.LedgerService/UpdateTransaction"
	LedgerService_DeleteTransaction_FullMethodName     = "/ledger.v1.LedgerService/DeleteTransaction"
	LedgerService_ListTransactions_FullMethodName      = "/ledger.v1.LedgerService/ListTransactions"
	LedgerService_CreateBudget_FullMethodName          = "/ledger.v1.LedgerService/CreateBudget"
	LedgerService_GetBudget_FullMethodName             = "/ledger.v1.LedgerService/GetBudget"
	LedgerService_UpdateBudget_FullMethodName          = "/ledger.v1.LedgerService/UpdateBudget"
	LedgerService_DeleteBudget_FullMethodName          = "/ledger.v1.LedgerService/DeleteBudget"
	LedgerService_ListBudgets_FullMethodName           = "/ledger.v1.LedgerService/ListBudgets"
	LedgerService_CreateReport_FullMethodName          = "/ledger.v1.LedgerService/CreateReport"
	LedgerService_GetReport_FullMethodName             = "/ledger.v1.LedgerService/GetReport"
	LedgerService_UpdateReport_FullMethodName          = "/ledger.v1.LedgerService/UpdateReport"
	LedgerService_DeleteReport_FullMethodName          = "/ledger.v1.LedgerService/DeleteReport"
	LedgerService_ListReports_FullMethodName           = "/ledger.v1.LedgerService/ListReports"
	LedgerService_ImportTransactionsCsv_FullMethodName = "/ledger.v1.LedgerService/ImportTransactionsCsv"
	LedgerService_ExportTransactionsCsv_FullMethodName = "/ledger.v1.LedgerService/ExportTransactionsCsv"
)

type LedgerServiceClient interface {
	CreateTransaction(ctx context.Context, in *CreateTransactionRequest, opts ...grpc.CallOption) (*TransactionResponse, error)
	GetTransaction(ctx context.Context, in *GetTransactionRequest, opts ...grpc.CallOption) (*TransactionResponse, error)
	UpdateTransaction(ctx context.Context, in *UpdateTransactionRequest, opts ...grpc.CallOption) (*TransactionResponse, error)
	DeleteTransaction(ctx context.Context, in *DeleteTransactionRequest, opts ...grpc.CallOption) (*DeleteResponse, error)
	ListTransactions(ctx context.Context, in *ListTransactionsRequest, opts ...grpc.CallOption) (*ListTransactionsResponse, error)
	CreateBudget(ctx context.Context, in *CreateBudgetRequest, opts ...grpc.CallOption) (*BudgetResponse, error)
	GetBudget(ctx context.Context, in *GetBudgetRequest, opts ...grpc.CallOption) (*BudgetResponse, error)
	UpdateBudget(ctx context.Context, in *UpdateBudgetRequest, opts ...grpc.CallOption) (*BudgetResponse, error)
	DeleteBudget(ctx context.Context, in *DeleteBudgetRequest, opts ...grpc.CallOption) (*DeleteResponse, error)
	ListBudgets(ctx context.Context, in *ListBudgetsRequest, opts ...grpc.CallOption) (*ListBudgetsResponse, error)
	CreateReport(ctx context.Context, in *CreateReportRequest, opts ...grpc.CallOption) (*ReportResponse, error)
	GetReport(ctx context.Context, in *GetReportRequest, opts ...grpc.CallOption) (*ReportResponse, error)
	UpdateReport(ctx context.Context, in *UpdateReportRequest, opts ...grpc.CallOption) (*ReportResponse, error)
	DeleteReport(ctx context.Context, in *DeleteReportRequest, opts ...grpc.CallOption) (*DeleteResponse, error)
	ListReports(ctx context.Context, in *ListReportsRequest, opts ...grpc.CallOption) (*ListReportsResponse, error)
	ImportTransactionsCsv(ctx context.Context, in *ImportTransactionsCsvRequest, opts ...grpc.CallOption) (*ImportTransactionsCsvResponse, error)
	ExportTransactionsCsv(ctx context.Context, in *ExportTransactionsCsvRequest, opts ...grpc.CallOption) (*ExportTransactionsCsvResponse, error)
}

type ledgerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLedgerServiceClient(cc grpc.ClientConnInterface) LedgerServiceClient {
	return &ledgerServiceClient{cc}
}

func (c *ledgerServiceClient) CreateTransaction(ctx context.Context, in *CreateTransactionRequest, opts ...grpc.CallOption) (*TransactionResponse, error) {
	out := new(TransactionResponse)
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	err := c.cc.Invoke(ctx, LedgerService_CreateTransaction_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ledgerServiceClient) GetTransaction(ctx context.Context, in *GetTransactionRequest, opts ...grpc.CallOption) (*TransactionResponse, error) {
	out := new(TransactionResponse)
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	err := c.cc.Invoke(ctx, LedgerService_GetTransaction_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ledgerServiceClient) UpdateTransaction(ctx context.Context, in *UpdateTransactionRequest, opts ...grpc.CallOption) (*TransactionResponse, error) {
	out := new(TransactionResponse)
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	err := c.cc.Invoke(ctx, LedgerService_UpdateTransaction_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ledgerServiceClient) DeleteTransaction(ctx context.Context, in *DeleteTransactionRequest, opts ...grpc.CallOption) (*DeleteResponse, error) {
	out := new(DeleteResponse)
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	err := c.cc.Invoke(ctx, LedgerService_DeleteTransaction_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ledgerServiceClient) ListTransactions(ctx context.Context, in *ListTransactionsRequest, opts ...grpc.CallOption) (*ListTransactionsResponse, error) {
	out := new(ListTransactionsResponse)
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	err := c.cc.Invoke(ctx, LedgerService_ListTransactions_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ledgerServiceClient) CreateBudget(ctx context.Context, in *CreateBudgetRequest, opts ...grpc.CallOption) (*BudgetResponse, error) {
	out := new(BudgetResponse)
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	err := c.cc.Invoke(ctx, LedgerService_CreateBudget_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ledgerServiceClient) GetBudget(ctx context.Context, in *GetBudgetRequest, opts ...grpc.CallOption) (*BudgetResponse, error) {
	out := new(BudgetResponse)
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	err := c.cc.Invoke(ctx, LedgerService_GetBudget_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ledgerServiceClient) UpdateBudget(ctx context.Context, in *UpdateBudgetRequest, opts ...grpc.CallOption) (*BudgetResponse, error) {
	out := new(BudgetResponse)
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	err := c.cc.Invoke(ctx, LedgerService_UpdateBudget_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ledgerServiceClient) DeleteBudget(ctx context.Context, in *DeleteBudgetRequest, opts ...grpc.CallOption) (*DeleteResponse, error) {
	out := new(DeleteResponse)
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	err := c.cc.Invoke(ctx, LedgerService_DeleteBudget_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ledgerServiceClient) ListBudgets(ctx context.Context, in *ListBudgetsRequest, opts ...grpc.CallOption) (*ListBudgetsResponse, error) {
	out := new(ListBudgetsResponse)
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	err := c.cc.Invoke(ctx, LedgerService_ListBudgets_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ledgerServiceClient) CreateReport(ctx context.Context, in *CreateReportRequest, opts ...grpc.CallOption) (*ReportResponse, error) {
	out := new(ReportResponse)
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	err := c.cc.Invoke(ctx, LedgerService_CreateReport_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ledgerServiceClient) GetReport(ctx context.Context, in *GetReportRequest, opts ...grpc.CallOption) (*ReportResponse, error) {
	out := new(ReportResponse)
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	err := c.cc.Invoke(ctx, LedgerService_GetReport_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ledgerServiceClient) UpdateReport(ctx context.Context, in *UpdateReportRequest, opts ...grpc.CallOption) (*ReportResponse, error) {
	out := new(ReportResponse)
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	err := c.cc.Invoke(ctx, LedgerService_UpdateReport_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ledgerServiceClient) DeleteReport(ctx context.Context, in *DeleteReportRequest, opts ...grpc.CallOption) (*DeleteResponse, error) {
	out := new(DeleteResponse)
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	err := c.cc.Invoke(ctx, LedgerService_DeleteReport_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ledgerServiceClient) ListReports(ctx context.Context, in *ListReportsRequest, opts ...grpc.CallOption) (*ListReportsResponse, error) {
	out := new(ListReportsResponse)
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	err := c.cc.Invoke(ctx, LedgerService_ListReports_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ledgerServiceClient) ImportTransactionsCsv(ctx context.Context, in *ImportTransactionsCsvRequest, opts ...grpc.CallOption) (*ImportTransactionsCsvResponse, error) {
	out := new(ImportTransactionsCsvResponse)
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	err := c.cc.Invoke(ctx, LedgerService_ImportTransactionsCsv_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ledgerServiceClient) ExportTransactionsCsv(ctx context.Context, in *ExportTransactionsCsvRequest, opts ...grpc.CallOption) (*ExportTransactionsCsvResponse, error) {
	out := new(ExportTransactionsCsvResponse)
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	err := c.cc.Invoke(ctx, LedgerService_ExportTransactionsCsv_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LedgerServiceServer is the server API for LedgerService service.
// All implementations must embed UnimplementedLedgerServiceServer for forward compatibility.
type LedgerServiceServer interface {
	CreateTransaction(context.Context, *CreateTransactionRequest) (*TransactionResponse, error)
	GetTransaction(context.Context, *GetTransactionRequest) (*TransactionResponse, error)
	UpdateTransaction(context.Context, *UpdateTransactionRequest) (*TransactionResponse, error)
	DeleteTransaction(context.Context, *DeleteTransactionRequest) (*DeleteResponse, error)
	ListTransactions(context.Context, *ListTransactionsRequest) (*ListTransactionsResponse, error)
	CreateBudget(context.Context, *CreateBudgetRequest) (*BudgetResponse, error)
	GetBudget(context.Context, *GetBudgetRequest) (*BudgetResponse, error)
	UpdateBudget(context.Context, *UpdateBudgetRequest) (*BudgetResponse, error)
	DeleteBudget(context.Context, *DeleteBudgetRequest) (*DeleteResponse, error)
	ListBudgets(context.Context, *ListBudgetsRequest) (*ListBudgetsResponse, error)
	CreateReport(context.Context, *CreateReportRequest) (*ReportResponse, error)
	GetReport(context.Context, *GetReportRequest) (*ReportResponse, error)
	UpdateReport(context.Context, *UpdateReportRequest) (*ReportResponse, error)
	DeleteReport(context.Context, *DeleteReportRequest) (*DeleteResponse, error)
	ListReports(context.Context, *ListReportsRequest) (*ListReportsResponse, error)
	ImportTransactionsCsv(context.Context, *ImportTransactionsCsvRequest) (*ImportTransactionsCsvResponse, error)
	ExportTransactionsCsv(context.Context, *ExportTransactionsCsvRequest) (*ExportTransactionsCsvResponse, error)
	mustEmbedUnimplementedLedgerServiceServer()
}

type UnimplementedLedgerServiceServer struct{}

func (UnimplementedLedgerServiceServer) CreateTransaction(context.Context, *CreateTransactionRequest) (*TransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTransaction not implemented")
}
func (UnimplementedLedgerServiceServer) GetTransaction(context.Context, *GetTransactionRequest) (*TransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTransaction not implemented")
}
func (UnimplementedLedgerServiceServer) UpdateTransaction(context.Context, *UpdateTransactionRequest) (*TransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTransaction not implemented")
}
func (UnimplementedLedgerServiceServer) DeleteTransaction(context.Context, *DeleteTransactionRequest) (*DeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTransaction not implemented")
}
func (UnimplementedLedgerServiceServer) ListTransactions(context.Context, *ListTransactionsRequest) (*ListTransactionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListTransactions not implemented")
}
func (UnimplementedLedgerServiceServer) CreateBudget(context.Context, *CreateBudgetRequest) (*BudgetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBudget not implemented")
}
func (UnimplementedLedgerServiceServer) GetBudget(context.Context, *GetBudgetRequest) (*BudgetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBudget not implemented")
}
func (UnimplementedLedgerServiceServer) UpdateBudget(context.Context, *UpdateBudgetRequest) (*BudgetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateBudget not implemented")
}
func (UnimplementedLedgerServiceServer) DeleteBudget(context.Context, *DeleteBudgetRequest) (*DeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteBudget not implemented")
}
func (UnimplementedLedgerServiceServer) ListBudgets(context.Context, *ListBudgetsRequest) (*ListBudgetsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListBudgets not implemented")
}
func (UnimplementedLedgerServiceServer) CreateReport(context.Context, *CreateReportRequest) (*ReportResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateReport not implemented")
}
func (UnimplementedLedgerServiceServer) GetReport(context.Context, *GetReportRequest) (*ReportResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetReport not implemented")
}
func (UnimplementedLedgerServiceServer) UpdateReport(context.Context, *UpdateReportRequest) (*ReportResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateReport not implemented")
}
func (UnimplementedLedgerServiceServer) DeleteReport(context.Context, *DeleteReportRequest) (*DeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteReport not implemented")
}
func (UnimplementedLedgerServiceServer) ListReports(context.Context, *ListReportsRequest) (*ListReportsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListReports not implemented")
}
func (UnimplementedLedgerServiceServer) ImportTransactionsCsv(context.Context, *ImportTransactionsCsvRequest) (*ImportTransactionsCsvResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ImportTransactionsCsv not implemented")
}
func (UnimplementedLedgerServiceServer) ExportTransactionsCsv(context.Context, *ExportTransactionsCsvRequest) (*ExportTransactionsCsvResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExportTransactionsCsv not implemented")
}
func (UnimplementedLedgerServiceServer) mustEmbedUnimplementedLedgerServiceServer() {}
func (UnimplementedLedgerServiceServer) testEmbeddedByValue()                       {}

type UnsafeLedgerServiceServer interface {
	mustEmbedUnimplementedLedgerServiceServer()
}

func RegisterLedgerServiceServer(s grpc.ServiceRegistrar, srv LedgerServiceServer) {
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&LedgerService_ServiceDesc, srv)
}

func _LedgerService_CreateTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTransactionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LedgerServiceServer).CreateTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{Server: srv, FullMethod: LedgerService_CreateTransaction_FullMethodName}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LedgerServiceServer).CreateTransaction(ctx, req.(*CreateTransactionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LedgerService_GetTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTransactionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LedgerServiceServer).GetTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{Server: srv, FullMethod: LedgerService_GetTransaction_FullMethodName}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LedgerServiceServer).GetTransaction(ctx, req.(*GetTransactionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LedgerService_UpdateTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateTransactionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LedgerServiceServer).UpdateTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{Server: srv, FullMethod: LedgerService_UpdateTransaction_FullMethodName}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LedgerServiceServer).UpdateTransaction(ctx, req.(*UpdateTransactionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LedgerService_DeleteTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteTransactionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LedgerServiceServer).DeleteTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{Server: srv, FullMethod: LedgerService_DeleteTransaction_FullMethodName}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LedgerServiceServer).DeleteTransaction(ctx, req.(*DeleteTransactionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LedgerService_ListTransactions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListTransactionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LedgerServiceServer).ListTransactions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{Server: srv, FullMethod: LedgerService_ListTransactions_FullMethodName}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LedgerServiceServer).ListTransactions(ctx, req.(*ListTransactionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LedgerService_CreateBudget_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateBudgetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LedgerServiceServer).CreateBudget(ctx, in)
	}
	info := &grpc.UnaryServerInfo{Server: srv, FullMethod: LedgerService_CreateBudget_FullMethodName}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LedgerServiceServer).CreateBudget(ctx, req.(*CreateBudgetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LedgerService_GetBudget_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBudgetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LedgerServiceServer).GetBudget(ctx, in)
	}
	info := &grpc.UnaryServerInfo{Server: srv, FullMethod: LedgerService_GetBudget_FullMethodName}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LedgerServiceServer).GetBudget(ctx, req.(*GetBudgetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LedgerService_UpdateBudget_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateBudgetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LedgerServiceServer).UpdateBudget(ctx, in)
	}
	info := &grpc.UnaryServerInfo{Server: srv, FullMethod: LedgerService_UpdateBudget_FullMethodName}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LedgerServiceServer).UpdateBudget(ctx, req.(*UpdateBudgetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LedgerService_DeleteBudget_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteBudgetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LedgerServiceServer).DeleteBudget(ctx, in)
	}
	info := &grpc.UnaryServerInfo{Server: srv, FullMethod: LedgerService_DeleteBudget_FullMethodName}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LedgerServiceServer).DeleteBudget(ctx, req.(*DeleteBudgetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LedgerService_ListBudgets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListBudgetsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LedgerServiceServer).ListBudgets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{Server: srv, FullMethod: LedgerService_ListBudgets_FullMethodName}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LedgerServiceServer).ListBudgets(ctx, req.(*ListBudgetsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LedgerService_CreateReport_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateReportRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LedgerServiceServer).CreateReport(ctx, in)
	}
	info := &grpc.UnaryServerInfo{Server: srv, FullMethod: LedgerService_CreateReport_FullMethodName}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LedgerServiceServer).CreateReport(ctx, req.(*CreateReportRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LedgerService_GetReport_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetReportRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LedgerServiceServer).GetReport(ctx, in)
	}
	info := &grpc.UnaryServerInfo{Server: srv, FullMethod: LedgerService_GetReport_FullMethodName}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LedgerServiceServer).GetReport(ctx, req.(*GetReportRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LedgerService_UpdateReport_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateReportRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LedgerServiceServer).UpdateReport(ctx, in)
	}
	info := &grpc.UnaryServerInfo{Server: srv, FullMethod: LedgerService_UpdateReport_FullMethodName}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LedgerServiceServer).UpdateReport(ctx, req.(*UpdateReportRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LedgerService_DeleteReport_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteReportRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LedgerServiceServer).DeleteReport(ctx, in)
	}
	info := &grpc.UnaryServerInfo{Server: srv, FullMethod: LedgerService_DeleteReport_FullMethodName}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LedgerServiceServer).DeleteReport(ctx, req.(*DeleteReportRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LedgerService_ListReports_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListReportsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LedgerServiceServer).ListReports(ctx, in)
	}
	info := &grpc.UnaryServerInfo{Server: srv, FullMethod: LedgerService_ListReports_FullMethodName}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LedgerServiceServer).ListReports(ctx, req.(*ListReportsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LedgerService_ImportTransactionsCsv_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ImportTransactionsCsvRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LedgerServiceServer).ImportTransactionsCsv(ctx, in)
	}
	info := &grpc.UnaryServerInfo{Server: srv, FullMethod: LedgerService_ImportTransactionsCsv_FullMethodName}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LedgerServiceServer).ImportTransactionsCsv(ctx, req.(*ImportTransactionsCsvRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LedgerService_ExportTransactionsCsv_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExportTransactionsCsvRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LedgerServiceServer).ExportTransactionsCsv(ctx, in)
	}
	info := &grpc.UnaryServerInfo{Server: srv, FullMethod: LedgerService_ExportTransactionsCsv_FullMethodName}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LedgerServiceServer).ExportTransactionsCsv(ctx, req.(*ExportTransactionsCsvRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var LedgerService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ledger.v1.LedgerService",
	HandlerType: (*LedgerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{MethodName: "CreateTransaction", Handler: _LedgerService_CreateTransaction_Handler},
		{MethodName: "GetTransaction", Handler: _LedgerService_GetTransaction_Handler},
		{MethodName: "UpdateTransaction", Handler: _LedgerService_UpdateTransaction_Handler},
		{MethodName: "DeleteTransaction", Handler: _LedgerService_DeleteTransaction_Handler},
		{MethodName: "ListTransactions", Handler: _LedgerService_ListTransactions_Handler},
		{MethodName: "CreateBudget", Handler: _LedgerService_CreateBudget_Handler},
		{MethodName: "GetBudget", Handler: _LedgerService_GetBudget_Handler},
		{MethodName: "UpdateBudget", Handler: _LedgerService_UpdateBudget_Handler},
		{MethodName: "DeleteBudget", Handler: _LedgerService_DeleteBudget_Handler},
		{MethodName: "ListBudgets", Handler: _LedgerService_ListBudgets_Handler},
		{MethodName: "CreateReport", Handler: _LedgerService_CreateReport_Handler},
		{MethodName: "GetReport", Handler: _LedgerService_GetReport_Handler},
		{MethodName: "UpdateReport", Handler: _LedgerService_UpdateReport_Handler},
		{MethodName: "DeleteReport", Handler: _LedgerService_DeleteReport_Handler},
		{MethodName: "ListReports", Handler: _LedgerService_ListReports_Handler},
		{MethodName: "ImportTransactionsCsv", Handler: _LedgerService_ImportTransactionsCsv_Handler},
		{MethodName: "ExportTransactionsCsv", Handler: _LedgerService_ExportTransactionsCsv_Handler},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ledger/v1/ledger.proto",
}
