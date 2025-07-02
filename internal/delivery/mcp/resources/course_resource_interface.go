package resources

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"
)

type CourseMCPResource interface {
	GetCourseResource() mcp.Resource
	GetCourseResourceHandler(ctx context.Context, request mcp.ReadResourceRequest) ([]mcp.ResourceContents, error)
}
