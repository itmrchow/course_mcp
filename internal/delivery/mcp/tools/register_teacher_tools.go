package tools

import (
	"github.com/mark3labs/mcp-go/server"
)

// RegisterTeacherTools 註冊教師工具
func RegisterTeacherTools(s *server.MCPServer, tool TeacherMCPTool) {
	s.AddTool(tool.GetTeacherTool(), tool.GetTeacherHandler)       // 取得教師
	s.AddTool(tool.CreateTeacherTool(), tool.CreateTeacherHandler) // 建立教師
	s.AddTool(tool.FindTeacherTool(), tool.FindTeacherHandler)     // 查詢教師
}