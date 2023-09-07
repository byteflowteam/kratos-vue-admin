# Kratos Vue admin v1.0

> `kratos vue admin` 简称 `KVA` 是后端基于 `Kratos 2.x + gorm + casbin`， 前端基于`vue3` 实现的`后台管理系统`，开源版本遵循 `Apache` 开源协议，企业和个人都可以根据协议自由安装使用。

## 特性

- 遵循 `RESTful API` 设计规范 & 基于接口的编程规范
- 基于 `Kratos 2.x` 框架（支持微服务架构）.
- 基于 `Casbin` 的 RBAC 访问控制模型 -- **权限控制可以细粒度到按钮 & 接口**
- 基于 `gorm` 的数据库存储
- 基于 `WIRE` 的依赖注入 -- 依赖注入本身的作用是解决了各个模块间层级依赖繁琐的初始化过程
- 基于 `Zap & Context` 实现了日志输出，通过结合 Context 实现了统一的 TraceID/UserID 等关键字段的输出(同时支持日志钩子写入到`Gorm`)
- 基于 `JWT` 的用户认证 -- 基于 JWT 的黑名单验证机制
- 基于 `Swaggo` 自动生成 `Swagger` 文档 -- 独立于接口的 mock 实现
- 基于 `net/http/httptest` 标准包实现了 API 的单元测试
- 基于 `go mod` 的依赖管理(国内源可使用：<https://goproxy.cn/>)

### 安装依赖工具

```shell
# 初始化

make init

# 生成全部代码
make all

# 下载依赖

go mod tidy

```

### 启动命令

```shell
kratos run
```

### 构建

```shell
go build -o kva
```

## 特别鸣谢

- `kratos` 微服务框架。
- `vue3` 使用该前端框架进行开发后台管理web 界面。
