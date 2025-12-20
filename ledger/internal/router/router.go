package router

import (
	"github.com/gin-gonic/gin"

	httpHandler "github.com/Deevins/final-task-course-2-go-lang/ledger/internal/handler/http"
)

func Register(engine *gin.Engine, healthHandler *httpHandler.HealthHandler) {
	api := engine.Group("/api/v1")
	healthHandler.RegisterRoutes(api)
}
