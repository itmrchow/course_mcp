package tools

import (
	"github.com/mark3labs/mcp-go/server"
)

// RegisterCourseTools 註冊課程工具
func RegisterCourseTools(s *server.MCPServer, tool CourseMCPTool) {
	s.AddTool(tool.GetCourseTool(), tool.GetCourseHandler)       // 取得課程
	s.AddTool(tool.CreateCourseTool(), tool.CreateCourseHandler) // 建立課程
}
