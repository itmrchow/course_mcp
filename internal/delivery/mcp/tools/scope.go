package tools

import (
	"context"
	"strings"

	"github.com/mark3labs/mcp-go/mcp"

	"course-mcp/internal/usecase/utils"
)

// ToolScope 定義可以使用哪些mcp tools , resources, prompts
type ToolScope string

const (
	// 課程相關權限
	ScopeCourseBasic ToolScope = "course_basic" // 課程讀取
	ScopeCourseEdit  ToolScope = "course_edit"  // 課程編輯

	// 報名相關權限
	ScopeCourseRegistrationBasic ToolScope = "course_registration_basic" // 報名讀取
	ScopeCourseRegistrationEdit  ToolScope = "course_registration_edit"  // 報名編輯

	// 教師相關權限
	ScopeTeacherBasic ToolScope = "teacher_basic" // 老師讀取
	ScopeTeacherEdit  ToolScope = "teacher_edit"  // 老師編輯

	// 管理員權限
	ScopeAdmin ToolScope = "admin" // 管理員操作
)

// ScopeToolMapping 定義 scope 到 tools 的映射
var ScopeToolMapping = map[ToolScope][]ToolName{
	// 課程基本權限 - 課程讀取
	ScopeCourseBasic: {
		ToolGetCourse,
		ToolFindCourse,
	},

	// 課程編輯權限
	ScopeCourseEdit: {
		ToolCreateCourse,
		ToolUpdateCourse,
	},

	// 報名基本權限 - 報名讀取
	ScopeCourseRegistrationBasic: {
		ToolGetCourseRegistration,
		ToolFindCourseRegistration,
	},

	// 報名編輯權限
	ScopeCourseRegistrationEdit: {
		ToolCreateCourseRegistration,
		ToolFindTeacher, // 根據你的定義，報名編輯包含FindTeacher
	},

	// 教師基本權限 - 老師讀取
	ScopeTeacherBasic: {
		ToolGetTeacher,
		ToolFindTeacher,
	},

	// 教師編輯權限
	ScopeTeacherEdit: {
		ToolCreateTeacher,
		ToolUpdateTeacher,
	},

	// 管理員權限 - 根據你的定義，目前只包含教師相關
	ScopeAdmin: {
		ToolGetTeacher,
		ToolFindTeacher,
	},
}

// ScopeManager 管理權限範圍
type ScopeManager struct{}

// NewScopeManager 創建新的權限管理器
func NewScopeManager() *ScopeManager {
	return &ScopeManager{}
}

// GetScopeTools 根據 scope 獲取可用工具
func (sm *ScopeManager) GetScopeTools(scope ToolScope) []ToolName {
	if tools, exists := ScopeToolMapping[scope]; exists {
		return tools
	}
	return []ToolName{}
}

// GetUserAllowedTools 根據用戶的scopes獲取所有可用工具
func (sm *ScopeManager) GetUserAllowedTools(scopes []ToolScope) []ToolName {
	toolSet := make(map[ToolName]bool)

	for _, scope := range scopes {
		tools := sm.GetScopeTools(scope)
		for _, tool := range tools {
			toolSet[tool] = true
		}
	}

	var allowedTools []ToolName
	for tool := range toolSet {
		allowedTools = append(allowedTools, tool)
	}

	return allowedTools
}

// FilterToolsByScope 根據用戶權限過濾工具列表
func (sm *ScopeManager) FilterToolsByScope(ctx context.Context, tools []mcp.Tool) []mcp.Tool {
	// 取得 token claims
	tokenClaims, isOk := utils.GetTokenClaims(ctx)
	if !isOk {
		return []mcp.Tool{} // 無權限時返回空列表
	}

	// 解析 scope 字串
	var userScopes []ToolScope
	if tokenClaims.Scope != "" {
		scopeStrs := strings.Split(tokenClaims.Scope, " ")
		for _, scopeStr := range scopeStrs {
			userScopes = append(userScopes, ToolScope(strings.TrimSpace(scopeStr)))
		}
	}

	// 如果沒有 scope，返回空列表
	if len(userScopes) == 0 {
		return []mcp.Tool{}
	}

	// 獲取所有允許的工具
	allowedToolSet := make(map[string]bool)
	for _, scope := range userScopes {
		tools := sm.GetScopeTools(scope)
		for _, tool := range tools {
			allowedToolSet[string(tool)] = true
		}
	}

	// 過濾工具列表
	filteredTools := make([]mcp.Tool, 0)
	for _, tool := range tools {
		if allowedToolSet[tool.Name] {
			filteredTools = append(filteredTools, tool)
		}
	}

	return filteredTools
}

// HasToolPermission 檢查是否有特定工具的權限
func (sm *ScopeManager) HasToolPermission(ctx context.Context, toolName string) bool {
	tokenClaims, isOk := utils.GetTokenClaims(ctx)
	if !isOk {
		return false
	}

	// 解析 scope
	var userScopes []ToolScope
	if tokenClaims.Scope != "" {
		scopeStrs := strings.Split(tokenClaims.Scope, " ")
		for _, scopeStr := range scopeStrs {
			userScopes = append(userScopes, ToolScope(strings.TrimSpace(scopeStr)))
		}
	}

	// 檢查權限
	for _, scope := range userScopes {
		tools := sm.GetScopeTools(scope)
		for _, tool := range tools {
			if string(tool) == toolName {
				return true
			}
		}
	}

	return false
}

// GetScopeDescription 獲取scope的描述
func (sm *ScopeManager) GetScopeDescription(scope ToolScope) string {
	descriptions := map[ToolScope]string{
		ScopeCourseBasic:             "課程讀取",
		ScopeCourseEdit:              "課程編輯",
		ScopeCourseRegistrationBasic: "報名讀取",
		ScopeCourseRegistrationEdit:  "報名編輯",
		ScopeTeacherBasic:            "老師讀取",
		ScopeTeacherEdit:             "老師編輯",
		ScopeAdmin:                   "管理員操作",
	}

	if desc, exists := descriptions[scope]; exists {
		return desc
	}
	return "未知權限"
}
