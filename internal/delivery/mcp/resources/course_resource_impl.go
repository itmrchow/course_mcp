package resources

import (
	"context"
	"log"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/rs/zerolog"
)

var _ CourseMCPResource = (*CourseMCPResourceImpl)(nil)

type CourseMCPResourceImpl struct {
	logger *zerolog.Logger
}

func NewCourseMCPResourceImpl(logger *zerolog.Logger) CourseMCPResource {
	return &CourseMCPResourceImpl{
		logger: logger,
	}
}

func (r *CourseMCPResourceImpl) GetCourseResource() mcp.Resource {
	return mcp.NewResource(
		"docs://get-course-response",
		"Get Course",
		mcp.WithResourceDescription("Get Course Response"),
		mcp.WithMIMEType("application/json"),
	)
}

func (r *CourseMCPResourceImpl) GetCourseResourceHandler(ctx context.Context, request mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {
	log.Printf("Handling Course Response Format request for URI: %s", request.Params.URI)

	// 建立完整的回應格式文檔，包含範例
	responseFormat := `{
	"schema": ` + CourseResponseSchema + `,
	"example": {
		"id": 1,
		"name": "課程1",
		"description": "課程1描述",
		"price": 100.0,
		"maxStudents": 100,
		"minStudents": 10,
		"registrationStartDate": "2024-01-01T00:00:00Z",
		"registrationEndDate": "2024-12-31T23:59:59Z",
		"startDate": "2024-02-01T00:00:00Z",
		"endDate": "2024-11-30T23:59:59Z",
		"isOnline": true,
		"status": 1,
		"note": "課程1備註"
	},
	"fieldNotes": {
		"price": "價格單位為新台幣（TWD）",
		"dates": "所有日期欄位均使用 UTC 時間格式（ISO 8601）",
		"status": {
			"0": "草稿 - 課程尚未完成設定",
			"1": "審核中 - 課程正在等待管理員審核",
			"2": "開放報名 - 學生可以開始報名",
			"3": "已結束 - 課程已結束，不再接受報名",
			"4": "暫停報名 - 課程暫時停止接受報名"
		}
	},
	"errorFormat": {
		"error": "course not found",
		"message": "course not found"
	}
}`

	log.Printf("Successfully generated Course Response Format (%d bytes)", len(responseFormat))

	return []mcp.ResourceContents{
		mcp.TextResourceContents{
			URI:      "docs://course-response-format",
			MIMEType: "application/json",
			Text:     responseFormat,
		},
	}, nil
}

const CourseResponseSchema = `{
    "type": "object",
    "properties": {
        "id": {"type": "integer", "description": "課程唯一識別碼"},
        "name": {"type": "string", "description": "課程名稱"},
        "description": {"type": "string", "description": "課程描述"},
        "price": {"type": "number", "description": "課程價格（新台幣）"},
        "maxStudents": {"type": "integer", "description": "最大學生人數"},
        "minStudents": {"type": "integer", "description": "最小學生人數"},
        "registrationStartDate": {"type": "string", "format": "date-time", "description": "報名開始日期（UTC時間）"},
        "registrationEndDate": {"type": "string", "format": "date-time", "description": "報名結束日期（UTC時間）"},
        "startDate": {"type": "string", "format": "date-time", "description": "課程開始日期（UTC時間）"},
        "endDate": {"type": "string", "format": "date-time", "description": "課程結束日期（UTC時間）"},
        "isOnline": {"type": "boolean", "description": "是否為線上課程"},
        "status": {
            "type": "integer", 
            "description": "課程狀態",
            "enum": [0, 1, 2, 3, 4],
            "enumDescriptions": {
                "0": "草稿",
                "1": "審核中", 
                "2": "開放報名",
                "3": "已結束",
                "4": "暫停報名"
            }
        },
        "note": {"type": "string", "description": "課程備註"}
    },
    "required": ["id", "name", "description", "price", "maxStudents", "minStudents", "registrationStartDate", "registrationEndDate", "startDate", "endDate", "isOnline", "status"]
}`
