package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
	GET("/admin-area", AuthorizationMiddleware("admin", "manager"), func(c *gin.Context) {
	   c.JSON(http.StatusOK, gin.H{"message": "Welcome to admin area"})
	})
*/
func AuthorizationMiddleware(allowedRoles ...string) gin.HandlerFunc {
	return func(context *gin.Context) {
		roleAny, exists := context.Get("permissions")

		if !exists {
			context.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized access: roles missing"})
			context.Abort()
			return
		}

		if roleAny == nil {
			context.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized access: roles are nil"})
			context.Abort()
			return
		}

		var roles []string

		switch v := roleAny.(type) {
		case string:
			roles = []string{v}
		case []string:
			roles = v
		case []interface{}:
			for _, r := range v {
				if strRole, ok := r.(string); ok {
					roles = append(roles, strRole)
				}
			}
		default:
			context.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized access: invalid role format"})
			context.Abort()
			return
		}

		for _, allowed := range allowedRoles {
			for _, userRole := range roles {
				if userRole == allowed {
					context.Next()
					return
				}
			}
		}

		context.JSON(http.StatusForbidden, gin.H{"error": "Forbidden - insufficient permissions"})
		context.Abort()
	}
}
