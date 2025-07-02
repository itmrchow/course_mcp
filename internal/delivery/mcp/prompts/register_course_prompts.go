package prompts

import (
	"github.com/mark3labs/mcp-go/server"
)

// RegisterCoursePrompts 註冊所有課程相關的 Prompt 到 MCP 伺服器。
//
// Arguments:
//
//	s: MCP 伺服器實例。
//	promptService: 實現 CoursePromptService 介面的實例。
//
// Example:
//
//	svc := NewCoursePromptServiceImpl()
//	RegisterCoursePrompts(mcpServer, svc)
func RegisterCoursePrompts(s *server.MCPServer, promptService CoursePromptService) {
	s.AddPrompt(promptService.GetGreetingPrompt(), promptService.GetGreetingHandler)
	s.AddPrompt(promptService.GetOperationSuggestionPrompt(), promptService.GetOperationSuggestionHandler)
}
