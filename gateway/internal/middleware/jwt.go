package middleware

import (
	"net/http"
	"strings"

	"github.com/Deevins/final-task-course-2-go-lang/gateway/internal/service"
	"github.com/gin-gonic/gin"
)

const userIDContextKey = "user_id"

func JWTAuth(authService service.AuthGatewayService) gin.HandlerFunc {
	if authService == nil {
		panic("JWT middleware requires auth service")
	}

	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "authorization header is required"})
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || !strings.EqualFold(parts[0], "Bearer") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid authorization header"})
			return
		}

		token := strings.TrimSpace(parts[1])
		if token == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "token is required"})
			return
		}

		resp, err := authService.ValidateToken(c.Request.Context(), token)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "token validation failed"})
			return
		}
		if !resp.GetValid() {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "token is invalid"})
			return
		}

		c.Set(userIDContextKey, resp.GetUserId())
		c.Next()
	}
}

func UserIDFromContext(c *gin.Context) string {
	if c == nil {
		return ""
	}
	value, ok := c.Get(userIDContextKey)
	if !ok {
		return ""
	}
	userID, _ := value.(string)
	return userID
}
