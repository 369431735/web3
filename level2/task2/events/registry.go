package events

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"task2/utils"

	ethTypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/goccy/go-json"

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

// 匹配 Hardhat 编译结果的完整结构
type ContractArtifact struct {
	ABI      []abiEntry `json:"abi"`
	Bytecode string     `json:"bytecode"`
}

// 定义 ABI 条目结构
type abiEntry struct {
	Type      string        `json:"type"`
	Name      string        `json:"name,omitempty"`
	Inputs    []abiArgument `json:"inputs,omitempty"`
	Anonymous bool          `json:"anonymous,omitempty"`
}

// 定义 ABI 参数结构
type abiArgument struct {
	Name         string `json:"name"`
	Type         string `json:"type"`
	Indexed      bool   `json:"indexed,omitempty"`
	InternalType string `json:"internalType,omitempty"`
}

func InitializeEventHandlersByAdress(filename string) {
	// 从编译后的文件中读取ABI
	dirPath := filepath.Join("contracts", "compile", filename+".sol")
	jsonPath := filepath.Join(dirPath, filename+".json")

	data, err := ioutil.ReadFile(jsonPath)
	if err != nil {
		log.Fatalf("文件读取失败: %v", err)
	}

	// 2. 解析为结构体
	var artifact ContractArtifact
	if err := json.Unmarshal(data, &artifact); err != nil {
		log.Fatalf("JSON 解析失败: %v", err)
	}

	// 3. 重新序列化 ABI 部分
	abiData, err := json.Marshal(artifact.ABI)
	if err != nil {
		log.Fatalf("ABI 序列化失败: %v", err)
	}

	// 4. 解析为 ABI 对象
	abiJSON, err := abi.JSON(bytes.NewReader(abiData))
	if err != nil {
		log.Fatalf("ABI 解析失败: %v", err)
	}

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
