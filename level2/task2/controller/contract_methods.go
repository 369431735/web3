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

// 合约方法调用的通用请求
type ContractMethodRequest struct {
	ContractName string        `json:"contract_name" binding:"required" example:"SimpleStorage"`
	Method       string        `json:"method" binding:"required" example:"set"`
	Params       []interface{} `json:"params" example:"[123]"`
}

// SimpleStorage合约的set方法请求
type SimpleStorageSetRequest struct {
	Value string `json:"value" binding:"required" example:"42"`
}

// SimpleStorage合约的get方法响应
type SimpleStorageGetResponse struct {
	Value string `json:"value" example:"42"`
}

// SimpleAuction合约的bid方法请求
type SimpleAuctionBidRequest struct {
	BidAmount string `json:"bid_amount" binding:"required" example:"1000000000000000000"`
}

// ArrayDemo合约的添加值方法请求
type ArrayDemoAddValueRequest struct {
	Value string `json:"value" binding:"required" example:"42"`
}

// getTransactOpts 获取交易选项
func getTransactOpts(client *ethclient.Client) (*bind.TransactOpts, error) {
	network := config.GetCurrentNetwork()
	if network == nil {
		return nil, fmt.Errorf("未找到网络配置")
	}

	// 解析私钥
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
	if err == nil {
		auth.GasPrice = gasPrice
	}

	return auth, nil
}

// 获取合约地址，优先从文件存储中获取
func getContractAddress(contractType string) (common.Address, error) {
	// 首先从文件存储中获取
	contractStorage := storage.GetInstance()
	addressStr, err := contractStorage.GetAddress(contractType)
	if err == nil {
		return common.HexToAddress(addressStr), nil
	}
	return common.Address{}, fmt.Errorf("地址格式无效: %v", err)

}

// ----- SimpleStorage合约方法 -----

// SimpleStorageSet godoc
// @Summary      调用SimpleStorage合约的set方法
// @Description  设置SimpleStorage合约的存储值
// @Tags         合约方法
// @Accept       json
// @Produce      json
// @Param        request body SimpleStorageSetRequest true "设置请求参数"
// @Success      200  {object}  map[string]interface{}
// @Failure      400  {object}  ErrorResponse
// @Failure      500  {object}  ErrorResponse
// @Router       /contract/simplestorage/set [post]
func SimpleStorageSet(c *gin.Context) {
	var req SimpleStorageSetRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "参数错误: " + err.Error(),
		})
		return
	}

	// 初始化客户端
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

	// 从已部署合约中获取地址
	contractAddress, err := getContractAddress("SimpleStorage")
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "获取SimpleStorage合约地址失败: " + err.Error(),
		})
		return
	}

	// 解析值参数
	value := new(big.Int)
	value.SetString(req.Value, 10)

	// 创建合约实例
	instance, err := bindings2.NewSimpleStorage(contractAddress, client)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "创建合约实例失败: " + err.Error(),
		})
		return
	}

	// 调用set方法
	tx, err := instance.Set(auth, value)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "调用set方法失败: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":          "调用SimpleStorage.set方法成功",
		"transaction_hash": tx.Hash().Hex(),
		"contract_address": contractAddress.Hex(),
	})
}

// SimpleStorageGet godoc
// @Summary      调用SimpleStorage合约的get方法
// @Description  获取SimpleStorage合约的存储值
// @Tags         合约方法
// @Accept       json
// @Produce      json
// @Success      200  {object}  SimpleStorageGetResponse
// @Failure      400  {object}  ErrorResponse
// @Failure      500  {object}  ErrorResponse
// @Router       /contract/simplestorage/get [get]
func SimpleStorageGet(c *gin.Context) {
	// 初始化客户端
	client, err := utils.GetEthClientHTTP()
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "初始化以太坊客户端失败: " + err.Error(),
		})
		return
	}

	// 从已部署合约中获取地址
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

	// 调用get方法
	result, err := instance.Get(&bind.CallOpts{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "调用get方法失败: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, SimpleStorageGetResponse{
		Value: result.String(),
	})
}

