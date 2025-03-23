package controller

import (
	"math/big"
	"net/http"
	bindings2 "task2/contracts/bindings"
	"task2/utils"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/gin-gonic/gin"
)

// ----- Lottery合约相关操作 -----

// LotteryEnter godoc
// @Summary      参与彩票合约抽奖
// @Description  用户通过支付一定的以太币参与抽奖，金额必须大于最低要求的金额
// @Tags         Lottery
// @Accept       application/json
// @Produce      application/json
// @Param        request body LotteryEnterRequest true "参与彩票抽奖请求参数"
// @Success      200  {object}  SuccessResponse
// @Failure      400  {object}  ErrorResponse
// @Failure      500  {object}  ErrorResponse
// @Router       /contracts/Lottery/enter [post]
func LotteryEnter(c *gin.Context) {
	var req LotteryEnterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":   http.StatusBadRequest,
			"errMsg": "请求参数验证失败: " + err.Error(),
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

	// 设置以太币金额
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
	contractAddress, err := getContractAddress("Lottery")
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":   http.StatusInternalServerError,
			"errMsg": "获取Lottery合约地址失败: " + err.Error(),
			"data":   nil,
		})
		return
	}

	// 创建合约实例
	instance, err := bindings2.NewLottery(contractAddress, client)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":   http.StatusInternalServerError,
			"errMsg": "创建合约实例失败: " + err.Error(),
			"data":   nil,
		})
		return
	}

	// 调用合约的Enter方法
	tx, err := instance.Enter(auth)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":   http.StatusInternalServerError,
			"errMsg": "调用Enter方法失败: " + err.Error(),
			"data":   nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": gin.H{
			"message":          "成功参与抽奖",
			"transaction_hash": tx.Hash().Hex(),
			"contract_address": contractAddress.Hex(),
			"eth_value":        req.EthValue,
		},
	})
}

// LotteryPickWinner godoc
// @Summary      选取彩票获奖者
// @Description  管理员选取彩票获奖者并发放奖金
// @Tags         Lottery
// @Accept       application/json
// @Produce      application/json
// @Success      200  {object}  SuccessResponse
// @Failure      500  {object}  ErrorResponse
// @Router       /contracts/Lottery/pick-winner [post]
func LotteryPickWinner(c *gin.Context) {
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
	contractAddress, err := getContractAddress("Lottery")
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":   http.StatusInternalServerError,
			"errMsg": "获取Lottery合约地址失败: " + err.Error(),
			"data":   nil,
		})
		return
	}

	// 创建合约实例
	instance, err := bindings2.NewLottery(contractAddress, client)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":   http.StatusInternalServerError,
			"errMsg": "创建合约实例失败: " + err.Error(),
			"data":   nil,
		})
		return
	}

	// 调用合约的PickWinner方法
	tx, err := instance.PickWinner(auth)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":   http.StatusInternalServerError,
			"errMsg": "调用PickWinner方法失败: " + err.Error(),
			"data":   nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": gin.H{
			"message":          "成功选出获奖者",
			"transaction_hash": tx.Hash().Hex(),
			"contract_address": contractAddress.Hex(),
		},
	})
}

// LotteryGetPlayers godoc
// @Summary      获取所有参与者地址
// @Description  获取所有参与抽奖的用户地址列表
// @Tags         Lottery
// @Accept       application/json
// @Produce      application/json
// @Success      200  {object}  SuccessResponse
// @Failure      500  {object}  ErrorResponse
// @Router       /contracts/Lottery/players [get]
func LotteryGetPlayers(c *gin.Context) {
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
	contractAddress, err := getContractAddress("Lottery")
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":   http.StatusInternalServerError,
			"errMsg": "获取Lottery合约地址失败: " + err.Error(),
			"data":   nil,
		})
		return
	}

	// 创建合约实例
	instance, err := bindings2.NewLottery(contractAddress, client)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":   http.StatusInternalServerError,
			"errMsg": "创建合约实例失败: " + err.Error(),
			"data":   nil,
		})
		return
	}

	// 调用合约的GetPlayers方法
	players, err := instance.GetPlayers(&bind.CallOpts{})
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":   http.StatusInternalServerError,
			"errMsg": "调用GetPlayers方法失败: " + err.Error(),
			"data":   nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": gin.H{
			"players":          players,
			"contract_address": contractAddress.Hex(),
			"count":            len(players),
		},
	})
}

// LotteryGetBalance godoc
// @Summary      获取彩票合约余额
// @Description  获取彩票合约当前的余额，即奖池金额
// @Tags         Lottery
// @Accept       application/json
// @Produce      application/json
// @Success      200  {object}  SuccessResponse
// @Failure      500  {object}  ErrorResponse
// @Router       /contracts/Lottery/balance [get]
func LotteryGetBalance(c *gin.Context) {
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
	contractAddress, err := getContractAddress("Lottery")
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":   http.StatusInternalServerError,
			"errMsg": "获取Lottery合约地址失败: " + err.Error(),
			"data":   nil,
		})
		return
	}

	// 获取合约余额
	balance, err := client.BalanceAt(c, contractAddress, nil)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":   http.StatusInternalServerError,
			"errMsg": "获取合约余额失败: " + err.Error(),
			"data":   nil,
		})
		return
	}

	// 将balance转换为wei、gwei和ether格式
	weiBalance := balance.String()

	// 将balance转换为gwei (1 gwei = 10^9 wei)
	gweiBalance := new(big.Float).Quo(
		new(big.Float).SetInt(balance),
		big.NewFloat(1e9),
	).Text('f', 9)

	// 将balance转换为ether (1 ether = 10^18 wei)
	etherBalance := new(big.Float).Quo(
		new(big.Float).SetInt(balance),
		big.NewFloat(1e18),
	).Text('f', 18)

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": gin.H{
			"balance": gin.H{
				"wei":   weiBalance,
				"gwei":  gweiBalance,
				"ether": etherBalance,
			},
			"contract_address": contractAddress.Hex(),
		},
	})
}

// LotteryEnterRequest 参与彩票抽奖请求参数
type LotteryEnterRequest struct {
	EthValue string `json:"eth_value" binding:"required"` // 参与彩票抽奖所支付的以太币值
}
