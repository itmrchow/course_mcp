package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/mark3labs/mcp-go/server"
	"github.com/rs/zerolog"
	"github.com/spf13/viper"

	"course-mcp/internal/delivery/api"
	"course-mcp/internal/delivery/mcp/prompts"
	"course-mcp/internal/delivery/mcp/resources"
	"course-mcp/internal/delivery/mcp/tools"
	"course-mcp/internal/delivery/mcp/utils"
	infra "course-mcp/internal/infrastructure"
)

func main() {

	infra.InitConfig()

	// 設定日誌格式
	logger := zerolog.New(os.Stdout).With().Timestamp().Logger()

	s := server.NewMCPServer(
		"course-mcp-server", // Server name
		"1.0.0",             // Server version
		server.WithResourceCapabilities(true, true),
		server.WithToolCapabilities(true),
		server.WithLogging(),
		server.WithRecovery(),
	)

	// mcpServer
	mcpServer := server.NewStreamableHTTPServer(
		s,
		server.WithHeartbeatInterval(30*time.Second),
		server.WithHTTPContextFunc(func(
			ctx context.Context,
			r *http.Request,
		) context.Context {
			authVal := r.Header.Get("Authorization")
			ctx = context.WithValue(ctx, "auth", authVal)

			// TODO: trace id
			return ctx
		}),
	)

	// tools
	courseTools := tools.NewCourseMCPToolImpl(&logger)
	teacherTools := tools.NewTeacherMCPToolImpl(&logger)

	// resources 給LLM看的資訊
	courseResource := resources.NewCourseMCPResourceImpl(&logger)

	// prompt
	coursePromptService := prompts.NewCoursePromptServiceImpl()

	// register
	resources.RegisterCourseResources(s, courseResource)
	tools.RegisterCourseTools(s, courseTools)
	tools.RegisterTeacherTools(s, teacherTools)
	prompts.RegisterCoursePrompts(s, coursePromptService)

	// auth provider
	authProvider := utils.NewKeycloakAuthProvider()

	// gin
	router := api.InitializeRouter(&logger, authProvider, mcpServer)

	port := viper.GetString("PORT")
	if port == "" {
		port = "3000"
	}
	logger.Info().Msg("Server started , port: " + port)

	srv := &http.Server{
		Addr:         fmt.Sprintf("localhost:%s", port),
		Handler:      router,
		ReadTimeout:  10 * time.Second, // 10 seconds
		WriteTimeout: 10 * time.Second, // 10 seconds
		IdleTimeout:  60 * time.Second, // 60 seconds
	}

	if err := srv.ListenAndServe(); err != nil {
		logger.Error().Err(err).Msg("server error")
		os.Exit(1)
	}

	// logger.Info().Msg("Server started , port: " + port)

	// if err := httpServer.Start(":" + port); err != nil {
	// 	log.Fatal("Failed to start server:", err)
	// }
}

type AuthMiddleware struct {
	jwtSecret []byte
	// userStore UserStore
}

func NewAuthMiddleware(secret []byte) *AuthMiddleware {
	return &AuthMiddleware{
		jwtSecret: secret,
		// userStore: store,
	}
}

func (m *AuthMiddleware) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Extract token from Authorization header
		authHeader := r.Header.Get("Authorization")
		if !strings.HasPrefix(authHeader, "Bearer ") {
			http.Error(w, "Missing or invalid authorization header", http.StatusUnauthorized)
			return
		}

		token := strings.TrimPrefix(authHeader, "Bearer ")

		// Validate JWT token
		claims, err := m.validateJWT(token)
		if err != nil {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		// Load user information
		// user, err := m.userStore.GetUser(claims.UserID)
		// if err != nil {
		// 	http.Error(w, "User not found", http.StatusUnauthorized)
		// 	return
		// }

		user := &Claims{
			UserID: claims.UserID,
			Role:   claims.Role,
		}

		// Add user to request context
		ctx := context.WithValue(r.Context(), "user", user)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (m *AuthMiddleware) validateJWT(tokenString string) (*Claims, error) {
	// Note: This example uses a hypothetical JWT library
	// In practice, you would use a real JWT library like github.com/golang-jwt/jwt
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return m.jwtSecret, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("invalid token")
}

type Claims struct {
	UserID string `json:"user_id"`
	Role   string `json:"role"`
	jwt.StandardClaims
}
