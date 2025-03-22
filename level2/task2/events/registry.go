package events

import (
	"context"
	"fmt"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
	"strings"
	"task2/utils"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
)

var (
	// 全局事件处理器映射
	EventHandlers = make(map[common.Hash]EventHandler)
)

// EventHandler 定义事件处理器接口
type EventHandler interface {
	Handle(log ethTypes.Log)
	GetEventName() string
	GetEventSignature() common.Hash
}

// RegisterEventHandler 注册事件处理器
func RegisterEventHandler(handler EventHandler) {
	signature := handler.GetEventSignature()
	EventHandlers[signature] = handler
	utils.LogInfo("注册事件处理器", map[string]interface{}{
		"event":     handler.GetEventName(),
		"signature": signature.Hex(),
	})
}

func InitializeEventHandlersByAdress(filename, adress string) {
	// 扫描合约绑定目录
	client, _ := utils.GetEthClientHTTP()
	code, _ := client.CodeAt(context.Background(), common.HexToAddress(adress), nil)

	// 需要合约有验证过源码
	abiJSON, _ := abi.JSON(strings.NewReader(string(code)))
	// 为每个事件创建处理器
	for eventName := range abiJSON.Events {
		eventSignature := abiJSON.Events[eventName].ID
		handler := createEventHandler(filename, eventName, eventSignature, &abiJSON)
		RegisterEventHandler(handler)
	}
}

// createEventHandler 创建特定事件的处理器
func createEventHandler(contractName, eventName string, eventSignatur common.Hash, contractABI *abi.ABI) EventHandler {
	baseHandler := BaseEventHandler{
		ContractABI:   contractABI,
		EventName:     eventName,
		ContractName:  contractName,
		EventSignatur: eventSignatur,
	}

	return &DefaultEventHandler{BaseEventHandler: baseHandler}
}

type BaseEventHandler struct {
	ContractABI   *abi.ABI
	EventName     string
	ContractName  string
	EventSignatur common.Hash
}

// GetEventName 获取事件名称
func (h *BaseEventHandler) GetEventName() string {
	return h.EventName
}

// GetEventSignature 获取事件签名
func (h *BaseEventHandler) GetEventSignature() common.Hash {
	event, exists := h.ContractABI.Events[h.EventName]
	if !exists {
		return common.Hash{}
	}
	return event.ID
}

// DefaultEventHandler 默认事件处理器
type DefaultEventHandler struct {
	BaseEventHandler
}

func (h *DefaultEventHandler) Handle(log ethTypes.Log) {
	utils.LogInfo("收到事件", map[string]interface{}{
		"contract":    h.ContractName,
		"event":       h.EventName,
		"blockNumber": log.BlockNumber,
		"txHash":      log.TxHash.Hex(),
		"topics":      log.Topics,
		"data":        fmt.Sprintf("%x", log.Data),
	})
}
