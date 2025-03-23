# Web3 区块链接口服务

该项目提供以太坊区块链相关的API服务，包括账户管理、交易处理、合约部署和事件订阅等功能。基于Go语言开发，集成了以太坊客户端和智能合约交互。

## 项目结构

```
.
├── api                  # API路由定义
│   ├── router.go        # 路由配置
│   └── handlers.go      # 处理器注册
├── config               # 配置相关代码
│   ├── config.yml       # 配置文件
│   └── config.go        # 配置加载与管理
├── contracts            # 智能合约相关
│   ├── bindings/        # 合约Go绑定
│   ├── deploy/          # 合约部署逻辑
│   └── README.md        # 合约说明
├── controller           # API控制器
│   ├── account.go       # 账户相关控制器
│   ├── block.go         # 区块相关控制器
│   ├── contract.go      # 合约部署控制器
│   ├── events.go        # 事件相关控制器
│   └── ...              # 其他控制器
├── data                 # 数据存储目录
│   └── contracts.json   # 合约地址存储
├── docs                 # Swagger文档
├── events               # 事件处理
│   └── registry.go      # 事件注册与处理
├── storage              # 数据存储
│   └── contract.go      # 合约地址存储
├── types                # 数据类型定义
├── utils                # 工具函数
│   ├── client.go        # 以太坊客户端
│   ├── wallet.go        # 钱包工具
│   └── ...              # 其他工具
├── .env                 # 环境变量
├── go.mod               # Go模块文件
├── go.sum               # Go依赖校验
├── main.go              # 主程序入口
├── cleanup.sh           # 清理脚本
└── init.sh              # 初始化脚本
```

## 核心功能

### 1. 账户管理
- 创建以太坊钱包（普通钱包、密钥库、HD钱包）
- 检查账户余额
- 获取账户交易历史
- 查询账户nonce和代码

### 2. 区块查询
- 获取最新区块信息
- 通过区块号或哈希查询区块
- 获取区块详细信息和区块高度

### 3. 交易处理
- 创建并发送交易
- 创建原始交易
- 查询交易详情

### 4. 合约管理
- 一键部署所有支持的合约
- 获取所有已部署合约地址
- 获取合约字节码

### 5. 合约调用
- 调用合约方法（view和非view方法）
- 支持多种合约的专用API
- 通用合约方法调用接口

### 6. 事件订阅
- 订阅合约事件
- 自动注册和处理事件

## 支持的合约

目前系统支持以下智能合约：

1. **SimpleStorage** - 简单存储合约，用于存储和读取数据
   - 方法：`set`, `get`

2. **Lock** - 锁定合约，在指定时间后允许提取资金
   - 方法：`withdraw`, `getUnlockTime`

3. **Shipping** - 物流跟踪合约，用于跟踪商品交付状态
   - 方法：`advanceState`, `getState`

4. **SimpleAuction** - 简单拍卖合约，允许竞价和揭示最高出价者
   - 方法：`bid`, `withdraw`

5. **ArrayDemo** - 数组演示合约，展示在智能合约中操作数组
   - 方法：数组操作相关方法

6. **Ballot** - 投票合约，实现基于区块链的投票系统
   - 方法：投票和提案相关方法

7. **Lottery** - 彩票合约，实现基于区块链的抽奖系统
   - 方法：购买彩票、开奖等

8. **Purchase** - 购买合约，实现基于区块链的商品购买
   - 方法：购买、确认收货等

## API文档

启动服务后，可以通过 `/swagger/index.html` 访问Swagger文档。

### 主要API端点

#### 区块相关
- `GET /block/latest` - 获取最新区块信息
- `GET /block/:number` - 通过区块号获取区块信息
- `GET /block/hash/:hash` - 通过区块哈希获取区块信息
- `GET /block/info` - 获取区块详细信息
- `GET /block/number` - 获取当前区块号

#### 账户相关
- `GET /accounts/:address/balance` - 获取账户余额
- `GET /accounts/:address/transactions` - 获取账户交易历史
- `GET /accounts/:address/nonce` - 获取账户nonce值
- `GET /accounts/:address/code` - 获取账户合约代码

