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
		transactions := ledger.Group("/transactions")
		{
			transactions.GET("", h.ListTransactions)
			transactions.POST("", h.CreateTransaction)
			transactions.GET("/:id", h.GetTransaction)
			transactions.PUT("/:id", h.UpdateTransaction)
			transactions.PATCH("/:id", h.UpdateTransaction)
			transactions.DELETE("/:id", h.DeleteTransaction)
		}
		budgets := ledger.Group("/budgets")
		{
			budgets.GET("", h.ListBudgets)
			budgets.POST("", h.CreateBudget)
			budgets.GET("/:id", h.GetBudget)
			budgets.PUT("/:id", h.UpdateBudget)
			budgets.PATCH("/:id", h.UpdateBudget)
			budgets.DELETE("/:id", h.DeleteBudget)
		}
		reports := ledger.Group("/reports")
		{
			reports.GET("", h.ListReports)
			reports.POST("", h.CreateReport)
			reports.GET("/:id", h.GetReport)
			reports.PUT("/:id", h.UpdateReport)
			reports.PATCH("/:id", h.UpdateReport)
			reports.DELETE("/:id", h.DeleteReport)
		}
		ledger.POST("/import", h.ImportTransactions)
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

// GetTransaction godoc
// @Summary Получить транзакцию
// @Description Возвращает транзакцию по идентификатору.
// @Tags ledger
// @Produce json
// @Security BearerAuth
// @Param id path string true "ID транзакции"
// @Success 200 {object} model.Transaction
// @Failure 400 {object} model.ErrorResponse
// @Failure 401 {object} model.ErrorResponse
// @Failure 500 {object} model.ErrorResponse
// @Router /api/ledger/transactions/{id} [get]
func (h *LedgerHandler) GetTransaction(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id is required"})
		return
	}

	item, err := h.service.GetTransaction(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, item)
}

// UpdateTransaction godoc
// @Summary Обновить транзакцию
// @Description Обновляет транзакцию по идентификатору.
// @Tags ledger
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "ID транзакции"
// @Param request body model.UpdateTransactionRequest true "Данные транзакции"
// @Success 200 {object} model.Transaction
// @Failure 400 {object} model.ErrorResponse
// @Failure 401 {object} model.ErrorResponse
// @Failure 500 {object} model.ErrorResponse
// @Router /api/ledger/transactions/{id} [put]
// @Router /api/ledger/transactions/{id} [patch]
func (h *LedgerHandler) UpdateTransaction(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id is required"})
		return
	}

	var req model.UpdateTransactionRequest
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

	updated, err := h.service.UpdateTransaction(c.Request.Context(), id, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, updated)
}

// DeleteTransaction godoc
// @Summary Удалить транзакцию
// @Description Удаляет транзакцию по идентификатору.
// @Tags ledger
// @Produce json
// @Security BearerAuth
// @Param id path string true "ID транзакции"
// @Success 200 {object} model.DeleteResponse
// @Failure 400 {object} model.ErrorResponse
// @Failure 401 {object} model.ErrorResponse
// @Failure 500 {object} model.ErrorResponse
// @Router /api/ledger/transactions/{id} [delete]
func (h *LedgerHandler) DeleteTransaction(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id is required"})
		return
	}

	deleted, err := h.service.DeleteTransaction(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, model.DeleteResponse{Deleted: deleted})
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

// GetBudget godoc
// @Summary Получить бюджет
// @Description Возвращает бюджет по идентификатору.
// @Tags ledger
// @Produce json
// @Security BearerAuth
// @Param id path string true "ID бюджета"
// @Success 200 {object} model.Budget
// @Failure 400 {object} model.ErrorResponse
// @Failure 401 {object} model.ErrorResponse
// @Failure 500 {object} model.ErrorResponse
// @Router /api/ledger/budgets/{id} [get]
func (h *LedgerHandler) GetBudget(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id is required"})
		return
	}

	item, err := h.service.GetBudget(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, item)
}

// UpdateBudget godoc
// @Summary Обновить бюджет
// @Description Обновляет бюджет по идентификатору.
// @Tags ledger
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "ID бюджета"
// @Param request body model.UpdateBudgetRequest true "Данные бюджета"
// @Success 200 {object} model.Budget
// @Failure 400 {object} model.ErrorResponse
// @Failure 401 {object} model.ErrorResponse
// @Failure 500 {object} model.ErrorResponse
// @Router /api/ledger/budgets/{id} [put]
// @Router /api/ledger/budgets/{id} [patch]
func (h *LedgerHandler) UpdateBudget(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id is required"})
		return
	}

	var req model.UpdateBudgetRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updated, err := h.service.UpdateBudget(c.Request.Context(), id, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, updated)
}

// DeleteBudget godoc
// @Summary Удалить бюджет
// @Description Удаляет бюджет по идентификатору.
// @Tags ledger
// @Produce json
// @Security BearerAuth
// @Param id path string true "ID бюджета"
// @Success 200 {object} model.DeleteResponse
// @Failure 400 {object} model.ErrorResponse
// @Failure 401 {object} model.ErrorResponse
// @Failure 500 {object} model.ErrorResponse
// @Router /api/ledger/budgets/{id} [delete]
func (h *LedgerHandler) DeleteBudget(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id is required"})
		return
	}

	deleted, err := h.service.DeleteBudget(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, model.DeleteResponse{Deleted: deleted})
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

// GetReport godoc
// @Summary Получить отчет
// @Description Возвращает отчет по идентификатору.
// @Tags ledger
// @Produce json
// @Security BearerAuth
// @Param id path string true "ID отчета"
// @Success 200 {object} model.Report
// @Failure 400 {object} model.ErrorResponse
// @Failure 401 {object} model.ErrorResponse
// @Failure 500 {object} model.ErrorResponse
// @Router /api/ledger/reports/{id} [get]
func (h *LedgerHandler) GetReport(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id is required"})
		return
	}

	item, err := h.service.GetReport(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, item)
}

// UpdateReport godoc
// @Summary Обновить отчет
// @Description Обновляет отчет по идентификатору.
// @Tags ledger
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "ID отчета"
// @Param request body model.UpdateReportRequest true "Данные отчета"
// @Success 200 {object} model.Report
// @Failure 400 {object} model.ErrorResponse
// @Failure 401 {object} model.ErrorResponse
// @Failure 500 {object} model.ErrorResponse
// @Router /api/ledger/reports/{id} [put]
// @Router /api/ledger/reports/{id} [patch]
func (h *LedgerHandler) UpdateReport(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id is required"})
		return
	}

	var req model.UpdateReportRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updated, err := h.service.UpdateReport(c.Request.Context(), id, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, updated)
}

// DeleteReport godoc
// @Summary Удалить отчет
// @Description Удаляет отчет по идентификатору.
// @Tags ledger
// @Produce json
// @Security BearerAuth
// @Param id path string true "ID отчета"
// @Success 200 {object} model.DeleteResponse
// @Failure 400 {object} model.ErrorResponse
// @Failure 401 {object} model.ErrorResponse
// @Failure 500 {object} model.ErrorResponse
// @Router /api/ledger/reports/{id} [delete]
func (h *LedgerHandler) DeleteReport(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id is required"})
		return
	}

	deleted, err := h.service.DeleteReport(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, model.DeleteResponse{Deleted: deleted})
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
