# Pledge Backend

这是一个基于Go语言开发的质押（Pledge）服务后端系统，主要用于管理区块链质押相关业务。项目分为API服务和定时任务两个主要部分。

## 项目结构

```
pledge-backend/
├── api/                # API服务模块
│   ├── common/         # 公共函数和常量
│   ├── controllers/    # 控制器层，处理具体业务逻辑
│   ├── middlewares/    # 中间件，如跨域处理等
│   ├── models/         # 数据模型定义
│   ├── routes/         # 路由配置
│   ├── services/       # 服务层，封装业务逻辑
│   ├── static/         # 静态资源文件
│   ├── validate/       # 请求验证
│   └── pledge_api.go   # API服务入口文件
├── config/             # 配置文件目录
│   ├── conf.go         # 配置结构定义
│   ├── configV21.toml  # v2.1版本配置
│   ├── configV22.toml  # v2.2版本配置
│   └── init.go         # 配置初始化
├── contract/           # 区块链智能合约相关
│   ├── abi/            # 合约ABI定义
│   └── bindings/       # Go绑定的合约代码
├── db/                 # 数据库相关
│   ├── init.go         # 数据库初始化
│   ├── mysql.go        # MySQL连接和操作
│   ├── pledge.sql      # 数据库结构SQL
│   └── redis.go        # Redis连接和操作
├── log/                # 日志文件目录
├── schedule/           # 定时任务模块
│   ├── common/         # 定时任务公共函数
│   ├── models/         # 定时任务数据模型
│   ├── services/       # 定时任务服务层
│   ├── tasks/          # 具体定时任务实现
│   └── pledge_task.go  # 定时任务入口文件
├── utils/              # 工具函数
│   ├── email.go        # 邮件相关功能
│   ├── file.go         # 文件操作
│   ├── functions.go    # 通用工具函数
│   ├── jwt_token.go    # JWT令牌处理
│   ├── map.go          # Map相关操作
│   ├── md5.go          # MD5加密
│   ├── strings.go      # 字符串处理
│   └── time_format.go  # 时间格式化
├── go.mod              # Go模块依赖
└── go.sum              # 依赖校验和
```

## 模块说明

### API服务 (api/)

API服务提供了区块链质押业务的HTTP接口，包括用户认证、质押管理、数据查询等功能。

- **controllers/**: 控制器层，处理HTTP请求并调用相应服务
- **middlewares/**: 中间件，包括跨域处理、身份验证等
- **models/**: 数据模型定义，包括请求/响应结构
- **routes/**: 路由配置，定义API路径与控制器的映射
- **services/**: 业务逻辑实现
- **static/**: 静态资源，如图片等
- **validate/**: 请求参数验证

### 定时任务 (schedule/)

定时任务模块负责执行周期性的后台任务，如区块链数据同步、状态更新等。

- **models/**: 定时任务相关的数据模型
- **services/**: 定时任务的服务实现
- **tasks/**: 具体定时任务的实现
- **pledge_task.go**: 定时任务服务的入口点

### 配置管理 (config/)

管理系统配置，支持不同版本的配置文件。

- **conf.go**: 定义配置结构
- **configV21.toml, configV22.toml**: 不同版本的配置文件
- **init.go**: 负责初始化和加载配置

### 数据库 (db/)

数据库相关的功能，包括MySQL和Redis连接管理。

- **mysql.go**: MySQL数据库连接和基本操作
- **redis.go**: Redis连接和数据操作
- **pledge.sql**: 数据库表结构定义

### 智能合约 (contract/)

与区块链智能合约交互的相关代码。

- **abi/**: 合约ABI定义
- **bindings/**: Go语言与智能合约交互的绑定代码

### 工具函数 (utils/)

提供各种通用工具函数。

- **email.go**: 电子邮件发送功能
- **file.go**: 文件操作工具
- **functions.go**: 通用功能函数
- **jwt_token.go**: JWT令牌生成和验证
- **map.go**: Map数据结构相关操作
- **md5.go**: MD5加密功能
- **strings.go**: 字符串处理函数
- **time_format.go**: 时间格式化工具

## 运行说明

### API服务

```bash
cd api
go run pledge_api.go
```

### 定时任务

```bash
cd schedule
go run pledge_task.go
```

## 系统服务

项目提供了系统服务配置文件，可以作为系统服务运行：

- **api/pledge-api.service**: API服务的systemd服务配置
- **schedule/pledge-task.service**: 定时任务的systemd服务配置

## 技术栈

- 语言: Go
- Web框架: Gin
- 数据库: MySQL, Redis
- 区块链交互: 以太坊/BSC智能合约
- 定时任务: gocron