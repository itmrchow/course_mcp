package resources

import "github.com/mark3labs/mcp-go/server"

// 註冊課程資源
func RegisterCourseResources(s *server.MCPServer, resource CourseMCPResource) {
	s.AddResource(resource.GetCourseResource(), resource.GetCourseResourceHandler)
}
