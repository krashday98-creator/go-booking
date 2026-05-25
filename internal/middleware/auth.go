package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const UserIDHeader = "X-User-ID"
const AdminKeyHeader = "X-Admin-Key"

func AdminAuth(adminAPIKey string) gin.HandlerFunc {
	return func(c *gin.Context) {
		key := c.GetHeader(AdminKeyHeader)
		if key == "" || key != adminAPIKey {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid or missing admin API key"})
			return
		}
		c.Next()
	}
}

func UserAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.GetHeader(UserIDHeader)
		if userID == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "X-User-ID header is required"})
			return
		}
		c.Set("user_id", userID)
		c.Next()
	}
}
