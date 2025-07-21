package tools

import (
	"github.com/mark3labs/mcp-go/server"
)

// RegisterCourseTools 註冊課程工具
func RegisterCourseTools(s *server.MCPServer, tool CourseMCPTool) {
	s.AddTool(tool.GetCourseTool(), tool.GetCourseHandler)       // 取得課程
	s.AddTool(tool.CreateCourseTool(), tool.CreateCourseHandler) // 建立課程
	s.AddTool(tool.FindCourseTool(), tool.FindCourseHandler)     // 查詢課程
}

// func WrapWithScopeCheck(scope string, handlerFunc server.ToolHandlerFunc) server.ToolHandlerFunc {
// 	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
// 		tokenClaims, _ := ctx.Value("tokenClaims").(*utils.TokenClaims) // 取得 tokenClaims

// 		if tokenClaims.Scope == "" {
// 			return handlerFunc(ctx, request)
// 		}

// 		// 沒有權限
// 		if !strings.Contains(tokenClaims.Scope, scope) {
// 			// TODO 回傳錯誤 + return
// 			return &mcp.CallToolResult{}
// 		}

// 		return handlerFunc(ctx, request)
// 	}
// }
