package controller

import (
	"math/big"
	"net/http"
	bindings2 "task2/contracts/bindings"
	"task2/utils"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/gin-gonic/gin"
)

// ----- Purchase合约相关操作 -----

// PurchaseAbort godoc
// @Summary      中止购买合约并返还资金
// @Description  卖家可以在未收到确认购买前取消合约，如果已经被确认购买则无法中止
// @Tags         Purchase
// @Accept       application/json
// @Produce      application/json
// @Success      200  {object}  SuccessResponse
// @Failure      500  {object}  ErrorResponse
// @Router       /contracts/Purchase/abort [post]
func PurchaseAbort(c *gin.Context) {
	// 初始化以太坊客户端
	client, err := utils.GetEthClientHTTP()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":   http.StatusInternalServerError,
			"errMsg": "初始化以太坊客户端失败: " + err.Error(),
			"data":   nil,
		})
		return
	}

	// 获取交易选项
	auth, err := getTransactOpts(client)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":   http.StatusInternalServerError,
			"errMsg": "获取交易选项失败: " + err.Error(),
			"data":   nil,
		})
		return
	}

	// 获取合约地址
	contractAddress, err := getContractAddress("Purchase")
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":   http.StatusInternalServerError,
			"errMsg": "获取Purchase合约地址失败: " + err.Error(),
			"data":   nil,
		})
		return
	}

	// 创建合约实例
	instance, err := bindings2.NewPurchase(contractAddress, client)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":   http.StatusInternalServerError,
			"errMsg": "创建合约实例失败: " + err.Error(),
			"data":   nil,
		})
		return
	}

	// 调用合约的Abort方法
	tx, err := instance.Abort(auth)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":   http.StatusInternalServerError,
			"errMsg": "调用Abort方法失败: " + err.Error(),
			"data":   nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": gin.H{
			"message":          "合约中止成功，资金已返还",
			"transaction_hash": tx.Hash().Hex(),
			"contract_address": contractAddress.Hex(),
		},
		"errMsg": "",
	})
}

// PurchaseConfirmPurchase godoc
// @Summary      确认购买合约
// @Description  买家可以确认购买合约，并支付相应的以太币
// @Tags         Purchase
// @Accept       application/json
// @Produce      application/json
// @Param        request body PurchaseConfirmRequest true "购买合约的请求参数"
// @Success      200  {object}  SuccessResponse
// @Failure      400  {object}  ErrorResponse
// @Failure      500  {object}  ErrorResponse
// @Router       /contracts/Purchase/confirm [post]
func PurchaseConfirmPurchase(c *gin.Context) {
	var req PurchaseConfirmRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":   http.StatusBadRequest,
			"errMsg": "请求参数错误: " + err.Error(),
			"data":   nil,
		})
		return
	}

	// 初始化以太坊客户端
	client, err := utils.GetEthClientHTTP()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":   http.StatusInternalServerError,
			"errMsg": "初始化以太坊客户端失败: " + err.Error(),
			"data":   nil,
		})
		return
	}

	// 获取交易选项
	auth, err := getTransactOpts(client)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":   http.StatusInternalServerError,
			"errMsg": "获取交易选项失败: " + err.Error(),
			"data":   nil,
		})
		return
	}

	// 设置交易金额
	amount := new(big.Int)
	_, ok := amount.SetString(req.EthValue, 10)
	if !ok {
		c.JSON(http.StatusOK, gin.H{
			"code":   http.StatusBadRequest,
			"errMsg": "以太币值格式错误",
			"data":   nil,
		})
		return
	}
	auth.Value = amount

	// 获取合约地址
	contractAddress, err := getContractAddress("Purchase")
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":   http.StatusInternalServerError,
			"errMsg": "获取Purchase合约地址失败: " + err.Error(),
			"data":   nil,
		})
		return
	}

	// 创建合约实例
	instance, err := bindings2.NewPurchase(contractAddress, client)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":   http.StatusInternalServerError,
			"errMsg": "创建合约实例失败: " + err.Error(),
			"data":   nil,
		})
		return
	}

	// 调用合约的ConfirmPurchase方法
	tx, err := instance.ConfirmPurchase(auth)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":   http.StatusInternalServerError,
			"errMsg": "调用ConfirmPurchase方法失败: " + err.Error(),
			"data":   nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": gin.H{
			"message":          "合约确认购买成功，交易已提交",
			"transaction_hash": tx.Hash().Hex(),
			"contract_address": contractAddress.Hex(),
			"eth_value":        req.EthValue,
		},
		"errMsg": "",
	})
}

