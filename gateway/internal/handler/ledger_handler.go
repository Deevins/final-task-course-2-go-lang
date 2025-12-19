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
		ledger.GET("/export", h.ExportTransactions)
	}
}

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

func (h *LedgerHandler) ListBudgets(c *gin.Context) {
	items, err := h.service.ListBudgets(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"budgets": items})
}

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

func (h *LedgerHandler) ListReports(c *gin.Context) {
	items, err := h.service.ListReports(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"reports": items})
}

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
