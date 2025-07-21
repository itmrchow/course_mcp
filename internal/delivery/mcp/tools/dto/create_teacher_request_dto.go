package dto

type CreateTeacherRequestDTO struct {
	UserID uint   `json:"user_id"` // 關聯 user（帳號）ID
	Name   string `json:"name"`    // 教師姓名
	Phone  string `json:"phone"`   // 聯絡電話
	Email  string `json:"email"`   // 教師信箱
	Bio    string `json:"bio"`     // 教師簡介/自我介紹
	Status int    `json:"status"`  // 狀態 , 0: 審核中 , 1: 審核通過 , 2: 審核失敗 , 3: 已停用
}
