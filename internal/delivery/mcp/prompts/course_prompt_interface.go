package prompts

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"
)

// GetPromptHandler 處理 Prompt 請求的函式簽名。
type GetPromptHandler func(ctx context.Context, request mcp.GetPromptRequest) (*mcp.GetPromptResult, error)

// CoursePromptService 定義課程相關的 Prompt 服務介面。
//
// 此介面負責提供與課程相關的 MCP Prompts 及其處理邏輯。
//
// Example:
//
//	var service CoursePromptService = NewCoursePromptServiceImpl(...)
//	greetingPrompt := service.GetGreetingPrompt()
type CoursePromptService interface {
	// GetGreetingPrompt 取得問候 prompt 的描述。
	GetGreetingPrompt() mcp.Prompt
	// GetGreetingHandler 處理問候 prompt 請求。
	GetGreetingHandler(ctx context.Context, request mcp.GetPromptRequest) (*mcp.GetPromptResult, error)

	// GetOperationSuggestionPrompt 取得操作建議 prompt 的描述。
	GetOperationSuggestionPrompt() mcp.Prompt
	// GetOperationSuggestionHandler 處理操作建議 prompt 請求。
	GetOperationSuggestionHandler(ctx context.Context, request mcp.GetPromptRequest) (*mcp.GetPromptResult, error)
}
