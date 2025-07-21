package tools

// ToolName 定義工具名稱
type ToolName string

const (
	// 課程工具
	ToolGetCourse    ToolName = "getCourse"
	ToolFindCourse   ToolName = "findCourse"
	ToolCreateCourse ToolName = "createCourse"
	ToolUpdateCourse ToolName = "updateCourse"

	// 報名工具
	ToolGetCourseRegistration    ToolName = "getCourseRegistration"
	ToolFindCourseRegistration   ToolName = "findCourseRegistration"
	ToolCreateCourseRegistration ToolName = "createCourseRegistration"

	// 教師工具
	ToolGetTeacher    ToolName = "getTeacher"
	ToolFindTeacher   ToolName = "findTeacher"
	ToolCreateTeacher ToolName = "createTeacher"
	ToolUpdateTeacher ToolName = "updateTeacher"
)