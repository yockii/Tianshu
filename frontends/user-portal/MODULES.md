# 租户/用户端（User Portal）模块梳理

面向外部租户及其用户，包括租户管理员和普通用户。支持租户自助注册、租户管理员自主管理本租户下的用户、设备、业务等。

## 主要模块

1. **租户注册与登录模块**
   - 租户自助注册
   - 租户登录/找回密码
   - 租户信息维护（Logo、主题、域名、欢迎语等）

2. **用户管理模块**
   - 用户注册/登录/找回密码
   - 用户信息维护
   - 用户分组与角色分配
   - 权限管理（基于租户自定义）

3. **设备管理与监控模块**
   - 设备列表/分组/详情
   - 设备状态监控（在线/离线/异常）
   - 设备绑定/解绑
   - 设备日志与告警

4. **实时数据与视频流模块**
   - 飞行器/机场实时数据展示
   - 视频流播放（支持多协议）
   - 历史数据回放

5. **航线与任务管理模块**
   - 航线规划与编辑（地图工具）
   - 航线上传/下发
   - 任务创建/调度/执行/监控
   - 任务历史与日志

6. **业务流程与表单自定义模块（预留）**
   - 业务流程建模
   - 表单自定义与数据采集

7. **个人中心与设置模块**
   - 个人信息/密码修改
   - 通知与消息
   - 主题切换

8. **多租户切换与主题隔离**
   - 支持不同租户独立主题、Logo、域名

---

如需进一步细化每个模块的页面与接口，可继续补充。
