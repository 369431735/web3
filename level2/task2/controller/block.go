package controller

import (
	"math/big"
	"net/http"
	"task2/abi"
	"task2/config"
	"task2/types"
	"task2/utils"

	// 避免循环导入，使用其他方式获取事件订阅功能
	"github.com/ethereum/go-ethereum/common"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/gin-gonic/gin"
)

// BlockInfo 区块信息
type BlockInfo struct {
	Number     uint64 `json:"number" example:"123456"`
	Hash       string `json:"hash" example:"0x123..."`
	Timestamp  uint64 `json:"timestamp" example:"1634567890"`
	ParentHash string `json:"parent_hash" example:"0x456..."`
}

// TransactionRequest 交易请求参数
type TransactionRequest struct {
	From   string   `json:"from" binding:"required" example:"0x123..."`
	To     string   `json:"to" binding:"required" example:"0x456..."`
	Amount *big.Int `json:"amount" binding:"required" example:"1000000000000000000"`
}

// SubscribeBlock godoc
// @Summary      订阅新区块
// @Description  订阅以太坊网络的新区块
// @Tags         区块
// @Accept       json
// @Produce      json
// @Success      200  {object}  map[string]string
// @Failure      500  {object}  ErrorResponse
// @Router       /block/subscribe [post]
func SubscribeBlock(c *gin.Context) {
	// 初始化以太坊客户端
	client, err := utils.InitClient()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "初始化以太坊客户端失败: " + err.Error(),
		})
		return
	}

	if err := utils.SubscribeNewBlock(client); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "区块订阅成功",
	})
}

// GetBlockInfo godoc
// @Summary      获取区块信息
// @Description  获取指定区块的详细信息
// @Tags         区块
// @Accept       json
// @Produce      json
// @Param        number  query     string  true  "区块号"
// @Success      200     {object}  BlockResponse
// @Failure      500     {object}  ErrorResponse
// @Router       /block/info [get]
func GetBlockInfo(c *gin.Context) {
	number := c.Query("number")
	if number == "" {
		c.JSON(http.StatusBadRequest, types.ErrorResponse{Code: http.StatusBadRequest, Message: "区块号不能为空"})
		return
	}

	client, err := utils.InitClient()
	if err != nil {
		c.JSON(http.StatusInternalServerError, types.ErrorResponse{Code: http.StatusInternalServerError, Message: "初始化以太坊客户端失败: " + err.Error()})
		return
	}

	block, err := client.BlockByNumber(c.Request.Context(), nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, types.ErrorResponse{Code: http.StatusInternalServerError, Message: "获取区块信息失败: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, types.BlockResponse{
		Number:     block.Number().String(),
		Hash:       block.Hash().Hex(),
		ParentHash: block.ParentHash().Hex(),
		Timestamp:  block.Time(),
		TxCount:    len(block.Transactions()),
	})
}

// GetLatestBlock godoc
// @Summary      获取最新区块
// @Description  获取最新的区块信息
// @Tags         区块
// @Accept       json
// @Produce      json
// @Success      200  {object}  BlockResponse
// @Failure      500  {object}  ErrorResponse
// @Router       /block/latest [get]
func GetLatestBlock(c *gin.Context) {
	client, err := utils.InitClient()
	if err != nil {
		c.JSON(http.StatusInternalServerError, types.ErrorResponse{Code: http.StatusInternalServerError, Message: "初始化以太坊客户端失败: " + err.Error()})
		return
	}

	block, err := client.BlockByNumber(c.Request.Context(), nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, types.ErrorResponse{Code: http.StatusInternalServerError, Message: "获取最新区块失败: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, types.BlockResponse{
		Number:     block.Number().String(),
		Hash:       block.Hash().Hex(),
		ParentHash: block.ParentHash().Hex(),
		Timestamp:  block.Time(),
		TxCount:    len(block.Transactions()),
	})
}

