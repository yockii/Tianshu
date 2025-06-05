# 第一阶段设计文档

## 一、技术选型
- 语言：Go
- Web框架：gofiber v2
- ORM：gorm
- 数据库：PostgreSQL
- 缓存：Redis（redigo）

## 二、第一阶段功能范围
1. 多租户与定制化
2. 用户与权限体系

---

## 三、功能需求细化

### 1. 多租户与定制化
- 租户注册/登录/信息维护（支持租户独立域名、Logo、主题色、登录界面、欢迎语等自定义）
- 租户数据隔离（数据库字段隔离，确保数据安全）
- 租户自定义业务流程/表单/界面布局（预留接口，第一阶段以租户基本信息和外观定制为主）

### 2. 用户与权限体系
- 用户注册/登录/找回密码（支持邮箱/手机号/用户名）
- 用户分组与角色管理（如管理员、普通用户、访客等）
- 多级权限控制（租户级、功能级、数据级、字段级权限，第一阶段以租户级和功能级为主）
- 操作日志审计（用户登录、登出、关键操作行为记录）

---

## 四、系统架构设计

### 1. 总体架构
- Fiber v2 作为 HTTP 服务入口，采用 RESTful API 设计
- gorm 负责数据持久化，PostgreSQL 作为主数据库，租户隔离
- redigo 作为 Redis 客户端，负责会话、缓存、验证码等
- 模块分层：
  - handler（API入口）
  - service（业务逻辑）
  - repository（数据访问）
  - model（数据结构）
  - middleware（中间件：认证、租户识别、日志等）

### 2. 主要数据表设计（gorm自动迁移）
- tenants（租户表）：id, name, logo, theme, domain, welcome_text, custom_config(JSONB), created_at, ...
- tenant_customizations（租户定制表）：id, tenant_id, logo, site_name, theme_color, favicon, extra_config(JSONB), created_at, ...
- users（用户表）：id, tenant_id, username, email, phone, password_hash, status, is_super_admin(bool), created_at, ...
- roles（角色表）：id, tenant_id, name, description, is_default(bool), created_at, ...
- permissions（权限表）：id, code, description, created_at, ...
- role_permissions（角色权限表）：id, role_id, permission_id
- user_roles（用户角色表）：id, user_id, role_id
- operation_logs（操作日志表）：id, tenant_id, user_id, action, detail, created_at, ...

#### 说明：
- gorm自动迁移所有表结构，支持字段变更和索引自动维护。
- 每个租户有一个不可删除的超级管理员账号（is_super_admin=true），系统自动创建，拥有最大权限。
- 租户可自定义角色，角色可分配权限，用户可分配多个角色。
- 租户定制化内容（如logo、网站名称、主题色等）存储于tenant_customizations表，支持JSONB扩展。

### 3. 关键接口设计（示例）
- POST   /api/tenant/register         // 租户注册（自动生成超级管理员账号）
- POST   /api/tenant/login           // 租户登录
- GET    /api/tenant/profile         // 获取租户信息及定制化内容
- PUT    /api/tenant/profile         // 修改租户信息及定制化内容
- POST   /api/user/register          // 用户注册
- POST   /api/user/login             // 用户登录
- GET    /api/user/profile           // 获取用户信息
- PUT    /api/user/profile           // 修改用户信息
- GET    /api/user/list              // 用户列表（支持分页、搜索）
- POST   /api/role                   // 新建角色
- GET    /api/role/list              // 角色列表
- POST   /api/role/assign            // 用户分配角色
- POST   /api/permission/assign      // 角色分配权限
- GET    /api/permission/list        // 权限列表
- GET    /api/logs                   // 操作日志查询

#### 说明：
- 超级管理员账号不可删除，删除接口需校验。
- 所有接口需校验租户隔离和权限。

---

## 五、部署与运维建议
- 支持Docker部署，环境变量配置数据库、Redis、端口等
- 日志统一输出，便于监控与追踪
- 预留多租户扩展能力，后续可平滑升级

---
如需详细ER图、接口文档或代码结构示例，可继续细化。
