cursor-mcp
mcp server , 透過mcp sse操作課程系統



# tools
| Tools Name         | 功能                    | Role | CheckPoint |
| ------------------ | ----------------------- | ---- | ---------- |
| GetCourse          | 取得課程 , 根據課程編號 |      | v          |
| CreateCourse       | 建立課程                |     |v            |
| FindCourse         | 查詢課程                |     | v           |
| UpdateCourse       | 更新課程                |      |            |
| CreateRegistration | 建立報名                |      |            |
| GetTeacher         | 取得老師 , 根據課程編號 |      |            |
| FindTeacher        | 查詢老師                |      |            |
| CreateTeacher      | 建立老師                |      |            |


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
