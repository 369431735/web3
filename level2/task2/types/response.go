package types

// BlockResponse 区块信息响应
type BlockResponse struct {
	Number     string `json:"number"`
	Hash       string `json:"hash"`
	ParentHash string `json:"parent_hash"`
	Timestamp  uint64 `json:"timestamp"`
	TxCount    int    `json:"tx_count"`
}
