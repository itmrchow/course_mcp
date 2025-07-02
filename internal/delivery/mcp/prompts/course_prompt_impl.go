package prompts

import (
	"context"
	"fmt"
	"log"

	"github.com/mark3labs/mcp-go/mcp"
)

// CoursePromptServiceImpl 是 CoursePromptService 介面的實作。
type CoursePromptServiceImpl struct {
	// 這裡可以加入你需要的依賴，例如 usecase 或 repository
	// usecase CourseUsecase
}

// NewCoursePromptServiceImpl 建立一個新的 CoursePromptServiceImpl 實例。
func NewCoursePromptServiceImpl() *CoursePromptServiceImpl {
	return &CoursePromptServiceImpl{}
}

// GetGreetingPrompt 取得問候 prompt 的描述。
func (s *CoursePromptServiceImpl) GetGreetingPrompt() mcp.Prompt {
	return mcp.NewPrompt("greeting",
		mcp.WithPromptDescription("A friendly greeting prompt for want use course mcp server"),
		mcp.WithArgument("name",
			mcp.ArgumentDescription("Name of the person to greet"),
		),
	)
}

// GetGreetingHandler 處理問候 prompt 請求。
func (s *CoursePromptServiceImpl) GetGreetingHandler(ctx context.Context, request mcp.GetPromptRequest) (*mcp.GetPromptResult, error) {
	name := request.Params.Arguments["name"]
	if name == "" {
		name = "friend"
	}
	log.Printf("Handling greeting prompt for %s\n", name) // Added logging

	return mcp.NewGetPromptResult(
		"A friendly greeting",
		[]mcp.PromptMessage{
			mcp.NewPromptMessage(
				mcp.RoleAssistant,
				mcp.NewTextContent(fmt.Sprintf("Hello, %s! How can I help you today? 我是一個課程系統的mcp server", name)),
			),
		},
	), nil
}

// GetOperationSuggestionPrompt 取得操作建議 prompt 的描述。
func (s *CoursePromptServiceImpl) GetOperationSuggestionPrompt() mcp.Prompt {
	return mcp.NewPrompt("操作建議",
		mcp.WithPromptDescription("操作建議"),
	)
}

// GetOperationSuggestionHandler 處理操作建議 prompt 請求。
func (s *CoursePromptServiceImpl) GetOperationSuggestionHandler(ctx context.Context, request mcp.GetPromptRequest) (*mcp.GetPromptResult, error) {
	log.Println("Handling operation suggestion prompt") // Added logging

	return mcp.NewGetPromptResult(
		"list course actions",
		[]mcp.PromptMessage{
			mcp.NewPromptMessage(
				mcp.RoleUser,
				mcp.NewTextContent("請列出課程系統可以操作的動作"),
			),
			mcp.NewPromptMessage(
				mcp.RoleAssistant,
				mcp.NewTextContent("我是一個課程系統的mcp server，可以操作的動作有："),
			),
			mcp.NewPromptMessage(
				mcp.RoleAssistant,
				mcp.NewTextContent("1. 建立課程"),
			),
			mcp.NewPromptMessage(
				mcp.RoleAssistant,
				mcp.NewTextContent("2. 查詢課程"),
			),
			mcp.NewPromptMessage(
				mcp.RoleAssistant,
				mcp.NewTextContent("3. 編輯課程"),
			),
		},
	), nil
}
