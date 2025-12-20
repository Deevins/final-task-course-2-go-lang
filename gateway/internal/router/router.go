package router

import (
	"github.com/gin-gonic/gin"

	"github.com/Deevins/final-task-course-2-go-lang/gateway/internal/handler"
)

func Register(engine *gin.Engine, authHandler *handler.AuthHandler, ledgerHandler *handler.LedgerHandler, authMiddleware gin.HandlerFunc) {
	registerSwagger(engine)

	api := engine.Group("/api")
	authHandler.Register(api)
	ledgerHandler.Register(api, authMiddleware)
}
