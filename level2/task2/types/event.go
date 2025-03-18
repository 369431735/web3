package types

import (
	"github.com/ethereum/go-ethereum/common"
)

// ContractEvent 合约事件
type ContractEvent struct {
	ContractAddress common.Address `json:"contract_address"`
	EventName       string         `json:"event_name"`
	Data            interface{}    `json:"data"`
	BlockNumber     uint64         `json:"block_number"`
	TxHash          common.Hash    `json:"tx_hash"`
}
