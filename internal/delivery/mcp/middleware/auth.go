package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(c *gin.Context) {
	if c.GetHeader("Authorization") == "" {

		// add header to response
		c.Header("WWW-Authenticate", `Bearer resource_metadata="localhost:3000/.well-known/oauth-authorization-server"`)
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	// TODO: 驗證token有效性
	c.Next()
}
