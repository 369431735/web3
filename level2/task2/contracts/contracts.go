package contracts

import (
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
)

// 创建一个模拟交易
func createMockTransaction() *types.Transaction {
	// 创建一个有效的接收地址
	toAddress := common.HexToAddress("0xabcdef1234567890abcdef1234567890abcdef12")

	// 使用一个固定的私钥用于签名
	privateKey, _ := crypto.HexToECDSA("0123456789abcdef0123456789abcdef0123456789abcdef0123456789abcdef")

	// 使用NewTransaction创建交易对象
	tx := types.NewTransaction(
		uint64(100),            // nonce
		toAddress,              // to
		big.NewInt(1000000),    // value - 0.001 ETH
		uint64(21000),          // gas limit
		big.NewInt(1000000000), // gas price - 1 Gwei
		nil,                    // data
	)

	// 创建一个签名的交易 - 这会确保tx.hash非空
	signer := types.NewEIP155Signer(big.NewInt(1)) // chainID = 1
	signedTx, _ := types.SignTx(tx, signer, privateKey)

	return signedTx
}

// SimpleStorage 合约接口
type SimpleStorage struct {
	SimpleStorageCaller     // Read-only binding to the contract
	SimpleStorageTransactor // Write-only binding to the contract
	SimpleStorageFilterer   // Log filterer for contract events
}

// SimpleStorageCaller 只读绑定
type SimpleStorageCaller struct {
	contract *bind.BoundContract
}

// SimpleStorageTransactor 只写绑定
type SimpleStorageTransactor struct {
	contract *bind.BoundContract
}

// SimpleStorageFilterer 事件过滤器
type SimpleStorageFilterer struct {
	contract *bind.BoundContract
}

// DeploySimpleStorage 部署 SimpleStorage 合约
func DeploySimpleStorage(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SimpleStorage, error) {
	// 模拟合约部署，返回虚拟交易
	return common.HexToAddress("0x1234567890123456789012345678901234567890"), createMockTransaction(), nil, nil
}

// Lock 合约接口
type Lock struct {
	LockCaller     // Read-only binding to the contract
	LockTransactor // Write-only binding to the contract
	LockFilterer   // Log filterer for contract events
}

// LockCaller 只读绑定
type LockCaller struct {
	contract *bind.BoundContract
}

// LockTransactor 只写绑定
type LockTransactor struct {
	contract *bind.BoundContract
}

// LockFilterer 事件过滤器
type LockFilterer struct {
	contract *bind.BoundContract
}

// DeployLock 部署 Lock 合约
func DeployLock(auth *bind.TransactOpts, backend bind.ContractBackend, unlockTime *big.Int) (common.Address, *types.Transaction, *Lock, error) {
	// 模拟合约部署，返回虚拟交易
	return common.HexToAddress("0x2345678901234567890123456789012345678901"), createMockTransaction(), nil, nil
}

// Shipping 合约接口
type Shipping struct {
	ShippingCaller     // Read-only binding to the contract
	ShippingTransactor // Write-only binding to the contract
	ShippingFilterer   // Log filterer for contract events
}

// ShippingCaller 只读绑定
type ShippingCaller struct {
	contract *bind.BoundContract
}

// ShippingTransactor 只写绑定
type ShippingTransactor struct {
	contract *bind.BoundContract
}

// ShippingFilterer 事件过滤器
type ShippingFilterer struct {
	contract *bind.BoundContract
}

// DeployShipping 部署 Shipping 合约
func DeployShipping(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Shipping, error) {
	// 模拟合约部署，返回虚拟交易
	return common.HexToAddress("0x3456789012345678901234567890123456789012"), createMockTransaction(), nil, nil
}

// SimpleAuction 合约接口
type SimpleAuction struct {
	SimpleAuctionCaller     // Read-only binding to the contract
	SimpleAuctionTransactor // Write-only binding to the contract
	SimpleAuctionFilterer   // Log filterer for contract events
}

// SimpleAuctionCaller 只读绑定
type SimpleAuctionCaller struct {
	contract *bind.BoundContract
}

// SimpleAuctionTransactor 只写绑定
type SimpleAuctionTransactor struct {
	contract *bind.BoundContract
}

// SimpleAuctionFilterer 事件过滤器
type SimpleAuctionFilterer struct {
	contract *bind.BoundContract
}

// DeploySimpleAuction 部署 SimpleAuction 合约
func DeploySimpleAuction(auth *bind.TransactOpts, backend bind.ContractBackend, biddingTime *big.Int, beneficiary common.Address) (common.Address, *types.Transaction, *SimpleAuction, error) {
	// 模拟合约部署，返回虚拟交易
	return common.HexToAddress("0x4567890123456789012345678901234567890123"), createMockTransaction(), nil, nil
}

// ArrayDemo 合约接口
type ArrayDemo struct {
	ArrayDemoCaller     // Read-only binding to the contract
	ArrayDemoTransactor // Write-only binding to the contract
	ArrayDemoFilterer   // Log filterer for contract events
}

// ArrayDemoCaller 只读绑定
type ArrayDemoCaller struct {
	contract *bind.BoundContract
}

// ArrayDemoTransactor 只写绑定
type ArrayDemoTransactor struct {
	contract *bind.BoundContract
}

// ArrayDemoFilterer 事件过滤器
type ArrayDemoFilterer struct {
	contract *bind.BoundContract
}

// DeployArrayDemo 部署 ArrayDemo 合约
func DeployArrayDemo(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ArrayDemo, error) {
	// 模拟合约部署，返回虚拟交易
	return common.HexToAddress("0x5678901234567890123456789012345678901234"), createMockTransaction(), nil, nil
}
