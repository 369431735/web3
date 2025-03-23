package controller

import (
	"math/big"
	"net/http"
	bindings2 "task2/contracts/bindings"
	"task2/utils"

	"github.com/gin-gonic/gin"
)

// SimpleAuctionBidRequest 竞拍请求
type SimpleAuctionBidRequest struct {
	BidAmount string `json:"bid_amount" binding:"required"`
}

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
// @Tags         合约操作
// @Accept       application/json
// @Produce      application/json
// @Param        request body SimpleAuctionBidRequest true "竞拍金额信息"
// @Success      200  {object}  SimpleAuctionBidResponse
// @Failure      400  {object}  ErrorResponse
// @Failure      500  {object}  ErrorResponse
// @Router       /contract/simpleauction/bid [post]
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
// @Tags         合约操作
// @Accept       application/json
// @Produce      application/json
// @Success      200  {object}  ContractResponse
// @Failure      400  {object}  ErrorResponse
// @Failure      500  {object}  ErrorResponse
// @Router       /contract/simpleauction/withdraw [post]
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
