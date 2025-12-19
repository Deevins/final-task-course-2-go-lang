package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/Deevins/final-task-course-2-go-lang/gateway/internal/model"
	"github.com/Deevins/final-task-course-2-go-lang/gateway/internal/service"
)

type BudgetHandler struct {
	service service.BudgetGatewayService
}

func NewBudgetHandler(s service.BudgetGatewayService) *BudgetHandler {
	if s == nil {
		panic("BudgetHandler requires service")
	}
	return &BudgetHandler{service: s}
}

func (h *BudgetHandler) Register(r *gin.RouterGroup) {
	budget := r.Group("/budget")
	{
		budget.POST("/export", h.ExportBudget)
		budget.GET("/import", h.ImportBudget)
		budget.GET("/download", h.DownloadBudget)
	}
}

// ExportBudget godoc
// @Summary Экспортировать бюджет в Google Sheets
// @Description Записывает строки бюджета в Google Sheets.
// @Tags budget
// @Accept json
// @Produce json
// @Param request body model.ExportSimpleRequest true "Данные для экспорта"
// @Success 200 {object} model.ExportSimpleResponse
// @Failure 400 {object} model.ErrorResponse
// @Failure 500 {object} model.ErrorResponse
// @Router /api/budget/export [post]
func (h *BudgetHandler) ExportBudget(c *gin.Context) {
	var req model.ExportSimpleRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if len(req.Rows) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "rows must not be empty"})
		return
	}
	resp, err := h.service.ExportBudget(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

// ImportBudget godoc
// @Summary Импортировать бюджет из Google Sheets
// @Description Возвращает строки бюджета из Google Sheets по идентификатору таблицы.
// @Tags budget
// @Produce json
// @Param spreadsheet_id query string true "ID таблицы Google Sheets"
// @Param sheet_name query string false "Имя листа" default(Report)
// @Success 200 {object} model.BudgetRowsResponse
// @Failure 400 {object} model.ErrorResponse
// @Failure 500 {object} model.ErrorResponse
// @Router /api/budget/import [get]
func (h *BudgetHandler) ImportBudget(c *gin.Context) {
	spreadsheetID := c.Query("spreadsheet_id")
	if spreadsheetID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "spreadsheet_id is required"})
		return
	}
	sheetName := c.DefaultQuery("sheet_name", "Report")
	rows, err := h.service.ImportBudget(c.Request.Context(), spreadsheetID, sheetName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"rows": rows})
}

// DownloadBudget godoc
// @Summary Скачать бюджет по умолчанию
// @Description Возвращает строки бюджета по умолчанию.
// @Tags budget
// @Produce json
// @Success 200 {object} model.BudgetRowsResponse
// @Failure 500 {object} model.ErrorResponse
// @Router /api/budget/download [get]
func (h *BudgetHandler) DownloadBudget(c *gin.Context) {
	rows, err := h.service.DownloadBudget(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"rows": rows})
}