// SubscribeBlocks godoc
// @Summary      订阅区块
// @Description  订阅新区块事件
// @Tags         区块
// @Accept       json
// @Produce      json
// @Success      200  {object}  BlockResponse
// @Failure      500  {object}  ErrorResponse
// @Router       /block/subscribe [get]
func SubscribeBlocks(c *gin.Context) {
	client, err := utils.InitClient()
	if err != nil {
		c.JSON(http.StatusInternalServerError, types.ErrorResponse{Code: http.StatusInternalServerError, Message: "初始化以太坊客户端失败: " + err.Error()})
		return
	}

	// 订阅新区块事件
	headers := make(chan *ethTypes.Header)
	sub, err := client.SubscribeNewHead(c.Request.Context(), headers)
	if err != nil {
		c.JSON(http.StatusInternalServerError, types.ErrorResponse{Code: http.StatusInternalServerError, Message: "订阅区块失败: " + err.Error()})
		return
	}

	// 处理新区块事件
	go func() {
		for {
			select {
			case err := <-sub.Err():
				utils.LogError("区块订阅错误", err)
			case header := <-headers:
				block, err := client.BlockByHash(c.Request.Context(), header.Hash())
				if err != nil {
					utils.LogError("获取区块信息失败", err)
					continue
				}

				utils.LogInfo("新区块", map[string]interface{}{
					"number":     block.Number().String(),
					"hash":       block.Hash().Hex(),
					"parentHash": block.ParentHash().Hex(),
					"timestamp":  block.Time(),
					"txCount":    len(block.Transactions()),
				})
			}
		}
	}()

	c.JSON(http.StatusOK, gin.H{
		"message": "区块订阅已启动",
	})
}

// CreateTransaction godoc
// @Summary      创建并发送交易
// @Description  创建一个新的交易并发送到以太坊网络
// @Tags         交易
// @Accept       json
// @Produce      json
// @Param        request body TransactionRequest true "交易请求参数"
// @Success      200  {object}  map[string]interface{} "返回交易哈希"
// @Failure      400  {object}  ErrorResponse "参数错误"
// @Failure      500  {object}  ErrorResponse "服务器内部错误"
// @Router       /transaction/create [post]
func CreateTransaction(c *gin.Context) {
	var req TransactionRequest
	// 绑定请求参数
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "参数错误: " + err.Error(),
		})
		return
	}

	// 调用工具函数发送交易
	txHash, err := utils.CreateAndSendTransaction(req.From, req.To, req.Amount)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "交易发送成功",
		"txHash":  txHash,
	})
}

// CreateRawTransaction godoc
// @Summary      创建原始交易
// @Description  创建一个未签名的原始交易
// @Tags         交易
// @Accept       json
// @Produce      json
// @Success      200  {object}  map[string]string
// @Failure      500  {object}  ErrorResponse
// @Router       /transaction/raw [post]
func CreateRawTransaction(c *gin.Context) {
	if err := utils.CreateRawTransaction(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "原始交易创建成功",
	})
}

// SubscribeContractEvents godoc
// @Summary      订阅合约事件
// @Description  订阅所有已部署合约的事件
// @Tags         事件
// @Accept       json
// @Produce      json
// @Success      200  {object}  map[string]string
// @Failure      500  {object}  ErrorResponse
// @Router       /events/subscribe [post]
func SubscribeContractEvents(c *gin.Context) {
	client, err := utils.InitClient()
	if err != nil {
		c.JSON(http.StatusInternalServerError, types.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "初始化以太坊客户端失败: " + err.Error(),
		})
		return
	}

	// 获取当前网络配置中的合约
	network := config.GetCurrentNetwork()
	if network == nil {
		c.JSON(http.StatusInternalServerError, types.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "获取网络配置失败",
		})
		return
	}

	// 准备合约地址映射
	contracts := make(map[string]common.Address)
	for name, contract := range network.Contracts {
		contracts[name] = common.HexToAddress(contract.Address)
	}

	// 调用abi包中的订阅所有合约事件方法
	if err := abi.SubscribeAllContracts(client, contracts); err != nil {
		c.JSON(http.StatusInternalServerError, types.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "订阅合约事件失败: " + err.Error(),
		})
		return
	}

	// 暂时返回成功消息
	c.JSON(http.StatusOK, gin.H{
		"message": "合约事件订阅已处理",
		"count":   len(contracts),
	})
}
