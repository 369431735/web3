package controller

import (
	"math/big"
	"net/http"
	"strings"
	"task2/utils"

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
	From   string `json:"from" binding:"required" example:"0x123..."`
	To     string `json:"to" binding:"required" example:"0x456..."`
	Amount string `json:"amount" binding:"required" example:"0xde0b6b3a7640000"` // 十六进制字符串表示的wei金额，例如"0xde0b6b3a7640000"表示1 ETH
}

// SubscribeBlock godoc
// @Summary      启动区块订阅
// @Description  启动后台服务订阅新区块
// @Tags         区块
// @Accept       application/json
// @Produce      application/json
// @Success      200  {object}  map[string]string
// @Failure      500  {object}  ErrorResponse
// @Router       /block/subscribe [post]
func SubscribeBlock(c *gin.Context) {
	// 初始化以太坊客户端
	client, err := utils.GetEthClientHTTP()
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
		"message": "区块订阅启动成功",
	})
}

// GetBlockInfo godoc
// @Summary      获取区块信息
// @Description  根据区块编号获取区块信息
// @Tags         区块
// @Accept       application/json
// @Produce      application/json
// @Param        number  query     string  true  "区块编号"
// @Success      200     {object}  BlockResponse
// @Failure      500     {object}  ErrorResponse
// @Router       /block/info [get]
func GetBlockInfo(c *gin.Context) {
	number := c.Query("number")
	if number == "" {
		c.JSON(http.StatusBadRequest, ErrorResponse{Code: http.StatusBadRequest, Message: "请输入区块编号"})
		return
	}

	// 初始化以太坊客户端
	client, err := utils.GetEthClientHTTP()
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Code: http.StatusInternalServerError, Message: "初始化以太坊客户端失败: " + err.Error()})
		return
	}

	block, err := client.BlockByNumber(c.Request.Context(), nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Code: http.StatusInternalServerError, Message: "获取区块信息失败: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, BlockResponse{
		Number:     block.Number().String(),
		Hash:       block.Hash().Hex(),
		ParentHash: block.ParentHash().Hex(),
		Timestamp:  block.Time(),
		TxCount:    len(block.Transactions()),
	})
}

// GetLatestBlock godoc
// @Summary      获取最新区块
// @Description  获取最新区块的信息
// @Tags         区块
// @Accept       application/json
// @Produce      application/json
// @Success      200  {object}  BlockResponse
// @Failure      500  {object}  ErrorResponse
// @Router       /block/latest [get]
func GetLatestBlock(c *gin.Context) {
	// 初始化以太坊客户端
	client, err := utils.GetEthClientHTTP()
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Code: http.StatusInternalServerError, Message: "初始化以太坊客户端失败: " + err.Error()})
		return
	}

	block, err := client.BlockByNumber(c.Request.Context(), nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Code: http.StatusInternalServerError, Message: "获取最新区块失败: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, BlockResponse{
		Number:     block.Number().String(),
		Hash:       block.Hash().Hex(),
		ParentHash: block.ParentHash().Hex(),
		Timestamp:  block.Time(),
		TxCount:    len(block.Transactions()),
	})
}

// SubscribeBlocks godoc
// @Summary      订阅区块
// @Description  使用WebSocket订阅新区块
// @Tags         区块
// @Accept       application/json
// @Produce      application/json
// @Success      200  {object}  BlockResponse
// @Failure      500  {object}  ErrorResponse
// @Router       /block/subscribe [get]
func SubscribeBlocks(c *gin.Context) {
	// 初始化WebSocket客户端
	client, err := utils.GetEthClientWS()
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Code: http.StatusInternalServerError, Message: "初始化WebSocket客户端失败: " + err.Error()})
		return
	}

	// 订阅新区块
	headers := make(chan *ethTypes.Header)
	sub, err := client.SubscribeNewHead(c.Request.Context(), headers)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Code: http.StatusInternalServerError, Message: "订阅区块失败: " + err.Error()})
		return
	}

	// 处理新区块
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

				utils.LogInfo("获取到新区块", map[string]interface{}{
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
		"message": "区块订阅成功，请等待新区块",
	})
}

