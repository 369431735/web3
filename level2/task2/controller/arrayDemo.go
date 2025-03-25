package controller

import (
	"context"
	"fmt"
	"math/big"
	"net/http"
	"task2/config"
	bindings2 "task2/contracts/bindings"
	"task2/storage"
	"task2/utils"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
)

// ArrayDemoAddValueRequest 数组演示添加值请求
type ArrayDemoAddValueRequest struct {
	Value string `json:"value" binding:"required" example:"42"`
}

// ArrayDemoAddValueResponse 数组演示添加值响应
type ArrayDemoAddValueResponse struct {
	Address string `json:"address"`
	TxHash  string `json:"txHash"`
	Value   string `json:"value"`
}

// ArrayDemoGetValuesResponse 数组演示获取值响应
type ArrayDemoGetValuesResponse struct {
	Address string   `json:"address"`
	Values  []string `json:"values"`
	Count   int      `json:"count"`
}

// ----- ArrayDemo合约相关操作 -----

// ArrayDemoAddValue godoc
// @Summary      向ArrayDemo合约添加值
// @Description  向数组合约添加新的整数值
// @Tags         ArrayDemo合约操作
// @Accept       application/json
// @Produce      application/json
// @Param        request body ArrayDemoAddValueRequest true "添加值请求"
// @Success      200  {object}  ArrayDemoAddValueResponse
// @Failure      400  {object}  ErrorResponse
// @Failure      500  {object}  ErrorResponse
// @Router       /contracts/arraydemo/add-value [post]
func ArrayDemoAddValue(c *gin.Context) {
	var req ArrayDemoAddValueRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "请求参数错误: " + err.Error(),
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
	contractAddress, err := getContractAddress("ArrayDemo")
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "获取ArrayDemo合约地址失败: " + err.Error(),
		})
		return
	}

	// 转换数值
	bigValue := new(big.Int)
	_, ok := bigValue.SetString(req.Value, 10)
	if !ok {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "无效的数值格式",
		})
		return
	}

	// 创建合约实例
	instance, err := bindings2.NewArrayDemo(contractAddress, client)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "创建合约实例失败: " + err.Error(),
		})
		return
	}

	// 发送交易
	tx, err := instance.Put(auth, bigValue)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "发送交易失败: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, ArrayDemoAddValueResponse{
		Address: contractAddress.Hex(),
		TxHash:  tx.Hash().Hex(),
		Value:   req.Value,
	})
}

// ArrayDemoGetValues godoc
// @Summary      获取ArrayDemo合约的值
// @Description  获取ArrayDemo合约的所有值
// @Tags         ArrayDemo合约操作
// @Accept       application/json
// @Produce      application/json
// @Success      200  {object}  ArrayDemoGetValuesResponse
// @Failure      400  {object}  ErrorResponse
// @Failure      500  {object}  ErrorResponse
// @Router       /contracts/arraydemo/get-values [get]
func ArrayDemoGetValues(c *gin.Context) {
	// 获取合约地址
	contractAddress, err := getContractAddress("ArrayDemo")
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "获取ArrayDemo合约地址失败: " + err.Error(),
		})
		return
	}

	// 创建合约实例
	client, err := utils.GetEthClientHTTP()
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "初始化以太坊客户端失败: " + err.Error(),
		})
		return
	}

	instance, err := bindings2.NewArrayDemo(contractAddress, client)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "创建合约实例失败: " + err.Error(),
		})
		return
	}

	// 获取所有值
	values, err := instance.GetArray(&bind.CallOpts{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "获取值失败: " + err.Error(),
		})
		return
	}

	// 转换为字符串数组
	strValues := make([]string, len(values))
	for i, v := range values {
		strValues[i] = v.String()
	}

	c.JSON(http.StatusOK, ArrayDemoGetValuesResponse{
		Address: contractAddress.Hex(),
		Values:  strValues,
		Count:   len(strValues),
	})
}

// getTransactOpts 获取交易选项
func getTransactOpts(client *ethclient.Client) (*bind.TransactOpts, error) {
	network := config.GetCurrentNetwork()
	if network == nil {
		return nil, fmt.Errorf("获取当前网络配置失败")
	}

	// 获取私钥
	privateKeyHex := network.Accounts["default"].PrivateKey
	privateKey, err := utils.GetPrivateKey(privateKeyHex)
	if err != nil {
		return nil, fmt.Errorf("获取私钥失败: %v", err)
	}

	// 获取链ID
	chainID, err := client.ChainID(context.Background())
	if err != nil {
		return nil, fmt.Errorf("获取链ID失败: %v", err)
	}

	// 创建交易选项
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		return nil, fmt.Errorf("创建交易选项失败: %v", err)
	}

	// 设置GasLimit和GasPrice
	auth.GasLimit = uint64(3000000)
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, fmt.Errorf("获取Gas价格失败: %v", err)
	}
	auth.GasPrice = gasPrice

	return auth, nil
}

// getContractAddress 获取合约地址
func getContractAddress(contractType string) (common.Address, error) {
	contractStorage := storage.GetInstance()
	addressStr, err := contractStorage.GetAddress(contractType)
	if err == nil {
		return common.HexToAddress(addressStr), nil
	}
	return common.Address{}, fmt.Errorf("获取合约地址失败: %v", err)
}
