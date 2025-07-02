package dto

import "time"

type CreateCourseRequestDTO struct {
	Name                  string    `json:"name"`                    // 課程名稱
	Description           string    `json:"description"`             // 課程描述
	TeacherID             string    `json:"teacher_id"`              // 教師ID
	Price                 uint      `json:"price"`                   // 課程價格 , 新台幣
	MaxStudents           uint      `json:"max_students"`            // 最大學生人數
	MinStudents           uint      `json:"min_students"`            // 最小學生人數 , 最少1人
	RegistrationStartDate time.Time `json:"registration_start_date"` // 報名開始時間
	RegistrationEndDate   time.Time `json:"registration_end_date"`   // 報名結束時間
	StartDate             time.Time `json:"start_date"`              // 上課開始時間
	EndDate               time.Time `json:"end_date"`                // 上課結束時間
	IsOnline              bool      `json:"is_online"`               // 是否是線上課程
	Status                uint      `json:"status"`                  // 0: 草稿, 1: 審核中, 2: 開放報名, 3: 已結束 , 4: 暫停報名
	Note                  string    `json:"note"`                    // 課程備註
}