// CreateTransaction godoc
// @Summary      创建交易
// @Description  创建一笔交易，并发送给以太坊网络
// @Tags         交易
// @Accept       application/json
// @Produce      application/json
// @Param        request body TransactionRequest true "交易请求参数"
// @Success      200  {object}  map[string]interface{} "返回创建交易的结果"
// @Failure      400  {object}  ErrorResponse "参数错误"
// @Failure      500  {object}  ErrorResponse "创建交易失败"
// @Router       /transaction/create [post]
func CreateTransaction(c *gin.Context) {
	var req TransactionRequest
	// 解析请求参数
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "参数错误: " + err.Error(),
		})
		return
	}

	// 初始化以太坊客户端
	_, err := utils.GetEthClientHTTP()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	// 创建并发送交易
	amount := new(big.Int)
	// 如果Amount是十六进制字符串（以0x开头），则解析十六进制
	if strings.HasPrefix(req.Amount, "0x") {
		amount.SetString(req.Amount[2:], 16)
	} else {
		// 否则尝试解析为十进制
		amount.SetString(req.Amount, 10)
	}

	txHash, err := utils.CreateAndSendTransaction(req.From, req.To, amount)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "交易创建成功",
		"txHash":  txHash,
	})
}

// CreateRawTransaction godoc
// @Summary      创建原始交易
// @Description  创建一笔原始交易，并发送给以太坊网络
// @Tags         交易
// @Accept       application/json
// @Produce      application/json
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

// GetBlockByNumber godoc
// @Summary      获取区块
// @Description  根据区块编号获取区块信息
// @Tags         区块
// @Accept       application/json
// @Produce      application/json
// @Param        number  path     string  true  "区块编号"
// @Success      200     {object}  BlockResponse
// @Failure      400     {object}  ErrorResponse
// @Failure      500     {object}  ErrorResponse
// @Router       /block/{number} [get]
func GetBlockByNumber(c *gin.Context) {
	number := c.Param("number")
	if number == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "请输入区块编号"})
		return
	}

	client, err := utils.GetEthClientHTTP()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "初始化以太坊客户端失败: " + err.Error()})
		return
	}

	blockNumber := new(big.Int)
	blockNumber.SetString(number, 10)
	block, err := client.BlockByNumber(c.Request.Context(), blockNumber)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "获取区块信息失败: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, BlockResponse{
		Number:     block.Number().String(),
		Hash:       block.Hash().Hex(),
		ParentHash: block.ParentHash().Hex(),
		Timestamp:  block.Time(),
		TxCount:    len(block.Transactions()),
	})
}

// GetBlockByHash godoc
// @Summary      根据区块哈希获取区块信息
// @Description  根据区块哈希获取区块信息
// @Tags         区块
// @Accept       application/json
// @Produce      application/json
// @Param        hash  path     string  true  "区块哈希"
// @Success      200   {object}  BlockResponse
// @Failure      400   {object}  ErrorResponse
// @Failure      500   {object}  ErrorResponse
// @Router       /block/hash/{hash} [get]
func GetBlockByHash(c *gin.Context) {
	hash := c.Param("hash")
	if hash == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "请输入区块哈希"})
		return
	}

	client, err := utils.GetEthClientHTTP()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "初始化以太坊客户端失败: " + err.Error()})
		return
	}

	block, err := client.BlockByHash(c.Request.Context(), common.HexToHash(hash))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "获取区块信息失败: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, BlockResponse{
		Number:     block.Number().String(),
		Hash:       block.Hash().Hex(),
		ParentHash: block.ParentHash().Hex(),
		Timestamp:  block.Time(),
		TxCount:    len(block.Transactions()),
	})
}
