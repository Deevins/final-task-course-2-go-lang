package handler

import (
	"net/http"

	"github.com/Deevins/final-task-course-2-go-lang/gateway/internal/model"
	"github.com/Deevins/final-task-course-2-go-lang/gateway/internal/service"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	service service.AuthGatewayService
}

func NewAuthHandler(s service.AuthGatewayService) *AuthHandler {
	if s == nil {
		panic("AuthHandler requires service")
	}
	return &AuthHandler{service: s}
}

func (h *AuthHandler) Register(r *gin.RouterGroup) {
	auth := r.Group("/auth")
	{
		auth.POST("/signup", h.SignUp)
		auth.POST("/signin", h.SignIn)
	}
}

func (h *AuthHandler) SignUp(c *gin.Context) {
	var req model.SignUpRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := h.service.SignUp(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, resp)
}

func (h *AuthHandler) SignIn(c *gin.Context) {
	var req model.SignInRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := h.service.SignIn(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}
