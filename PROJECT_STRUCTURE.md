# 项目结构说明（第一阶段）

本项目采用 Go 语言最佳实践进行目录结构设计，结合 Fiber v2、gorm、redigo、PostgreSQL、Redis。

## 目录结构

- cmd/                // 各服务入口（如 main.go）
- internal/           // 业务核心代码（不可被外部依赖）
  - handler/          // HTTP API 路由与处理器
  - service/          // 业务逻辑层
  - repository/       // 数据访问层
  - model/            // 数据结构与ORM模型
  - middleware/       // 中间件（认证、租户识别、日志等）
  - config/           // 配置加载
  - utils/            // 工具函数
- pkg/                // 可被外部依赖的通用包（如JWT、邮件、验证码等）
- scripts/            // 启动、部署、数据库迁移等脚本
- build/              // Dockerfile、CI/CD等构建相关
- docs/               // 项目文档
- .env                // 环境变量配置
- go.mod/go.sum       // Go依赖管理

---

后续将在上述结构下逐步实现第一阶段的功能模块。
