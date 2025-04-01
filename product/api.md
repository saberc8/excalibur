## 接口 api

以下是按照新规则调整后的接口设计文档（Markdown 格式）：

# 项目管理系统接口设计文档

## 1. 接口设计规则

- **GET 请求**：仅用于获取资源列表或单个资源（通过查询参数指定 ID）
- **POST 请求**：用于所有其他操作（创建、更新、删除、关联等）
- **路径规范**：
  - 列表接口：`/space/list`
  - 单个资源获取：`/space?{id}`
  - 操作接口：`/space/action`（如`/user/create`）

## 2. 空间管理接口

| 路径            | 方法 | 说明         | 参数（示例）                    |
| --------------- | ---- | ------------ | ------------------------------- |
| `/space/list`   | GET  | 获取空间列表 | `page=1&page_size=10`           |
| `/space`        | GET  | 获取单个空间 | `id=sp_123`                     |
| `/space/create` | POST | 创建空间     | `name=研发空间&description=...` |
| `/space/update` | POST | 更新空间     | `id=sp_123&name=新名称`         |
| `/space/delete` | POST | 删除空间     | `id=sp_123`                     |

## 3. 项目管理接口

| 路径              | 方法 | 说明         | 参数（示例）                  |
| ----------------- | ---- | ------------ | ----------------------------- |
| `/project/list`   | GET  | 获取项目列表 | `space_id=sp_123`             |
| `/project`        | GET  | 获取单个项目 | `id=prj_456`                  |
| `/project/create` | POST | 创建项目     | `space_id=sp_123&name=新项目` |
| `/project/update` | POST | 更新项目     | `id=prj_456&owner=user_789`   |
| `/project/delete` | POST | 删除项目     | `id=prj_456`                  |

## 4. 任务管理接口

| 路径           | 方法 | 说明         | 参数（示例）                        |
| -------------- | ---- | ------------ | ----------------------------------- |
| `/task/list`   | GET  | 获取任务列表 | `project_id=prj_456&status=进行中`  |
| `/task`        | GET  | 获取单个任务 | `id=tsk_789`                        |
| `/task/create` | POST | 创建任务     | `project_id=prj_456&title=紧急任务` |
| `/task/update` | POST | 更新任务状态 | `id=tsk_789&status=已完成`          |
| `/task/delete` | POST | 删除任务     | `id=tsk_789`                        |

## 6. 其他接口

| 路径           | 方法 | 说明     | 参数（示例）                   |
| -------------- | ---- | -------- | ------------------------------ |
| `/search/task` | POST | 搜索任务 | `keyword=紧急&space_id=sp_123` |
