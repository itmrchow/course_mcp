// 此 interface 已移動至 internal/delivery/mcp/course_tool_interface.go
// 保留此檔案以避免編譯錯誤，請更新 import 路徑。

package tools

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"
)

type CourseMCPTool interface {
	// 取得課程
	GetCourseTool() mcp.Tool
	GetCourseHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error)

	// 建立課程
	CreateCourseTool() mcp.Tool
	CreateCourseHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error)

	// 查詢課程
	FindCourseTool() mcp.Tool
	FindCourseHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error)
}
