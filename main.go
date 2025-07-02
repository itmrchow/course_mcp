package main

import (
	"log"
	"os"

	"github.com/mark3labs/mcp-go/server"
	"github.com/rs/zerolog"

	"course-mcp/internal/delivery/mcp/prompts"
	"course-mcp/internal/delivery/mcp/resources"
	"course-mcp/internal/delivery/mcp/tools"
)

func main() {
	// 設定日誌格式
	logger := zerolog.New(os.Stdout).With().Timestamp().Logger()

	s := server.NewMCPServer(
		"course-mcp-server", // Server name
		"1.0.0",             // Server version
		server.WithResourceCapabilities(true, true),
		server.WithToolCapabilities(true),
		server.WithRecovery(),
	)

	// tools
	courseTools := tools.NewCourseMCPToolImpl(&logger)

	// resources 給LLM看的資訊
	courseResource := resources.NewCourseMCPResourceImpl(&logger)

	// prompt
	coursePromptService := prompts.NewCoursePromptServiceImpl()

	// register
	resources.RegisterCourseResources(s, courseResource)
	tools.RegisterCourseTools(s, courseTools)
	prompts.RegisterCoursePrompts(s, coursePromptService)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	logger.Info().Msg("Server started , port: " + port)
	sseServer := server.NewSSEServer(s)
	if err := sseServer.Start(":" + port); err != nil {
		log.Fatal("Failed to start server:", err)
	}

	// err := server.ServeStdio(s)
	// if err != nil {
	// 	log.Fatal("Failed to start server:", err)
	// }
}
