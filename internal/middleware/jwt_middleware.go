package middleware

import (
	"alpha-core/internal/config"
	"alpha-core/pkg/logger"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

type JWTMiddleware struct {
	configure *config.Config
	log       *logger.Logger
}

func NewJWTMiddleware(configure *config.Config, log *logger.Logger) *JWTMiddleware {
	return &JWTMiddleware{
		configure: configure,
		log:       log,
	}
}

func (middleware *JWTMiddleware) Middleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		authorization := context.GetHeader("Authorization")

		if authorization == "" {
			middleware.log.ErrorLog("Authorization header is missing")
			context.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
			context.Abort()
			return
		}

		parts := strings.SplitN(authorization, " ", 2)

		if len(parts) != 2 || parts[0] != "Bearer" {
			middleware.log.ErrorLog("Authorization header format is invalid")
			context.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header format is invalid"})
			context.Abort()
			return
		}

		tokenString := parts[1]
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				middleware.log.ErrorLog("Unexpected signing method: %v", token.Header["alg"])
				return nil, jwt.ErrSignatureInvalid
			}
			return []byte(middleware.configure.JWTSecret), nil
		})

		if err != nil {
			middleware.log.ErrorLog("Failed to parse token: %v", err)
			context.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			context.Abort()
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			// push useful claims to context
			if sub, ok := claims["sub"]; ok {
				context.Set("sub", sub)
			}
			if email, ok := claims["email"]; ok {
				context.Set("email", email)
			}
			if roles, ok := claims["roles"]; ok {
				context.Set("roles", roles)
			}
			if perms, ok := claims["perms"]; ok {
				context.Set("perms", perms)
			}
		}

		context.Next()
	}
}
