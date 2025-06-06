# 第一阶段设计文档

## 一、技术选型与框架
- 语言：Go 1.25
- Web 框架：Fiber v2
- ORM：GORM (自动迁移支持 JSONB 索引和外键)
- 数据库：PostgreSQL
- 缓存/会话：Redis (redigo)，会话数据存于 Redis，JWT 存储 sessionKey
- 配置管理：Viper (统一通过 `config.Cfg` 获取)
- 日志：Fiber 自带日志 + OperationLogService 写入 `operation_logs` 表
- 密码哈希：bcrypt

## 二、第一阶段功能范围与完成情况
| 功能模块               | 需求描述                                             | 实现情况         |
|------------------------|------------------------------------------------------|------------------|
| 多租户注册与登录       | 支持租户独立域名、Logo、主题色、欢迎语等基本定制             | 已完成           |
| 租户信息与定制维护     | 获取/修改租户 profile 与 JSONB 扩展配置                  | 已完成           |
| 用户注册/登录          | 支持邮箱/手机号/用户名注册登录，密码加密，JWT+Redis 会话       | 已完成           |
| 用户 Profile 管理      | 获取/更新/列表用户信息（分页、搜索）                      | 已完成           |
| 角色、权限管理         | CRUD 角色、权限；多角色分配；权限码检查；超级管理员绕过权限    | 已完成           |
| 关系维护 (User-Role、Role-Permission) | 关联增删查                                        | 已完成           |
| 权限中间件             | `AuthMiddleware` 验证 JWT+Redis，会话校验；`RequirePermission` 检查权限 | 已完成           |
| 操作日志审计           | 关键操作（用户/角色/权限/关联）写入日志表                    | 已完成           |
| 前端 Dashboard         | Vue3+Vite+TS，Element Plus，Pinia，动态菜单与权限控制         | 已完成           |
| CI/CD & 测试           | 单元测试／集成测试；前端 E2E；Docker 化；持续集成                | 未完成 (Pending) |


## 三、系统架构设计
### 1. 总体架构
- Fiber v2 作为 HTTP 服务入口，RESTful 风格
- 中间件层：认证（JWT+Redis 会话）、租户隔离、权限校验、操作日志
- Handler 层：路由与请求解析
- Service 层：核心业务逻辑与事务管理
- Repository 层：GORM 数据访问、自动迁移
- Model 层：实体定义，包含 JSON 与 GORM 标签
- Cache 层：Redis 会话与缓存访问

### 2. 数据库表设计
- `tenants`: 租户信息（基础信息，不含定制配置）
- `tenant_customizations`: 租户外观定制（logo、主题色、favicon、extra_config JSONB）
- `users`：用户表（tenant_id、用户名/邮箱/手机号、password_hash、is_super_admin）
- `roles`：角色表（默认角色标记）
- `permissions`：权限码表
- `user_roles`：用户-角色关联
- `role_permissions`：角色-权限关联
- `operation_logs`：操作日志（tenant_id、user_id、action、detail）

## 四、关键接口设计

### 4.1 租户相关
- POST   /api/tenant/register         租户注册（自动创建超级管理员）
- POST   /api/tenant/login            租户登录
- GET    /api/tenant/profile          获取租户信息及定制化
- PUT    /api/tenant/profile          修改租户信息及定制化

### 4.2 用户相关
- POST   /api/user/register           用户注册
- POST   /api/user/login              用户登录
- GET    /api/user/profile            获取当前用户
- PUT    /api/user/profile            更新当前用户
- GET    /api/user/list               用户列表（分页、搜索）

### 4.3 角色与权限
- POST   /api/role                    新建角色
- GET    /api/role/list               角色列表
- PUT    /api/role/{id}               更新角色
- DELETE /api/role/{id}               删除角色
- POST   /api/permission              新建权限
- GET    /api/permission/list         权限列表
- PUT    /api/permission/{id}         更新权限
- DELETE /api/permission/{id}         删除权限

### 4.4 关系管理
- GET    /api/relation/user-roles     查询用户角色
- POST   /api/relation/user-roles     分配角色
- DELETE /api/relation/user-roles     删除用户角色关联 (DELETE body)
- POST   /api/relation/role-permissions 分配权限
- DELETE /api/relation/role-permissions 删除角色权限关联

### 4.5 操作日志
- GET    /api/logs                    查询操作日志 (支持分页、过滤)

> **注**：所有接口均基于 JWT 会话中间件，并校验租户隔离与权限。

## 五、未完成与后续计划
1. 单元与集成测试覆盖各 Handler 与 Service
2. 前端 E2E 测试（用户/角色/权限流程）
3. Docker 化部署与 CI/CD 脚本
4. 设计文档详尽 ER 图与 API 文档自动生成

---
## 六、已实现功能

### 1. 多租户与定制化
- 租户注册/登录/信息维护：已完成，支持独立域名、Logo、主题色、欢迎语等自定义，自动创建超级管理员
- 租户定制化内容存储：已实现 `tenant_customizations` 表（JSONB）存储定制信息

### 2. 用户与权限体系
- 用户注册/登录/找回密码：已完成，支持邮箱/手机号/用户名注册与登录
- JWT 鉴权与会话：已实现 `AuthMiddleware`（JWT + Redis 存储）
- 用户信息维护：已完成个人信息查询与修改接口
- 角色管理：已实现角色的创建/查询/更新/删除
- 权限管理：已实现权限的创建/查询接口
- 用户-角色、角色-权限分配：已实现对应关联接口
- 权限校验中间件：已实现 `RequirePermission`，超级管理员绕过
- 操作日志审计：已实现所有关键操作写入 `operation_logs` 表并提供查询接口

### 3. 前端功能
- 技术栈：Vue3 + Vite + TypeScript + Element Plus
- 统一 API 封装：封装 `{code,message,data}` 格式的响应，并在请求拦截器中解包
- Pinia 状态管理：实现用户信息与权限列表的获取与校验
- 动态侧边栏与路由守卫：根据权限动态渲染菜单并保护路由
- CRUD 页面与弹窗：用户管理、角色设置、租户定制、个人信息等页面及操作弹窗均已实现

---
## 七、未完成项
- 单元测试与集成测试：尚需为各 handler 和 service 编写测试用例
- GORM 迁移验证：完善 JSONB 索引、外键约束的验证与回滚策略
- 前端 E2E 测试：尚缺少仪表盘功能流的端到端测试
- CI/CD：Dockerfile 与 GitHub Actions 自动化迁移、构建与测试流水线待完善
- API 文档与 ER 图：需补充详细接口文档和数据库 ER 图
- 分页与搜索优化：后续可在列表接口中添加高级筛选与性能优化
- 错误处理与国际化：统一错误响应格式，并考虑多语言支持
- 日志管理查询前端：尚需实现操作日志查询的前端界面
- 角色默认标记管理 UI：需要在角色设置页面支持默认角色标记管理
- 默认角色自动分配：新建用户时若租户存在默认角色则自动分配

---
## 八、总结
第一阶段设计中核心功能已基本完成，后续聚焦测试覆盖、部署自动化和性能优化。
