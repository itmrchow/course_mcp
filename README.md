cursor-mcp
mcp server , 透過mcp sse操作課程系統

# User Roles
## 前台
| Role           | Description     |
| -------------- | --------------- |
| Student        | 學生 (預設權限) |
| Teacher        | 老師            |
| CompanyManager | 公司管理者      |

## 後台
| Role   | Description       |
| ------ | ----------------- |
| Admin  | 管理員 (預設權限) |
| Member | 成員              |


# tools
| Resource | Tools Name               | 功能                    | Role | CheckPoint |
| -------- | ------------------------ | ----------------------- | ---- | ---------- |
| Course   | GetCourse                | 取得課程 , 根據課程編號 |      | v          |
| Course   | CreateCourse             | 建立課程                |      | v          |
| Course   | FindCourse               | 查詢課程                |      | v          |
| Course   | UpdateCourse             | 更新課程                |      |            |
| Course   | CreateCourseRegistration | 建立報名                |      |            |
| Course   | FindCourseRegistration   | 查詢報名                |      |            |
| Course   | GetCourseRegistration    | 取得報名                |      |            |
| Teacher  | GetTeacher               | 取得老師 , 根據教師編號 |      | v          |
| Teacher  | FindTeacher              | 查詢老師                |      | v          |
| Teacher  | CreateTeacher            | 建立老師                |      | v          |

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
