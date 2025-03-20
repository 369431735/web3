# Web3 区块链接口服务

该项目提供以太坊区块链相关的API服务，包括账户管理、交易处理、合约部署和事件订阅等功能。

## 项目结构

```
.
├── abi                  # 已合并到contracts目录（旧目录）
├── api                  # API路由定义
├── bindings             # 已合并到contracts目录（旧目录）
├── config               # 配置相关代码
├── contracts            # 智能合约Go绑定
├── controller           # API控制器
├── docs                 # Swagger文档
├── handler              # 请求处理函数
├── storage              # 数据存储
├── types                # 数据类型定义
├── utils                # 工具函数
├── go.mod               # Go模块文件
├── go.sum               # Go依赖校验
├── main.go              # 主程序入口
└── README.md            # 项目说明
```

## 功能特性

- 账户管理：创建钱包、检查余额等
- 区块查询：获取区块信息、订阅新区块
- 合约部署：支持多种以太坊智能合约
- 事件订阅：订阅合约事件
- 交易处理：创建和发送交易

## 支持的合约

目前系统支持以下智能合约：

1. **SimpleStorage** - 简单存储合约，用于存储和读取数据
2. **Lock** - 锁定合约，在指定时间后允许提取资金
3. **Shipping** - 物流跟踪合约，用于跟踪商品交付状态
4. **SimpleAuction** - 简单拍卖合约，允许竞价和揭示最高出价者
5. **ArrayDemo** - 数组演示合约，展示在智能合约中操作数组
6. **Ballot** - 投票合约，实现基于区块链的投票系统
7. **Lottery** - 彩票合约，实现基于区块链的抽奖系统
8. **Purchase** - 购买合约，实现基于区块链的商品购买

## API文档

启动服务后，可以通过 `/swagger/index.html` 访问Swagger文档。

### 主要API端点

#### 合约部署

- `POST /api/v1/contract/deploy/simplestorage` - 部署SimpleStorage合约
- `POST /api/v1/contract/deploy/lock` - 部署Lock合约
  - 请求体:
    ```json
    {
      "unlockTime": 1680000000
    }
    ```
- `POST /api/v1/contracts/deploy-all` - 一键部署所有支持的合约

#### 合约调用

- `POST /api/v1/contract/simplestorage/set` - 设置SimpleStorage合约的值
- `GET /api/v1/contract/simplestorage/get` - 获取SimpleStorage合约的值
- `POST /api/v1/contract/lock/withdraw` - 从Lock合约中提取资金

#### 事件订阅

- `POST /api/v1/events/subscribe/:contractType` - 订阅指定类型合约的事件
  - 路径参数 `contractType`: `simplestorage` 或 `lock`
  - 请求体:
    ```json
    {
      "address": "0x合约地址"
    }
    ```

## 运行项目

1. 确保已安装Go 1.16+
2. 配置本地以太坊节点（或使用Ganache等测试节点）
3. 编译运行程序

```bash
go mod tidy
go run main.go
```

## 配置说明

配置文件位于`config/config.json`，可以配置以下内容：

- 服务器端口和基础路径
- 以太坊节点URL
- 账户信息
- 日志级别和输出位置

## 贡献代码

欢迎提交Pull Request或Issue。请确保代码格式化并通过测试。

## 项目变更说明

### 2024年3月20日更新

1. **合约支持扩展**
   - 增加了对contracts/bindings目录下所有合约的支持
   - 新增支持的合约：Ballot、Lottery、Purchase、SimpleAuction、Shipping、ArrayDemo

2. **一键部署功能增强**
   - 修改了DeployAllContracts函数，使其能够一键部署所有8种合约
   - 增加了合约的中间层封装，提高部署函数的复用性

3. **存储系统更新**
   - 扩展了合约地址存储结构，支持存储新增合约的地址
   - 更新了合约地址加载逻辑，支持加载所有类型的合约

4. **API路由优化**
   - 增加了合约事件订阅路由
   - 优化了合约部署路由结构

5. **文档更新**
   - 完善了README文档，添加了所有支持的合约说明
   - 增加了API使用示例 