// ----- Lock合约方法 -----

// LockWithdraw godoc
// @Summary      调用Lock合约的withdraw方法
// @Description  从Lock合约中提取资金
// @Tags         合约方法
// @Accept       json
// @Produce      json
// @Param        request body LockWithdrawRequest true "提取请求参数"
// @Success      200  {object}  map[string]interface{}
// @Failure      400  {object}  ErrorResponse
// @Failure      500  {object}  ErrorResponse
// @Router       /contract/lock/withdraw [post]
func LockWithdraw(c *gin.Context) {
	// 初始化客户端
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

	// 从已部署合约中获取地址
	contractAddress, err := getContractAddress("Lock")
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "获取Lock合约地址失败: " + err.Error(),
		})
		return
	}

	// 创建合约实例
	instance, err := bindings2.NewLock(contractAddress, client)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "创建合约实例失败: " + err.Error(),
		})
		return
	}

	// 调用withdraw方法
	tx, err := instance.Withdraw(auth)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "调用withdraw方法失败: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":          "调用Lock.withdraw方法成功",
		"transaction_hash": tx.Hash().Hex(),
		"contract_address": contractAddress.Hex(),
	})
}

// ----- SimpleAuction合约方法 -----

// SimpleAuctionBid godoc
// @Summary      调用SimpleAuction合约的bid方法
// @Description  参与拍卖出价
// @Tags         合约方法
// @Accept       json
// @Produce      json
// @Param        request body SimpleAuctionBidRequest true "出价请求参数"
// @Success      200  {object}  map[string]interface{}
// @Failure      400  {object}  ErrorResponse
// @Failure      500  {object}  ErrorResponse
// @Router       /contract/simpleauction/bid [post]
func SimpleAuctionBid(c *gin.Context) {
	var req SimpleAuctionBidRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "参数错误: " + err.Error(),
		})
		return
	}

	// 初始化客户端
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

	// 从已部署合约中获取地址
	contractAddress, err := getContractAddress("SimpleAuction")
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "获取SimpleAuction合约地址失败: " + err.Error(),
		})
		return
	}

	// 解析出价金额
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

	// 调用bid方法
	tx, err := instance.Bid(auth)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "调用bid方法失败: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":          "调用SimpleAuction.bid方法成功",
		"transaction_hash": tx.Hash().Hex(),
		"contract_address": contractAddress.Hex(),
		"bid_amount":       bidAmount.String(),
	})
}

// SimpleAuctionWithdraw godoc
// @Summary      调用SimpleAuction合约的withdraw方法
// @Description  从拍卖合约中提取资金
// @Tags         合约方法
// @Accept       json
// @Produce      json
// @Success      200  {object}  map[string]interface{}
// @Failure      400  {object}  ErrorResponse
// @Failure      500  {object}  ErrorResponse
// @Router       /contract/simpleauction/withdraw [post]
func SimpleAuctionWithdraw(c *gin.Context) {
	// 初始化客户端
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

	// 从已部署合约中获取地址
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

	// 调用withdraw方法
	tx, err := instance.AuctionEnd(auth)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "调用withdraw方法失败: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":          "调用SimpleAuction.withdraw方法成功",
		"transaction_hash": tx.Hash().Hex(),
		"contract_address": contractAddress.Hex(),
	})
}

// ----- Shipping合约方法 -----