// PurchaseConfirmReceived godoc
// @Summary      确认收到货物
// @Description  买家可以确认收到货物，并支付相应的以太币
// @Tags         Purchase
// @Accept       application/json
// @Produce      application/json
// @Success      200  {object}  SuccessResponse
// @Failure      500  {object}  ErrorResponse
// @Router       /contracts/Purchase/confirm-received [post]
func PurchaseConfirmReceived(c *gin.Context) {
	// 初始化以太坊客户端
	client, err := utils.GetEthClientHTTP()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":   http.StatusInternalServerError,
			"errMsg": "初始化以太坊客户端失败: " + err.Error(),
			"data":   nil,
		})
		return
	}

	// 获取交易选项
	auth, err := getTransactOpts(client)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":   http.StatusInternalServerError,
			"errMsg": "获取交易选项失败: " + err.Error(),
			"data":   nil,
		})
		return
	}

	// 获取合约地址
	contractAddress, err := getContractAddress("Purchase")
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":   http.StatusInternalServerError,
			"errMsg": "获取Purchase合约地址失败: " + err.Error(),
			"data":   nil,
		})
		return
	}

	// 创建合约实例
	instance, err := bindings2.NewPurchase(contractAddress, client)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":   http.StatusInternalServerError,
			"errMsg": "创建合约实例失败: " + err.Error(),
			"data":   nil,
		})
		return
	}

	// 调用合约的ConfirmReceived方法
	tx, err := instance.ConfirmReceived(auth)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":   http.StatusInternalServerError,
			"errMsg": "调用ConfirmReceived方法失败: " + err.Error(),
			"data":   nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": gin.H{
			"message":          "货物确认收到，资金已支付",
			"transaction_hash": tx.Hash().Hex(),
			"contract_address": contractAddress.Hex(),
		},
		"errMsg": "",
	})
}

// PurchaseGetValue godoc
// @Summary      获取合约的当前价值
// @Description  获取合约的当前价值，以wei、gwei和ether为单位
// @Tags         Purchase
// @Accept       application/json
// @Produce      application/json
// @Success      200  {object}  SuccessResponse
// @Failure      500  {object}  ErrorResponse
// @Router       /contracts/Purchase/value [get]
func PurchaseGetValue(c *gin.Context) {
	// 初始化以太坊客户端
	client, err := utils.GetEthClientHTTP()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":   http.StatusInternalServerError,
			"errMsg": "初始化以太坊客户端失败: " + err.Error(),
			"data":   nil,
		})
		return
	}

	// 获取合约地址
	contractAddress, err := getContractAddress("Purchase")
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":   http.StatusInternalServerError,
			"errMsg": "获取Purchase合约地址失败: " + err.Error(),
			"data":   nil,
		})
		return
	}

	// 创建合约实例
	instance, err := bindings2.NewPurchase(contractAddress, client)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":   http.StatusInternalServerError,
			"errMsg": "创建合约实例失败: " + err.Error(),
			"data":   nil,
		})
		return
	}

	// 调用合约的Value方法获取合约当前价值
	value, err := instance.Value(&bind.CallOpts{})
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":   http.StatusInternalServerError,
			"errMsg": "调用Value方法失败: " + err.Error(),
			"data":   nil,
		})
		return
	}

	// 将value转换为wei、gwei和ether格式
	weiValue := value.String()

	// 将value转换为gwei格式
	gweiValue := new(big.Float).Quo(
		new(big.Float).SetInt(value),
		big.NewFloat(1e9),
	).Text('f', 9)

	// 将value转换为ether格式
	etherValue := new(big.Float).Quo(
		new(big.Float).SetInt(value),
		big.NewFloat(1e18),
	).Text('f', 18)

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": gin.H{
			"value": gin.H{
				"wei":   weiValue,
				"gwei":  gweiValue,
				"ether": etherValue,
			},
			"contract_address": contractAddress.Hex(),
		},
		"errMsg": "",
	})
}

// PurchaseGetState godoc
// @Summary      获取合约的状态
// @Description  获取合约的状态，包括状态码和状态文本
// @Tags         Purchase
// @Accept       application/json
// @Produce      application/json
// @Success      200  {object}  SuccessResponse
// @Failure      500  {object}  ErrorResponse
// @Router       /contracts/Purchase/state [get]
func PurchaseGetState(c *gin.Context) {
	// 初始化以太坊客户端
	client, err := utils.GetEthClientHTTP()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":   http.StatusInternalServerError,
			"errMsg": "初始化以太坊客户端失败: " + err.Error(),
			"data":   nil,
		})
		return
	}

	// 获取合约地址
	contractAddress, err := getContractAddress("Purchase")
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":   http.StatusInternalServerError,
			"errMsg": "获取Purchase合约地址失败: " + err.Error(),
			"data":   nil,
		})
		return
	}

	// 创建合约实例
	instance, err := bindings2.NewPurchase(contractAddress, client)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":   http.StatusInternalServerError,
			"errMsg": "创建合约实例失败: " + err.Error(),
			"data":   nil,
		})
		return
	}

	// 调用合约的State方法获取合约状态
	state, err := instance.State(&bind.CallOpts{})
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":   http.StatusInternalServerError,
			"errMsg": "调用State方法失败: " + err.Error(),
			"data":   nil,
		})
		return
	}

	// 将状态码转换为状态文本
	var stateText string
	switch state {
	case 0:
		stateText = "已创建(Created)"
	case 1:
		stateText = "已锁定(Locked)"
	case 2:
		stateText = "已释放(Released)"
	case 3:
		stateText = "已失效(Inactive)"
	default:
		stateText = "未知状态"
	}

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": gin.H{
			"state": gin.H{
				"code": state,
				"text": stateText,
			},
			"contract_address": contractAddress.Hex(),
		},
		"errMsg": "",
	})
}

// PurchaseConfirmRequest 购买合约的请求参数
type PurchaseConfirmRequest struct {
	EthValue string `json:"eth_value" binding:"required"` // 购买合约的以太币值
}
