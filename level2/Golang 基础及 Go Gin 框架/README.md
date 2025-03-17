# Web3 Go 项目

这是一个使用 Go 语言开发的 Web3 项目，包含了以太坊智能合约的部署和交互功能。

## 项目结构

```
.
├── abi/                # 智能合约 ABI 和绑定文件
│   ├── bindings/      # 生成的合约绑定
│   └── contracts/     # 合约源文件
├── utils/             # 工具函数
├── main.go           # 主程序入口
├── go.mod            # Go 模块文件
└── README.md         # 项目说明文档
```

## 功能特性

- 以太坊客户端连接
- 智能合约部署
- 区块订阅
- 交易发送
- 钱包管理
- 地址转换
- 余额查询

## 环境要求

- Go 1.20 或更高版本
- 以太坊节点（如 Hardhat、Ganache 等）
- abigen 工具（用于生成合约绑定）

## 安装

1. 克隆项目
```bash
git clone <repository-url>
```

2. 安装依赖
```bash
go mod download
```

3. 生成合约绑定
```bash
cd abi
go run generate_bindings.go
```

## 使用

1. 启动本地以太坊节点（如 Hardhat）
```bash
npx hardhat node
```

2. 运行程序
```bash
go run main.go
```

## 配置

- 默认连接到 `http://localhost:8545`
- 使用 Hardhat 默认的第一个账户（私钥：`ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80`）

## 注意事项

- 确保以太坊节点已经启动
- 确保账户有足够的 ETH 用于部署合约
- 合约部署可能需要一些时间，请耐心等待 