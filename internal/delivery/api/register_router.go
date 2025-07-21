package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"

	"course-mcp/internal/delivery/api/handlers"
	"course-mcp/internal/delivery/api/middlewares"
	"course-mcp/internal/usecase/utils"
)

func NewRouter(
	logger *zerolog.Logger,
	authProvider utils.AuthProvider,
	mcpServer http.Handler,
	authMiddleware *middlewares.AuthMiddleware,
) *gin.Engine {

	router := gin.Default()

	// global middleware
	router.Use(middlewares.CorsMiddleware())

	// register POST, GET, DELETE methods for the /mcp path, all handled by MCPServer
	for _, method := range []string{http.MethodPost, http.MethodGet, http.MethodDelete} {
		router.Handle(method, "/mcp",
			authMiddleware.Authentication(),
			// authMiddleware.Authorization(),
			gin.WrapH(mcpServer))
	}

	// register Auth routes
	registerAuthRoutes(logger, authProvider, router)

	return router
}

func registerAuthRoutes(logger *zerolog.Logger, authProvider utils.AuthProvider, router *gin.Engine) {
	// API handlers
	authHandler := handlers.NewAuthHandler(logger, authProvider)
	oauthMetadataHandler := handlers.NewOAuthMetadataHandler(logger, authProvider)
	tokenHandler := handlers.NewTokenHandler(logger, authProvider)

	// OAuth Protected Resources
	router.GET("/.well-known/oauth-protected-resource", oauthMetadataHandler.HandleOAuthProtectedResource)

	// OAuth Authorization Server
	router.GET("/.well-known/oauth-authorization-server", middlewares.CorsMiddleware(), oauthMetadataHandler.HandleOAuthAuthorizationServer)

	// Register Client
	router.POST("/register", middlewares.CorsMiddleware(), authHandler.HandleRegister)

	// Token endpoint
	router.POST("/token", middlewares.CorsMiddleware("Authorization", "Content-Type"), tokenHandler.HandleToken)

	// Authorize endpoint
	router.GET("/authorize", middlewares.CorsMiddleware("Authorization", "Content-Type"), tokenHandler.HandleAuthorize)
}
