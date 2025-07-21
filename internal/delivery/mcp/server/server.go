package server

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
	"github.com/rs/zerolog"

	"course-mcp/internal/delivery/mcp/scope"
	"course-mcp/internal/usecase/utils"
)

// Server package for MCP

// NewMCPServer 建立新的 MCP 伺服器實例
func NewMCPServer(
	logger *zerolog.Logger,
	manager *MCPServerManager,
) *server.MCPServer {
	return server.NewMCPServer(
		"course-mcp-server", // Server name
		"1.0.0",             // Server version
		server.WithResourceCapabilities(true, true),
		server.WithToolCapabilities(true),
		server.WithLogging(),
		server.WithRecovery(),
		server.WithToolFilter(manager.ToolFilterFunc()),
		server.WithHooks(&server.Hooks{
			OnAfterListTools: []server.OnAfterListToolsFunc{
				func(ctx context.Context, id any, req *mcp.ListToolsRequest, result *mcp.ListToolsResult) {

					tokenClaims, isOk := utils.GetTokenClaims(ctx)
					if !isOk {
						logger.Error().Msg("Token claims not found in context")
						return
					}

					logger.Info().Any("scope", tokenClaims.Scope).Msg("Token claims scope")

					// 這裡可以添加額外的日誌或處理
					logger.Info().Msgf("Tools listed: %v", result.Tools)
				},
			},
		}),
		server.WithToolHandlerMiddleware(manager.AuthorizationMiddleware()),
	)
}

type MCPServerManager struct {
	logger       *zerolog.Logger
	scopeManager *scope.ScopeManager
}

// NewMCPServerManager 創建新的 MCP 伺服器管理器 , 定義MCP Server所需的Handler funcs
func NewMCPServerManager(logger *zerolog.Logger, scopeManager *scope.ScopeManager) *MCPServerManager {
	return &MCPServerManager{
		logger:       logger,
		scopeManager: scopeManager,
	}
}

func (m *MCPServerManager) ToolFilterFunc() server.ToolFilterFunc {
	return func(ctx context.Context, tools []mcp.Tool) []mcp.Tool {
		return m.scopeManager.FilterToolsByScope(ctx, tools) // 使用 scope 管理器過濾工具
	}
}

// AuthorizationMiddleware 建立授權檢查中間件
func (m *MCPServerManager) AuthorizationMiddleware() server.ToolHandlerMiddleware {
	return func(next server.ToolHandlerFunc) server.ToolHandlerFunc {
		return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			toolName := request.Params.Name
			m.logger.Info().Msgf("Tool name: %s", toolName)

			// 使用 scopeManager 檢查工具權限
			if !m.scopeManager.HasToolPermission(ctx, toolName) {
				m.logger.Warn().Msgf("Access denied for tool: %s", toolName)
				return &mcp.CallToolResult{
					Content: []mcp.Content{
						mcp.TextContent{
							Type: "text",
							Text: `{"error": "forbidden", "code": 403, "message": "insufficient permissions"}`,
						},
					},
				}, nil
			}

			return next(ctx, request)
		}
	}
}
