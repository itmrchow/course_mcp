package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"course-mcp/internal/usecase/utils"
)

type AuthMiddleware struct {
	tokenValidator *utils.TokenValidator
}

func NewAuthMiddleware(tokenValidator *utils.TokenValidator) *AuthMiddleware {
	return &AuthMiddleware{
		tokenValidator: tokenValidator,
	}
}

func (a *AuthMiddleware) RequireAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString := ctx.GetHeader("Authorization")
		if tokenString == "" {
			// add header to response
			ctx.Header("WWW-Authenticate", `Bearer resource_metadata="localhost:3000/.well-known/oauth-authorization-server"`) // TODO: value from env
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		tokenClaims, err := a.tokenValidator.Validate(tokenString)
		if err != nil {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		// Set the token claims in the context for further use
		ctx.Set("tokenClaims", tokenClaims)
		ctx.Next()
	}
}
