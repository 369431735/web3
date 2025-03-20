# 合约绑定

本目录包含了以太坊智能合约的 Go 绑定代码。

## 目录结构说明

- `Lock.go` - Lock 合约的绑定代码
- `SimpleStorage.go` - SimpleStorage 合约的绑定代码
- `deploy.go` - 合约部署辅助函数

## 使用说明

### 部署合约

```go
// 部署 Lock 合约
address, tx, instance, err := contracts.DeployLock(auth, client, big.NewInt(unlockTime))

// 部署 SimpleStorage 合约
address, tx, instance, err := contracts.DeploySimpleStorage(auth, client)
```

### 订阅合约事件

```go
// 创建Lock合约实例
instance, err := contracts.NewLock(contractAddress, client)

// 订阅Withdrawal事件
sink := make(chan *contracts.LockWithdrawal)
sub, err := instance.LockFilterer.WatchWithdrawal(opts, sink)
```

### 前端订阅

通过API接口订阅合约事件：

```
GET /events/subscribe/Lock
GET /events/subscribe/SimpleStorage
```

## 注意事项

1. 在订阅事件前，确保使用 WebSocket 连接
2. 合约二进制代码已包含在绑定文件中
3. 所有合约绑定都已从 `contracts/bindings` 移动到 `contracts` 目录下 