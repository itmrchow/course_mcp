package tools

import (
	"context"
	"encoding/json"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/rs/zerolog"

	"course-mcp/internal/delivery/mcp/tools/dto"
)

var _ TeacherMCPTool = (*TeacherMCPToolImpl)(nil)

type TeacherMCPToolImpl struct {
	logger *zerolog.Logger
}

// NewTeacherMCPToolImpl 建立教師MCP Tool實例
func NewTeacherMCPToolImpl(logger *zerolog.Logger) TeacherMCPTool {
	return &TeacherMCPToolImpl{
		logger: logger,
	}
}

// GetTeacherTool 取得教師MCP Tool
func (t *TeacherMCPToolImpl) GetTeacherTool() mcp.Tool {
	return mcp.NewTool(
		"getTeacher",
		mcp.WithDescription("取得教師"),
		mcp.WithNumber("teacherId",
			mcp.Required(),
			mcp.Description("教師ID"),
		),
	)
}

// GetTeacherHandler 取得教師MCP Handler
func (t *TeacherMCPToolImpl) GetTeacherHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	args := request.GetArguments()
	teacherId := args["teacherId"].(float64)

	t.logger.Info().Msgf("get teacher id: %f", teacherId)

	// TODO: call service to get teacher by id
	return &mcp.CallToolResult{
		Content: []mcp.Content{
			mcp.TextContent{
				Type: "text",
				Text: `{"message": "get teacher service not implemented yet"}`,
			},
		},
	}, nil
}

// CreateTeacherTool 建立教師MCP Tool
func (t *TeacherMCPToolImpl) CreateTeacherTool() mcp.Tool {
	return mcp.NewTool(
		"createTeacher",
		mcp.WithDescription("建立教師"),
		mcp.WithNumber("user_id", mcp.Required(), mcp.Description("關聯 user（帳號）ID")),
		mcp.WithString("name", mcp.Required(), mcp.Description("教師姓名")),
		mcp.WithString("phone", mcp.Required(), mcp.Description("聯絡電話")),
		mcp.WithString("email", mcp.Required(), mcp.Description("教師信箱")),
		mcp.WithString("bio", mcp.Required(), mcp.Description("教師簡介/自我介紹")),
		mcp.WithNumber("status", mcp.Required(), mcp.Description("狀態（0: 審核中, 1: 審核通過, 2: 審核失敗, 3: 已停用）")),
	)
}

// CreateTeacherHandler 建立教師MCP Handler
func (t *TeacherMCPToolImpl) CreateTeacherHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	// 1. 取得參數
	args := request.GetArguments()

	// 2. 轉換為 DTO
	var teacherDTO dto.CreateTeacherRequestDTO
	teacherJson, _ := json.Marshal(args)
	if err := json.Unmarshal(teacherJson, &teacherDTO); err != nil {
		return &mcp.CallToolResult{
			Content: []mcp.Content{
				mcp.TextContent{
					Type: "text",
					Text: `{"error": "invalid arguments", "message": "teacher object format error"}`,
				},
			},
		}, nil
	}

	// 3. 欄位驗證
	// TODO: 欄位驗證（檢查必填、型別、範圍等）

	// 4. 呼叫建立教師服務
	// TODO: call create teacher service

	t.logger.Info().Msgf("create teacher dto: %+v", teacherDTO)

	// 5. 回傳 response
	return &mcp.CallToolResult{
		Content: []mcp.Content{
			mcp.TextContent{
				Type: "text",
				Text: `{"success": true, "message": "教師建立成功"}`,
			},
		},
	}, nil
}

// FindTeacherTool 查詢教師MCP Tool
func (t *TeacherMCPToolImpl) FindTeacherTool() mcp.Tool {
	return mcp.NewTool(
		"findTeacher",
		mcp.WithDescription("查詢教師"),
		mcp.WithString("teacher_name", mcp.Description("教師姓名")),
	)
}

// FindTeacherHandler 查詢教師MCP Handler
func (t *TeacherMCPToolImpl) FindTeacherHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	args := request.GetArguments()

	// 轉換為 DTO
	var findDTO dto.FindTeacherRequestDTO
	findJson, _ := json.Marshal(args)
	if err := json.Unmarshal(findJson, &findDTO); err != nil {
		return &mcp.CallToolResult{
			Content: []mcp.Content{
				mcp.TextContent{
					Type: "text",
					Text: `{"error": "invalid arguments", "message": "find teacher request format error"}`,
				},
			},
		}, nil
	}

	t.logger.Info().Msgf("find teacher dto: %+v", findDTO)

	// TODO: call service to find teachers
	return &mcp.CallToolResult{
		Content: []mcp.Content{
			mcp.TextContent{
				Type: "text",
				Text: `{"message": "find teacher service not implemented yet"}`,
			},
		},
	}, nil
}