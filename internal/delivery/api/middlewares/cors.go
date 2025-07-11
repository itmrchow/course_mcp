package middlewares

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CorsMiddleware(allowedHeaders ...string) gin.HandlerFunc {
	// CORS middleware for handling preflight and actual requests

	fmt.Println("CORS middleware initialized with allowed headers:", allowedHeaders)
	headers := "Mcp-Protocol-Version, Authorization, Content-Type"
	if len(allowedHeaders) > 0 {
		// 添加基本的必需 headers 以及用戶指定的 headers
		baseHeaders := "Mcp-Protocol-Version, Authorization, Content-Type"
		additionalHeaders := ""
		for i, h := range allowedHeaders {
			if i > 0 {
				additionalHeaders += ", "
			}
			additionalHeaders += h
		}
		headers = baseHeaders + ", " + additionalHeaders
	}
	return func(c *gin.Context) {
		origin := c.Request.Header.Get("Origin")
		if origin != "" {
			c.Header("Access-Control-Allow-Origin", origin)
		} else {
			c.Header("Access-Control-Allow-Origin", "*")
		}
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", headers)
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Max-Age", "86400")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusOK)
			return
		}
		c.Next()
	}
}
