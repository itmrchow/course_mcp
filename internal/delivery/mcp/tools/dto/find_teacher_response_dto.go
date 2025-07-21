package dto

// FindTeacherResponseDTO 查詢教師回應 DTO
type FindTeacherResponseDTO struct {
	Teachers []GetTeacherResponseDTO `json:"teachers"` // 教師列表
	Total    int                     `json:"total"`    // 總數量
}