// ShippingAdvanceState godoc
// @Summary      调用Shipping合约的Shipped方法
// @Description  更新Shipping合约的运输状态
// @Tags         合约方法
// @Accept       json
// @Produce      json
// @Success      200  {object}  map[string]interface{}
// @Failure      400  {object}  ErrorResponse
// @Failure      500  {object}  ErrorResponse
// @Router       /contract/shipping/advance-state [post]
func ShippingAdvanceState(c *gin.Context) {
	// 初始化客户端
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

	// 从已部署合约中获取地址
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

	// 调用Shipped方法
	tx, err := instance.Shipped(auth)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "调用Shipped方法失败: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":          "调用Shipping.Shipped方法成功",
		"transaction_hash": tx.Hash().Hex(),
		"contract_address": contractAddress.Hex(),
	})
}

// ShippingGetState godoc
// @Summary      调用Shipping合约的Status方法
// @Description  获取Shipping合约的运输状态
// @Tags         合约方法
// @Accept       json
// @Produce      json
// @Success      200  {object}  map[string]interface{}
// @Failure      400  {object}  ErrorResponse
// @Failure      500  {object}  ErrorResponse
// @Router       /contract/shipping/get-state [get]
func ShippingGetState(c *gin.Context) {
	// 初始化客户端
	client, err := utils.GetEthClientHTTP()
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "初始化以太坊客户端失败: " + err.Error(),
		})
		return
	}

	// 从已部署合约中获取地址
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

	// 调用Status方法
	status, err := instance.Status(&bind.CallOpts{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "调用Status方法失败: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":           status,
		"contract_address": contractAddress.Hex(),
	})
}

// ----- ArrayDemo合约方法 -----

// ArrayDemoAddValue godoc
// @Summary      调用ArrayDemo合约的put方法
// @Description  添加一个值到数组中
// @Tags         合约方法
// @Accept       json
// @Produce      json
// @Param        request body ArrayDemoAddValueRequest true "请求参数"
// @Success      200  {object}  map[string]interface{}
// @Failure      400  {object}  ErrorResponse
// @Failure      500  {object}  ErrorResponse
// @Router       /contract/arraydemo/add-value [post]
func ArrayDemoAddValue(c *gin.Context) {
	var req ArrayDemoAddValueRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "参数错误: " + err.Error(),
		})
		return
	}

	// 初始化客户端
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

	// 从已部署合约中获取地址
	contractAddress, err := getContractAddress("ArrayDemo")
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "获取ArrayDemo合约地址失败: " + err.Error(),
		})
		return
	}

	// 解析值
	bigValue := new(big.Int)
	_, ok := bigValue.SetString(req.Value, 10)
	if !ok {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "无效的数值",
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

	// 调用Put方法
	tx, err := instance.Put(auth, bigValue)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "调用Put方法失败: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":          "调用ArrayDemo.Put方法成功",
		"transaction_hash": tx.Hash().Hex(),
		"contract_address": contractAddress.Hex(),
		"value":            req.Value,
	})
}

// ArrayDemoGetValues godoc
// @Summary      调用ArrayDemo合约的getArray方法
// @Description  获取数组中的所有值
// @Tags         合约方法
// @Accept       json
// @Produce      json
// @Success      200  {object}  map[string]interface{}
// @Failure      400  {object}  ErrorResponse
// @Failure      500  {object}  ErrorResponse
// @Router       /contract/arraydemo/get-values [get]
func ArrayDemoGetValues(c *gin.Context) {
	// 初始化客户端
	client, err := utils.GetEthClientHTTP()
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "初始化以太坊客户端失败: " + err.Error(),
		})
		return
	}

	// 从已部署合约中获取地址
	contractAddress, err := getContractAddress("ArrayDemo")
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "获取ArrayDemo合约地址失败: " + err.Error(),
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

	// 调用getArray方法
	values, err := instance.GetArray(&bind.CallOpts{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "调用getArray方法失败: " + err.Error(),
		})
		return
	}

	// 将大整数转换为字符串，方便JSON化
	strValues := make([]string, len(values))
	for i, v := range values {
		strValues[i] = v.String()
	}

	c.JSON(http.StatusOK, gin.H{
		"values":           strValues,
		"count":            len(strValues),
		"contract_address": contractAddress.Hex(),
	})
}
