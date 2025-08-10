package middleware

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
)

func Middleware(middlewares []string, authMiddleWare *JWTMiddleware) []gin.HandlerFunc {
	var handlers []gin.HandlerFunc

	for _, m := range middlewares {
		parts := strings.SplitN(m, ":", 2)
		name := parts[0]
		param := ""
		if len(parts) == 2 {
			param = parts[1]
		}

		switch name {
		case "auth":
			handlers = append(handlers, authMiddleWare.Middleware())
		case "permissions":
			if param == "" {
				fmt.Println("authorization middleware requires a role param, skipping")
				continue
			}
			roles := strings.Split(param, "|")
			handlers = append(handlers, AuthorizationMiddleware(roles...))
		case "rate-limited":
			handlers = append(handlers, RateLimitMiddleware())
		case "cors":
			handlers = append(handlers, CORSMiddleware())
		case "recovery":
			handlers = append(handlers, gin.Recovery())
		default:
			fmt.Printf("Unknown middleware: %s, skipping\n", name)
		}
	}

	return handlers
}
