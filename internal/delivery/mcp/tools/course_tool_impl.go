package tools

import (
	"context"
	"encoding/json"
	"time"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/rs/zerolog"

	"course-mcp/internal/delivery/mcp/tools/dto"
)

var _ CourseMCPTool = (*CourseMCPToolImpl)(nil)

type CourseMCPToolImpl struct {
	logger *zerolog.Logger
}

// NewCourseMCPToolImpl 建立課程MCP Tool實例
func NewCourseMCPToolImpl(logger *zerolog.Logger) CourseMCPTool {
	return &CourseMCPToolImpl{
		logger: logger,
	}
}

// GetCourseTool 取得課程MCP Tool
func (c *CourseMCPToolImpl) GetCourseTool() mcp.Tool {
	return mcp.NewTool(
		string(ToolGetCourse),
		mcp.WithDescription("取得課程"),
		mcp.WithNumber("courseId",
			mcp.Required(),
			mcp.Description("課程ID"),
		),
	)
}

// GetCourseHandler 取得課程MCP Handler
func (c *CourseMCPToolImpl) GetCourseHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {

	args := request.GetArguments()
	courseId := args["courseId"].(float64)

	// TODO: call service
	if courseId == 1 {
		dto := dto.GetCourseResponseDTO{
			ID:                    1,
			Name:                  "課程1",
			Description:           "課程1描述",
			TeacherID:             "1",
			TeacherName:           "教師1",
			Price:                 100,
			MaxStudents:           100,
			MinStudents:           10,
			RegistrationStartDate: time.Now(),
			RegistrationEndDate:   time.Now(),
			StartDate:             time.Now(),
			EndDate:               time.Now(),
			IsOnline:              true,
			Status:                1,
			Note:                  "課程1備註",
		}

		jsonData, _ := json.Marshal(dto)

		return &mcp.CallToolResult{
			Content: []mcp.Content{
				mcp.TextContent{
					Type: "text",
					Text: string(jsonData),
				},
			},
		}, nil
	} else {
		return &mcp.CallToolResult{
			Content: []mcp.Content{
				mcp.TextContent{
					Type: "text",
					Text: `{"error": "course not found", "message": "course not found"}`,
				},
			},
		}, nil
	}

	// prepare response dto

}

// CreateCourseTool 建立課程MCP Tool
func (c *CourseMCPToolImpl) CreateCourseTool() mcp.Tool {
	return mcp.NewTool(
		string(ToolCreateCourse),
		mcp.WithDescription("建立課程"),
		mcp.WithString("name", mcp.Required(), mcp.Description("課程名稱")),
		mcp.WithString("description", mcp.Required(), mcp.Description("課程描述")),
		mcp.WithString("teacher_id", mcp.Required(), mcp.Description("教師ID")),
		mcp.WithNumber("price", mcp.Required(), mcp.Description("課程價格（新台幣）")),
		mcp.WithNumber("max_students", mcp.Required(), mcp.Description("最大學生人數")),
		mcp.WithNumber("min_students", mcp.Required(), mcp.Description("最小學生人數，最少1人")),
		mcp.WithString("registration_start_date", mcp.Required(), mcp.Description("報名開始時間（UTC時間）")),
		mcp.WithString("registration_end_date", mcp.Required(), mcp.Description("報名結束時間（UTC時間）")),
		mcp.WithString("start_date", mcp.Required(), mcp.Description("上課開始時間（UTC時間）")),
		mcp.WithString("end_date", mcp.Required(), mcp.Description("上課結束時間（UTC時間）")),
		mcp.WithBoolean("is_online", mcp.Required(), mcp.Description("是否是線上課程")),
		mcp.WithNumber("status", mcp.Required(), mcp.Description("課程狀態（0: 草稿, 1: 審核中）")),
		mcp.WithString("note", mcp.Required(), mcp.Description("課程備註")),
		mcp.WithArray("tags", mcp.Description("課程標籤")),
	)
}

// CreateCourseHandler 建立課程MCP Handler
func (c *CourseMCPToolImpl) CreateCourseHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	// 1. 取得參數（所有屬性都在最外層）
	args := request.GetArguments()

	// 2. 轉換為 DTO
	var dto dto.CreateCourseRequestDTO
	courseJson, _ := json.Marshal(args)
	if err := json.Unmarshal(courseJson, &dto); err != nil {
		return &mcp.CallToolResult{
			Content: []mcp.Content{
				mcp.TextContent{
					Type: "text",
					Text: `{"error": "invalid arguments", "message": "course object format error"}`,
				},
			},
		}, nil
	}

	// 3. 欄位驗證
	// TODO: 欄位驗證（檢查必填、型別、範圍等）

	// 4. 呼叫建立課程服務
	// TODO: call create course

	c.logger.Info().Msgf("create course dto: %+v", dto)
	// 5. 回傳 response
	return &mcp.CallToolResult{
		Content: []mcp.Content{
			mcp.TextContent{
				Type: "text",
				Text: `{"success": true, "message": "課程建立成功"}`,
			},
		},
	}, nil
}

func (c *CourseMCPToolImpl) FindCourseTool() mcp.Tool {
	return mcp.NewTool(
		string(ToolFindCourse),
		mcp.WithDescription("查詢課程"),
		mcp.WithDescription("將使用者輸入的內容,轉成1~3個tags"),
		mcp.WithString("course_id", mcp.Description("課程ID")),
		mcp.WithString("course_name", mcp.Description("課程名稱")),
		mcp.WithString("teacher_name", mcp.Description("教師名稱")),
		mcp.WithString("teacher_id", mcp.Description("教師ID")),
		mcp.WithString("tags"),
	)
}

func (c *CourseMCPToolImpl) FindCourseHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	args := request.GetArguments()

	c.logger.Info().Msgf("find course args: %+v", args)

	// 修改DTO
	dto := dto.GetCourseResponseDTO{
		ID:                    1,
		Name:                  "課程1",
		Description:           "課程1描述",
		TeacherID:             "1",
		TeacherName:           "教師1",
		Price:                 100,
		MaxStudents:           100,
		MinStudents:           10,
		RegistrationStartDate: time.Now(),
		RegistrationEndDate:   time.Now(),
		StartDate:             time.Now(),
		EndDate:               time.Now(),
		IsOnline:              true,
		Status:                1,
		Note:                  "課程1備註",
	}

	jsonData, _ := json.Marshal(dto)

	return &mcp.CallToolResult{
		Content: []mcp.Content{
			mcp.TextContent{
				Type: "text",
				Text: string(jsonData),
			},
		},
	}, nil
}