#### 钱包相关
- `POST /account/wallet` - 创建新的以太坊钱包
- `POST /account/keystore` - 创建新的以太坊密钥库文件
- `POST /account/hdwallet` - 创建新的分层确定性钱包
- `GET /account/balance/:address` - 获取账户余额

#### 交易相关
- `POST /transaction/create` - 创建并发送交易
- `POST /transaction/raw` - 创建原始交易
- `GET /transaction/:hash` - 获取交易详情

#### 合约相关
- `POST /contracts/deploy-all` - 一键部署所有合约
- `GET /contracts/allAddresses` - 获取所有已部署合约地址
- `POST /contracts/bytecode` - 获取合约字节码

#### 合约方法调用
- `POST /contracts/SimpleStorage/set` - 设置SimpleStorage合约的值
- `GET /contracts/SimpleStorage/get` - 获取SimpleStorage合约的值
- `POST /contracts/Lock/withdraw` - 从Lock合约中提取资金
- `GET /contracts/Lock/unlockTime` - 获取Lock合约的解锁时间
- `POST /contracts/Shipping/advance-state` - 更新物流状态
- `GET /contracts/Shipping/get-state` - 获取当前物流状态
- `POST /contracts/SimpleAuction/bid` - 参与拍卖出价
- `POST /contracts/SimpleAuction/withdraw` - 提取拍卖资金
- `POST /contracts/ArrayDemo/add-value` - 向数组中添加值
- `GET /contracts/ArrayDemo/get-values` - 获取数组中的所有值
- `POST /contracts/Ballot/vote` - 对提案进行投票
- `GET /contracts/Ballot/winner` - 获取当前获胜的提案ID
- `GET /contracts/Ballot/winner-name` - 获取当前获胜提案的名称
- `POST /contracts/Lottery/enter` - 参与彩票
- `POST /contracts/Lottery/draw` - 抽取彩票获胜者
- `GET /contracts/Lottery/players` - 获取参与彩票的玩家列表
- `GET /contracts/Lottery/balance` - 获取彩票合约余额
- `POST /contracts/Purchase/abort` - 卖家中止购买
- `POST /contracts/Purchase/confirm` - 买家确认购买
- `POST /contracts/Purchase/confirm-received` - 买家确认收货
- `GET /contracts/Purchase/value` - 获取商品价值
- `GET /contracts/Purchase/state` - 获取购买状态

#### 事件订阅
- `GET /events/subscribe` - 通过WebSocket订阅合约事件

## 运行项目

1. 确保已安装Go 1.16+
2. 配置本地以太坊节点（或使用Ganache、Infura等）
3. 设置配置文件`config/config.yml`
4. 编译运行程序

```bash
# 安装依赖
go mod tidy

# 运行服务
go run main.go
```

## 配置说明

配置文件位于`config/config.yml`，包含以下内容：

```yaml
# 服务器配置
server:
  port: 8080
  mode: debug  # debug或release
  basePath: /api/v1

# 以太坊网络配置
networks:
  # 本地开发网络
  development:
    url: http://localhost:8545
    chainId: 1337
    accounts:
      default:
        address: "0x..."
        privateKey: "0x..."

  # 测试网络
  goerli:
    url: https://goerli.infura.io/v3/YOUR_PROJECT_ID
    chainId: 5
    accounts:
      default:
        address: "0x..."
        privateKey: "0x..."

# 当前使用的网络
currentNetwork: development

# 日志配置
log:
  level: info
  filename: web3.log
```

## 开发说明

### 添加新合约

1. 在`contracts/bindings`目录下添加合约的Go绑定
2. 在`contracts/deploy/deploy.go`中添加部署方法
3. 在`controller`目录下添加合约方法控制器
4. 在`api/router.go`中注册合约API路由

### 项目依赖

- [go-ethereum](https://github.com/ethereum/go-ethereum) - 以太坊Go客户端
- [gin](https://github.com/gin-gonic/gin) - HTTP Web框架
- [gin-swagger](https://github.com/swaggo/gin-swagger) - Swagger文档生成

## 贡献代码

欢迎提交Pull Request或Issue。请确保代码格式化并通过测试。