package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/mark3labs/mcp-go/server"
	"github.com/rs/zerolog"
	"github.com/spf13/viper"

	"course-mcp/internal/delivery/api"
	"course-mcp/internal/delivery/api/middlewares"
	"course-mcp/internal/delivery/mcp/prompts"
	"course-mcp/internal/delivery/mcp/resources"
	mcpServer "course-mcp/internal/delivery/mcp/server"
	"course-mcp/internal/delivery/mcp/tools"
	infra "course-mcp/internal/infrastructure"
	"course-mcp/internal/usecase/utils"
)

func main() {

	// config + env
	infra.InitConfig()

	// 設定日誌格式
	logger := zerolog.New(os.Stdout).With().Timestamp().Logger()

	// utils
	// - scope 管理器
	scopeManager := tools.NewScopeManager()
	// - token validator
	tokenValidator := utils.NewTokenValidator()
	// - auth provider -> 呼叫 auth server
	authProvider := utils.NewKeycloakAuthProvider()

	// mcp
	// - 建立 MCP 伺服器管理器 , 包含MCP Server所需的Handler funcs
	mcpServerManager := mcpServer.NewMCPServerManager(&logger, scopeManager)

	// - 建立 MCP 伺服器實例
	s := mcpServer.NewMCPServer(&logger, mcpServerManager)

	// - mcp http server
	mcpServer := server.NewStreamableHTTPServer(
		s,
		server.WithHeartbeatInterval(30*time.Second),
	)

	// - mcp tools
	courseTools := tools.NewCourseMCPToolImpl(&logger)
	teacherTools := tools.NewTeacherMCPToolImpl(&logger)

	// - resources 給LLM看的資訊
	courseResource := resources.NewCourseMCPResourceImpl(&logger)

	// - prompt
	coursePromptService := prompts.NewCoursePromptServiceImpl()

	// - register
	resources.RegisterCourseResources(s, courseResource)
	tools.RegisterCourseTools(s, courseTools)
	tools.RegisterTeacherTools(s, teacherTools)
	prompts.RegisterCoursePrompts(s, coursePromptService)

	// gin (http server)

	// - API middleware
	authMiddleware := middlewares.NewAuthMiddleware(tokenValidator)

	// - router
	router := api.NewRouter(&logger, authProvider, mcpServer, authMiddleware)

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
