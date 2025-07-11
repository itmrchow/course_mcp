package tools

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"
)

type TeacherMCPTool interface {
	// 取得教師
	GetTeacherTool() mcp.Tool
	GetTeacherHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error)

	// 建立教師
	CreateTeacherTool() mcp.Tool
	CreateTeacherHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error)

	// 查詢教師
	FindTeacherTool() mcp.Tool
	FindTeacherHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error)
}