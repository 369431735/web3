package controller

import (
	"math/big"
	"net/http"
	bindings2 "task2/contracts/bindings"
	"task2/utils"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/gin-gonic/gin"
)

// SimpleStorageSetRequest 设置存储值的请求
type SimpleStorageSetRequest struct {
	Value string `json:"value" binding:"required"`
}

// SimpleStorageGetValueResponse 获取存储值的响应
type SimpleStorageGetValueResponse struct {
	Value   string `json:"value"`
	Address string `json:"address"`
}

// ----- SimpleStorage合约相关操作 -----

// SimpleStorageSet godoc
// @Summary      设置SimpleStorage值
// @Description  调用SimpleStorage合约的set方法设置存储值
// @Tags         SimpleStorage
// @Accept       application/json
// @Produce      application/json
// @Param        value  body     string  true  "要存储的整数值"
// @Success      200    {object}  ContractResponse
// @Failure      400    {object}  ErrorResponse
// @Failure      500    {object}  ErrorResponse
// @Router       /contracts/SimpleStorage/set [post]
func SimpleStorageSet(c *gin.Context) {
	var req SimpleStorageSetRequest
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
	contractAddress, err := getContractAddress("SimpleStorage")
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "获取SimpleStorage合约地址失败: " + err.Error(),
		})
		return
	}

	// 创建合约实例
	instance, err := bindings2.NewSimpleStorage(contractAddress, client)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "创建合约实例失败: " + err.Error(),
		})
		return
	}

	// 转换为big.Int
	value := new(big.Int)
	value.SetString(req.Value, 10)

	// 调用合约的Set方法
	tx, err := instance.Set(auth, value)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "调用合约方法失败: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, ContractResponse{
		ContractType: "SimpleStorage",
		Address:      contractAddress.Hex(),
		TxHash:       tx.Hash().Hex(),
	})
}

// SimpleStorageGet godoc
// @Summary      获取SimpleStorage的值
// @Description  调用SimpleStorage合约的get方法获取当前存储的值
// @Tags         SimpleStorage
// @Accept       application/json
// @Produce      application/json
// @Success      200  {object}  SimpleStorageGetValueResponse
// @Failure      500  {object}  ErrorResponse
// @Router       /contracts/SimpleStorage/get [get]
func SimpleStorageGet(c *gin.Context) {
	// 初始化以太坊客户端
	client, err := utils.GetEthClientHTTP()
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "初始化以太坊客户端失败: " + err.Error(),
		})
		return
	}

	// 获取合约地址
	contractAddress, err := getContractAddress("SimpleStorage")
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "获取SimpleStorage合约地址失败: " + err.Error(),
		})
		return
	}

	// 创建合约实例
	instance, err := bindings2.NewSimpleStorage(contractAddress, client)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "创建合约实例失败: " + err.Error(),
		})
		return
	}

	// 调用合约的Get方法
	value, err := instance.Get(&bind.CallOpts{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "调用合约方法失败: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, SimpleStorageGetValueResponse{
		Value:   value.String(),
		Address: contractAddress.Hex(),
	})
}
