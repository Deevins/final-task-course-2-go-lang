package handler

import (
	"net/http"

	"github.com/Deevins/final-task-course-2-go-lang/gateway/internal/middleware"
	"github.com/Deevins/final-task-course-2-go-lang/gateway/internal/model"
	"github.com/Deevins/final-task-course-2-go-lang/gateway/internal/service"
	"github.com/gin-gonic/gin"
)

type LedgerHandler struct {
	service service.LedgerGatewayService
}

func NewLedgerHandler(s service.LedgerGatewayService) *LedgerHandler {
	if s == nil {
		panic("LedgerHandler requires service")
	}
	return &LedgerHandler{service: s}
}

func (h *LedgerHandler) Register(r *gin.RouterGroup, authMiddleware gin.HandlerFunc) {
	ledger := r.Group("/ledger")
	ledger.Use(authMiddleware)
	{
		ledger.GET("/transactions", h.ListTransactions)
		ledger.POST("/transactions", h.CreateTransaction)
		ledger.GET("/budgets", h.ListBudgets)
		ledger.POST("/budgets", h.CreateBudget)
		ledger.GET("/reports", h.ListReports)
		ledger.POST("/reports", h.CreateReport)
		ledger.POST("/import", h.ImportTransactions)
		ledger.POST("/sheets/import", h.ImportTransactionsSheet)
		ledger.GET("/export", h.ExportTransactions)
	}
}

// ListTransactions godoc
// @Summary Получить список транзакций
// @Description Возвращает транзакции пользователя из Ledger.
// @Tags ledger
// @Produce json
// @Security BearerAuth
// @Success 200 {object} model.TransactionsResponse
// @Failure 401 {object} model.ErrorResponse
// @Failure 500 {object} model.ErrorResponse
// @Router /api/ledger/transactions [get]
func (h *LedgerHandler) ListTransactions(c *gin.Context) {
	accountID := middleware.UserIDFromContext(c)
	if accountID == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user not found in context"})
		return
	}

	items, err := h.service.ListTransactions(c.Request.Context(), accountID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"transactions": items})
}

// CreateTransaction godoc
// @Summary Создать транзакцию
// @Description Создает транзакцию. Если account_id не указан, берется из JWT.
// @Tags ledger
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body model.CreateTransactionRequest true "Данные транзакции"
// @Success 201 {object} model.Transaction
// @Failure 400 {object} model.ErrorResponse
// @Failure 401 {object} model.ErrorResponse
// @Failure 500 {object} model.ErrorResponse
// @Router /api/ledger/transactions [post]
func (h *LedgerHandler) CreateTransaction(c *gin.Context) {
	var req model.CreateTransactionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if req.AccountID == "" {
		req.AccountID = middleware.UserIDFromContext(c)
	}
	if req.AccountID == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "account_id is required"})
		return
	}

	created, err := h.service.CreateTransaction(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, created)
}

// ListBudgets godoc
// @Summary Получить список бюджетов
// @Description Возвращает бюджеты из Ledger.
// @Tags ledger
// @Produce json
// @Security BearerAuth
// @Success 200 {object} model.BudgetsResponse
// @Failure 401 {object} model.ErrorResponse
// @Failure 500 {object} model.ErrorResponse
// @Router /api/ledger/budgets [get]
func (h *LedgerHandler) ListBudgets(c *gin.Context) {
	items, err := h.service.ListBudgets(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"budgets": items})
}

// CreateBudget godoc
// @Summary Создать бюджет
// @Description Создает бюджет в Ledger.
// @Tags ledger
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body model.CreateBudgetRequest true "Данные бюджета"
// @Success 201 {object} model.Budget
// @Failure 400 {object} model.ErrorResponse
// @Failure 401 {object} model.ErrorResponse
// @Failure 500 {object} model.ErrorResponse
// @Router /api/ledger/budgets [post]
func (h *LedgerHandler) CreateBudget(c *gin.Context) {
	var req model.CreateBudgetRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	created, err := h.service.CreateBudget(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, created)
}

// ListReports godoc
// @Summary Получить список отчетов
// @Description Возвращает отчеты из Ledger.
// @Tags ledger
// @Produce json
// @Security BearerAuth
// @Success 200 {object} model.ReportsResponse
// @Failure 401 {object} model.ErrorResponse
// @Failure 500 {object} model.ErrorResponse
// @Router /api/ledger/reports [get]
func (h *LedgerHandler) ListReports(c *gin.Context) {
	items, err := h.service.ListReports(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"reports": items})
}

// CreateReport godoc
// @Summary Создать отчет
// @Description Создает отчет в Ledger.
// @Tags ledger
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body model.CreateReportRequest true "Данные отчета"
// @Success 201 {object} model.Report
// @Failure 400 {object} model.ErrorResponse
// @Failure 401 {object} model.ErrorResponse
// @Failure 500 {object} model.ErrorResponse
// @Router /api/ledger/reports [post]
func (h *LedgerHandler) CreateReport(c *gin.Context) {
	var req model.CreateReportRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	created, err := h.service.CreateReport(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, created)
}

// ImportTransactions godoc
// @Summary Импортировать транзакции из CSV
// @Description Принимает CSV контент и импортирует транзакции.
// @Tags ledger
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body model.ImportTransactionsRequest true "CSV данные"
// @Success 200 {object} model.ImportTransactionsResponse
// @Failure 400 {object} model.ErrorResponse
// @Failure 401 {object} model.ErrorResponse
// @Failure 500 {object} model.ErrorResponse
// @Router /api/ledger/import [post]
func (h *LedgerHandler) ImportTransactions(c *gin.Context) {
	var req model.ImportTransactionsRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	imported, err := h.service.ImportTransactionsCSV(c.Request.Context(), []byte(req.CSVContent), req.HasHeader)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, model.ImportTransactionsResponse{Imported: imported})
}

// ExportTransactions godoc
// @Summary Экспортировать транзакции в CSV
// @Description Возвращает CSV контент транзакций пользователя.
// @Tags ledger
// @Produce json
// @Security BearerAuth
// @Success 200 {object} model.ExportTransactionsResponse
// @Failure 401 {object} model.ErrorResponse
// @Failure 500 {object} model.ErrorResponse
// @Router /api/ledger/export [get]
func (h *LedgerHandler) ExportTransactions(c *gin.Context) {
	accountID := middleware.UserIDFromContext(c)
	if accountID == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user not found in context"})
		return
	}

	csvContent, err := h.service.ExportTransactionsCSV(c.Request.Context(), accountID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, model.ExportTransactionsResponse{CSVContent: string(csvContent)})
}

// ImportTransactionsSheet godoc
// @Summary Импортировать транзакции из Google Sheets
// @Description Принимает данные транзакций в формате Sheets и проксирует в Ledger.
// @Tags ledger
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body model.ImportTransactionsSheetRequest true "Данные транзакций из Sheets"
// @Success 200 {object} model.ImportTransactionsSheetResponse
// @Failure 400 {object} model.ErrorResponse
// @Failure 401 {object} model.ErrorResponse
// @Failure 500 {object} model.ErrorResponse
// @Router /api/ledger/sheets/import [post]
func (h *LedgerHandler) ImportTransactionsSheet(c *gin.Context) {
	var req model.ImportTransactionsSheetRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	imported, err := h.service.ImportTransactionsSheet(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, model.ImportTransactionsSheetResponse{Imported: imported})
}
