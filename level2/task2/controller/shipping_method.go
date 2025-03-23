package controller

import (
	"net/http"
	bindings2 "task2/contracts/bindings"
	"task2/utils"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/gin-gonic/gin"
)

// ShippingStateResponse 运输合约状态响应结构
type ShippingStateResponse struct {
	Status  string `json:"status"`
	Address string `json:"address"`
}

// ----- Shipping合约方法 -----

// ShippingAdvanceState godoc
// @Summary      执行Shipping合约的Shipped状态转换
// @Description  调用Shipping合约的Shipped方法，将状态推进到下一阶段
// @Tags         运输合约
// @Accept       application/json
// @Produce      application/json
// @Success      200  {object}  ContractResponse
// @Failure      400  {object}  ErrorResponse
// @Failure      500  {object}  ErrorResponse
// @Router       /contract/shipping/advance-state [post]
func ShippingAdvanceState(c *gin.Context) {
	// 创建以太坊客户端连接
	client, err := utils.GetEthClientHTTP()
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "创建以太坊客户端连接失败: " + err.Error(),
		})
		return
	}

	// 获取合约交易选项
	auth, err := getTransactOpts(client)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "获取合约交易选项失败: " + err.Error(),
		})
		return
	}

	// 获取合约地址
	contractAddress, err := getContractAddress("Shipping")
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "获取Shipping合约地址失败: " + err.Error(),
		})
		return
	}

	// 创建合约实例
	instance, err := bindings2.NewShipping(contractAddress, client)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "创建合约实例失败: " + err.Error(),
		})
		return
	}

	// 执行Shipped方法
	tx, err := instance.Shipped(auth)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "执行Shipped方法失败: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, ContractResponse{
		Address: contractAddress.Hex(),
		TxHash:  tx.Hash().Hex(),
	})
}

// ShippingGetState godoc
// @Summary      获取Shipping合约的Status状态
// @Description  获取Shipping合约的状态信息，包括当前阶段
// @Tags         运输合约
// @Accept       application/json
// @Produce      application/json
// @Success      200  {object}  ShippingStateResponse
// @Failure      400  {object}  ErrorResponse
// @Failure      500  {object}  ErrorResponse
// @Router       /contract/shipping/get-state [get]
func ShippingGetState(c *gin.Context) {
	// 创建以太坊客户端连接
	client, err := utils.GetEthClientHTTP()
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "创建以太坊客户端连接失败: " + err.Error(),
		})
		return
	}

	// 获取合约地址
	contractAddress, err := getContractAddress("Shipping")
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "获取Shipping合约地址失败: " + err.Error(),
		})
		return
	}

	// 创建合约实例
	instance, err := bindings2.NewShipping(contractAddress, client)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "创建合约实例失败: " + err.Error(),
		})
		return
	}

	// 获取Status状态
	status, err := instance.Status(&bind.CallOpts{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "获取Status状态失败: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, ShippingStateResponse{
		Status:  status,
		Address: contractAddress.Hex(),
	})
}
