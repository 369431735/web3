package abi

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

// ContractEvent 合约事件结构
type ContractEvent struct {
	ContractName string
	EventName    string
	Address      common.Address
	Topics       []common.Hash
	Data         []byte
}

// SubscribeContractEvents 订阅指定合约的所有事件
func SubscribeContractEvents(client *ethclient.Client, contractName string, contractAddress common.Address, contractABI *abi.ABI) error {
	// 创建一个查询过滤器
	query := ethereum.FilterQuery{
		Addresses: []common.Address{contractAddress},
	}

	// 创建日志通道
	logs := make(chan types.Log)

	// 订阅事件
	sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		return fmt.Errorf("订阅事件失败: %v", err)
	}

	// 在后台处理事件
	go func() {
		for {
			select {
			case err := <-sub.Err():
				log.Printf("事件订阅错误: %v", err)
				return
			case vLog := <-logs:
				// 解析事件
				event := ContractEvent{
					ContractName: contractName,
					Address:      vLog.Address,
					Topics:       vLog.Topics,
					Data:         vLog.Data,
				}

				// 尝试匹配事件名称
				for name, eventABI := range contractABI.Events {
					// 检查事件签名是否匹配
					if eventABI.ID == vLog.Topics[0] {
						event.EventName = name
						break
					}
				}

				// 打印事件信息
				log.Printf("收到合约事件:\n合约: %s\n事件: %s\n地址: %s\n数据: 0x%x\n",
					event.ContractName,
					event.EventName,
					event.Address.Hex(),
					event.Data,
				)

				// 如果有事件名称，尝试解析参数
				if event.EventName != "" {
					data := make(map[string]interface{})

					err := contractABI.UnpackIntoMap(data, event.EventName, event.Data)
					if err != nil {
						log.Printf("解析事件参数失败: %v", err)
						continue
					}

					// 打印解析后的参数
					log.Printf("事件参数:\n")
					for key, value := range data {
						log.Printf("%s: %v\n", key, value)
					}
				}
			}
		}
	}()

	return nil
}

// SubscribeAllContracts 订阅所有已部署合约的事件
func SubscribeAllContracts(client *ethclient.Client, contracts map[string]common.Address) error {
	// 遍历所有合约
	for contractName, address := range contracts {
		// 获取合约 ABI
		var contractABI *abi.ABI

		switch strings.ToLower(contractName) {
		case "simplestorage":
			parsed, err := abi.JSON(strings.NewReader(SimpleStorageABI))
			if err != nil {
				return fmt.Errorf("解析 SimpleStorage ABI 失败: %v", err)
			}
			contractABI = &parsed
		case "lock":
			parsed, err := abi.JSON(strings.NewReader(LockABI))
			if err != nil {
				return fmt.Errorf("解析 Lock ABI 失败: %v", err)
			}
			contractABI = &parsed
		case "shipping":
			parsed, err := abi.JSON(strings.NewReader(ShippingABI))
			if err != nil {
				return fmt.Errorf("解析 Shipping ABI 失败: %v", err)
			}
			contractABI = &parsed
		case "simpleauction":
			parsed, err := abi.JSON(strings.NewReader(SimpleAuctionABI))
			if err != nil {
				return fmt.Errorf("解析 SimpleAuction ABI 失败: %v", err)
			}
			contractABI = &parsed
		case "arraydemo":
			parsed, err := abi.JSON(strings.NewReader(ArrayDemoABI))
			if err != nil {
				return fmt.Errorf("解析 ArrayDemo ABI 失败: %v", err)
			}
			contractABI = &parsed
		default:
			log.Printf("警告: 未知合约类型 %s，跳过订阅", contractName)
			continue
		}

		// 订阅合约事件
		if err := SubscribeContractEvents(client, contractName, address, contractABI); err != nil {
			return fmt.Errorf("订阅合约 %s 事件失败: %v", contractName, err)
		}

		log.Printf("成功订阅合约 %s 的事件 (地址: %s)", contractName, address.Hex())
	}

	return nil
}
