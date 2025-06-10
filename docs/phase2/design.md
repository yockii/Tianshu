# 第二阶段设计文档

## 一、阶段目标
第二阶段聚焦大疆上云API核心集成，实现无人机与机场设备的接入与管理、实时数据与视频流、智能航线与任务编排、机场(Dock)管理等功能。

## 二、技术选型与延续
- 语言：Go 1.25
- Web 框架：Fiber v2
- ORM：GORM（自动迁移、JSONB 索引、外键）
- 数据库：PostgreSQL
- 缓存/会话：Redis (redigo)，JWT 存储 sessionKey
- 配置管理：Viper (统一通过 `config.Cfg` 获取)
- 日志：Fiber 日志 + OperationLogService 写入 `operation_logs` 表
- 密码哈希：bcrypt
- 消息协议：Mochi-MQTT（内嵌 MQTT 客户端），启动时自动连接至大疆 MQTT 网关并订阅主题
- 前端：Vue3 + TypeScript + Vite + Element Plus（已有基础架构）

## 三、功能模块与设计

### 3.1 设备接入与管理
- **Pilot 与 Dock 设备注册**：对接大疆上云API，支持设备批量注册、分组、标签管理
- **状态监控**：周期性或实时获取设备飞行状态、电量、GPS、信号强度等；存入时序数据表或缓存
- **OTA 升级与参数配置**：远程下发固件升级包、组件参数调整、预设方案保存与回滚
- **设备日志与告警**：采集设备端日志，支持搜索与导出；告警规则引擎与历史记录

### 3.2 实时数据与视频流
- **遥测数据订阅与推送**：使用 WebSocket 或 MQTT（Mochi-MQTT）订阅设备上云推送的数据，服务端分发至前端与后端消费
- **历史数据回放**：按时间区间查询存储的飞行轨迹、传感器数据；支持分页与导出
- **视频流接入与转发**：支持 RTMP/RTSP/WebRTC 协议流媒体接入，集成大疆推流鉴权；前端播放组件封装

### 3.3 智能航线与任务编排
- **航线规划与解析**：前端地图绘制航线，后端兼容大疆航线格式导入导出
- **任务下发与调度**：支持定时、循环、应急任务，下发到指定设备或组；任务优先级与并发策略
- **任务监控与断点续飞**：实时跟踪任务状态，异常断点自动重试或切换备降点
- **任务历史与日志**：记录任务执行详情与事件流，支持分页检索与导出

### 3.4 机场 (Dock) 管理
- **机场注册与分组**：与 Pilot 类似，实现机场设备元数据管理
- **远程控制命令**：开关舱门、充电管理、重启控制，接入大疆 Dock API
- **机场状态与告警**：实时监控机场健康状态、环境数据（温度、湿度等）、告警推送与存储
- **绑定关系管理**：一键绑定/解绑飞机与机场，支持批量操作与权限校验

## 四、系统架构与接口设计
- 继续采用分层架构：Middleware → Handler → Service → Repository → Model
- 新增 RPC/WebSocket 层：处理数据订阅推送与流媒体转发
- 数据库表新增：
  - `devices`、`device_status`、`ota_tasks`、`device_logs`
  - `video_streams`、`stream_sessions`
  - `flight_plans`、`missions`
  - `docks`、`dock_commands`、`dock_logs`

### 关键 RESTful 接口示例
- POST   /api/device/register
- GET    /api/device/status/{id}
- POST   /api/device/ota
- GET    /api/video/streams
- POST   /api/mission
- GET    /api/mission/{id}/status
- POST   /api/dock/{id}/command

## 五、数据模型与迁移
- 编写 GORM Migration，确保 JSONB 索引与外键约束生效
- 设计时序数据存储策略：分表或按月分区

## 六、前端设计
- 复用现有布局与权限控制，新增模块菜单与路由
- 设备列表、实时监控面板、视频播放页、任务管理页、机场运维页
- 使用 Pinia 状态管理订阅数据与 WebSocket 连接
- Pilot/Dock 登录界面：新增设备侧登录页面，集成大疆 JSBridge，支持在飞控端或机场端内部加载并调用大疆原生 API 进行鉴权与授权

## 七、测试与部署
- 单元测试和集成测试：覆盖新 Service/Handler，模拟大疆 API 返回
- 前端 E2E：设备接入与任务流程
- Docker 化：补充 second-stage Dockerfile 与 GitHub Actions 流水线

---
## 下一步
请审阅以上第二阶段设计，与第一阶段设计保持一致的结构与规范。如有补充、调整或优先级建议，欢迎反馈。
