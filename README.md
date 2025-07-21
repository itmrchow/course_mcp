course-mcp
mcp server , 透過mcp sse操作課程系統

# User Roles
## 前台
| Role                 | Description     | Scopes                                                                                                           |
| -------------------- | --------------- | ---------------------------------------------------------------------------------------------------------------- |
| role_student         | 學生 (預設權限) | course_basic , course_registration_basic , course_registration_edit , teacher_basic                              |
| role_teacher         | 老師            | course_basic , course_edit , teacher_edit , teacher_basic , course_registration_basic , course_registration_edit |
| role_company_manager | 公司管理者      |

## 後台
| Role   | Description       | Scopes |
| ------ | ----------------- | ------ |
| Admin  | 管理員 (預設權限) |
| Member | 成員              |

# Scopes
| Scope name                | Description | 功能                                             |
| ------------------------- | ----------- | ------------------------------------------------ |
| course_basic              | 課程讀取    | GetCourse<br> FindCourse                         |
| course_edit               | 課程編輯    | CreateCourse<br> UpdateCourse                    |
| course_registration_basic | 報名讀取    | GetCourseRegistration<br> FindCourseRegistration |
| course_registration_edit  | 報名編輯    | CreateCourseRegistration<br> FindTeacher         |
| teacher_basic             | 老師讀取    | GetTeacher<br> FindTeacher                       |
| teacher_edit              | 老師編輯    | CreateTeacher<br> UpdateTeacher                  |
| admin                     | 管理員操作  | GetTeacher<br> FindTeacher                       |

# tools
| Resource | Tools Name               | 功能                    | Role         | Check |
| -------- | ------------------------ | ----------------------- | ------------ | ---------- |
| Course   | GetCourse                | 取得課程 , 根據課程編號 | role_student | v          |
| Course   | CreateCourse             | 建立課程                |              | v          |
| Course   | FindCourse               | 查詢課程                |              | v          |
| Course   | UpdateCourse             | 更新課程                |              |            |
| Course   | CreateCourseRegistration | 建立報名                |              |            |
| Course   | UpdateCourseRegistration | 更新報名                |              |            |
| Course   | FindCourseRegistration   | 查詢報名                |              |            |
| Course   | GetCourseRegistration    | 取得報名                |              |            |
| Teacher  | GetTeacher               | 取得老師 , 根據教師編號 |              | v          |
| Teacher  | FindTeacher              | 查詢老師                |              | v          |
| Teacher  | CreateTeacher            | 建立老師                |              | v          |
| Teacher  | CreateTeacher            | 更新老師                |              |            |

# Workflow
前台
- Student
  - 報名: FindCourse -> CreateCourseRegistration
  - 查詢報名:
  - 取消報名:
- Teacher
  - 課程建立:
  - 課程修改:
  - 課程查詢:
後台
- 成員
  - 建立成員
  - 查詢成員
  - 刪除成員
  - 更新成員
- 用戶
  - 建立用戶
  - 查詢用戶
  - 刪除用戶
  - 更新用戶
- 課程
  - 新增課程
  - 修改課程
  - 查詢課程
  - 刪除課程

# config
## 設定方式
修改容器環境變數

## env
| 欄位名稱 | 型別 | 預設值 |
| -------- | ---- | ------ |
| PORT     | int  | 3000   |

# install

# todo

# pkgs
- github.com/mark3labs/mcp-go go mcp
- github.com/rs/zerolog logger
- ... other

# test tool
inspector
```
npx @modelcontextprotocol/inspector
```
