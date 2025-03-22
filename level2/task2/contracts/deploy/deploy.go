package deploy

import (
	"math/big"
	"task2/contracts/bindings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

// 部署SimpleStorage合约
func DeploySimpleStorageFromBindings(auth *bind.TransactOpts, client *ethclient.Client) (common.Address, *types.Transaction, error) {
	address, tx, _, err := bindings.DeploySimpleStorage(auth, client)
	return address, tx, err
}

// 部署Lock合约
func DeployLockFromBindings(auth *bind.TransactOpts, client *ethclient.Client) (common.Address, *types.Transaction, error) {
	// 默认锁定时间为当前时间后24小时
	unlockTime := time.Now().Add(24 * time.Hour).Unix()
	address, tx, _, err := bindings.DeployLock(auth, client, big.NewInt(unlockTime))
	return address, tx, err
}

// 部署Shipping合约
func DeployShippingFromBindings(auth *bind.TransactOpts, client *ethclient.Client) (common.Address, *types.Transaction, error) {
	address, tx, _, err := bindings.DeployShipping(auth, client)
	return address, tx, err
}

// 部署SimpleAuction合约
func DeploySimpleAuctionFromBindings(auth *bind.TransactOpts, client *ethclient.Client) (common.Address, *types.Transaction, error) {
	// 拍卖时间为7天，受益人为交易发起人
	biddingTime := big.NewInt(7 * 24 * 60 * 60) // 7天（秒）
	beneficiary := auth.From
	address, tx, _, err := bindings.DeploySimpleAuction(auth, client, biddingTime, beneficiary)
	return address, tx, err
}

// 部署ArrayDemo合约
func DeployArrayDemoFromBindings(auth *bind.TransactOpts, client *ethclient.Client) (common.Address, *types.Transaction, error) {
	address, tx, _, err := bindings.DeployArrayDemo(auth, client)
	return address, tx, err
}

// 部署Ballot合约
func DeployBallotFromBindings(auth *bind.TransactOpts, client *ethclient.Client) (common.Address, *types.Transaction, error) {
	// 创建两个提案作为示例
	proposals := [][32]byte{
		toBytes32("提案1"),
		toBytes32("提案2"),
	}
	address, tx, _, err := bindings.DeployBallot(auth, client, proposals)
	return address, tx, err
}

// 部署Lottery合约
func DeployLotteryFromBindings(auth *bind.TransactOpts, client *ethclient.Client) (common.Address, *types.Transaction, error) {
	address, tx, _, err := bindings.DeployLottery(auth, client)
	return address, tx, err
}

// 部署Purchase合约
func DeployPurchaseFromBindings(auth *bind.TransactOpts, client *ethclient.Client) (common.Address, *types.Transaction, error) {
	// 设置价格为1 ETH
	value := big.NewInt(1000000000000000000) // 1 ETH (以wei为单位)
	auth.Value = value                       // 设置发送的ETH数量
	address, tx, _, err := bindings.DeployPurchase(auth, client)
	// 重置Value，避免影响其他部署
	auth.Value = big.NewInt(0)
	return address, tx, err
}

// 辅助函数：将字符串转换为[32]byte
func toBytes32(s string) [32]byte {
	var result [32]byte
	copy(result[:], s)
	return result
}
