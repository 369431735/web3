package controller

import (
	"context"
	"github.com/gin-gonic/gin"
	"log"
	"math/big"
	"net/http"
	"sync"
	bindings2 "task2/contracts/bindings"
	"task2/contracts/watch"
	"task2/storage"
	"task2/utils"
)

// SimpleAuctionBidRequest 竞拍请求
type SimpleAuctionBidRequest struct {
	BidAmount string `json:"bid_amount" binding:"required"`
}

var (
	auctionWatcher *watch.AuctionWatcher
	watcherOnce    sync.Once
	ctx            context.Context
	cancelCtx      context.CancelFunc
)

// SimpleAuctionBidResponse 竞拍响应结果
type SimpleAuctionBidResponse struct {
	Address   string `json:"address"`
	TxHash    string `json:"txHash"`
	BidAmount string `json:"bid_amount"`
}

// ----- SimpleAuction合约相关操作 -----

// SimpleAuctionBid godoc
// @Summary      调用SimpleAuction合约进行竞拍
// @Description  向拍卖合约发起竞拍请求
// @Tags         SimpleAuction合约操作
// @Accept       application/json
// @Produce      application/json
// @Param        request body SimpleAuctionBidRequest true "竞拍金额信息"
// @Success      200  {object}  SimpleAuctionBidResponse
// @Failure      400  {object}  ErrorResponse
// @Failure      500  {object}  ErrorResponse
// @Router       /contracts/SimpleAuction/bid [post]
func SimpleAuctionBid(c *gin.Context) {
	var req SimpleAuctionBidRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "参数验证失败: " + err.Error(),
		})
		return
	}

	// 初始化以太坊客户端
	client, err := utils.GetEthClientHTTP()
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "初始化以太坊客户端失败: " + err.Error(),
		})
		return
	}

	// 获取交易选项
	auth, err := getTransactOpts(client)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "获取交易选项失败: " + err.Error(),
		})
		return
	}

	// 获取合约地址
	contractAddress, err := getContractAddress("SimpleAuction")
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "获取SimpleAuction合约地址失败: " + err.Error(),
		})
		return
	}

	// 转换为big.Int并设置竞拍金额
	bidAmount := new(big.Int)
	bidAmount.SetString(req.BidAmount, 10)
	auth.Value = bidAmount

	// 创建合约实例
	instance, err := bindings2.NewSimpleAuction(contractAddress, client)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "创建合约实例失败: " + err.Error(),
		})
		return
	}

	// 调用合约的Bid方法
	tx, err := instance.Bid(auth)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "调用合约方法失败: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, SimpleAuctionBidResponse{
		Address:   contractAddress.Hex(),
		TxHash:    tx.Hash().Hex(),
		BidAmount: bidAmount.String(),
	})
}

// SimpleAuctionWithdraw godoc
// @Summary      调用SimpleAuction合约结束拍卖并提取资金
// @Description  结束竞拍并允许中标者支付款项，非中标者提取资金
// @Tags         SimpleAuction合约操作
// @Accept       application/json
// @Produce      application/json
// @Success      200  {object}  ContractResponse
// @Failure      400  {object}  ErrorResponse
// @Failure      500  {object}  ErrorResponse
// @Router       /contracts/SimpleAuction/withdraw [post]
func SimpleAuctionWithdraw(c *gin.Context) {
	// 初始化以太坊客户端
	client, err := utils.GetEthClientHTTP()
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "初始化以太坊客户端失败: " + err.Error(),
		})
		return
	}

	// 获取交易选项
	auth, err := getTransactOpts(client)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "获取交易选项失败: " + err.Error(),
		})
		return
	}

	// 获取合约地址
	contractAddress, err := getContractAddress("SimpleAuction")
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "获取SimpleAuction合约地址失败: " + err.Error(),
		})
		return
	}

	// 创建合约实例
	instance, err := bindings2.NewSimpleAuction(contractAddress, client)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "创建合约实例失败: " + err.Error(),
		})
		return
	}

	// 调用合约的Withdraw方法
	tx, err := instance.AuctionEnd(auth)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "调用Withdraw方法失败: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, ContractResponse{
		ContractType: "SimpleAuction",
		Address:      contractAddress.Hex(),
		TxHash:       tx.Hash().Hex(),
	})
}

// SimpleAuctionWithdraw godoc
// @Summary      监听HighestBidIncreased
// @Description  监听HighestBidIncreased
// @Tags         SimpleAuction合约操作
// @Accept       application/json
// @Produce      application/json
// @Failure      400  {object}  ErrorResponse
// @Failure      500  {object}  ErrorResponse
// @Router       /contracts/SimpleAuction/watchHighestBidIncreased [post]
func WatchHighestBidIncreased(c *gin.Context) {
	var err error

	// 使用 sync.Once 确保只初始化一次 AuctionWatcher
	watcherOnce.Do(func() {
		contractAddr, _ := storage.GetInstance().GetAddress("SimpleAuction")
		auctionWatcher, err = watch.NewAuctionWatcher(contractAddr)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"code":    http.StatusInternalServerError,
				"message": "创建监控实例失败: " + err.Error(),
			})
			return
		}

		// 启动后台监听
		go func() {
			if err := auctionWatcher.WatchHighestBidIncreased(); err != nil {
				log.Fatalf("监听失败: %v", err)
			}
		}()
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": "启动监听失败: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":         http.StatusOK,
		"message":      "监听已启动",
		"contractType": "SimpleAuction",
	})
}

// Shutdown 优雅关闭监听器和上下文
func Shutdown(c *gin.Context) {
	cancelCtx()
	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "监听器已关闭",
	})
}
