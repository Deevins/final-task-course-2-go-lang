package http

import (
	"errors"
	"net/http"

	"github.com/Deevins/final-task-course-2-go-lang/ledger/internal/service"
	"github.com/Deevins/final-task-course-2-go-lang/ledger/internal/storage"
	"github.com/gin-gonic/gin"
)

type ReportHandler struct {
	service service.LedgerService
}

func NewReportHandler(service service.LedgerService) *ReportHandler {
	if service == nil {
		panic("ReportHandler requires service")
	}
	return &ReportHandler{service: service}
}

func (h *ReportHandler) RegisterRoutes(router *gin.RouterGroup) {
	router.GET("/reports/export", h.ExportReportSheet)
}

func (h *ReportHandler) ExportReportSheet(c *gin.Context) {
	reportID := c.Query("report_id")
	if reportID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "report_id is required"})
		return
	}

	report, err := h.service.ExportReportSheet(reportID)
	if err != nil {
		if errors.Is(err, storage.ErrNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "report not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, report)
}
