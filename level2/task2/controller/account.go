package controller

import (
	"net/http"
	"task2/utils"

	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
)

// WalletResponse 钱包响应
type WalletResponse struct {
	Address string `json:"address" example:"0x742d35Cc6634C0532925a3b844Bc454e4438f44e"`
}

// GetAccountBalance godoc
// @Summary      获取账户余额
// @Description  获取指定以太坊地址的账户余额
// @Tags         账户
// @Accept       application/json
// @Produce      application/json
// @Param        address  path     string  true  "以太坊地址"
// @Success      200      {object}  AccountBalance
// @Failure      400      {object}  ErrorResponse
// @Failure      500      {object}  ErrorResponse
// @Router       /accounts/{address}/balance [get]
func GetAccountBalance(c *gin.Context) {
	address := c.Param("address")
	if address == "" {
		c.JSON(http.StatusBadRequest, ErrorResponse{Code: http.StatusBadRequest, Message: "地址不能为空"})
		return
	}

	account := common.HexToAddress(address)
	client, err := utils.GetEthClientHTTP()
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Code: http.StatusInternalServerError, Message: "初始化以太坊客户端失败: " + err.Error()})
		return
	}

	balance, err := client.BalanceAt(c.Request.Context(), account, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Code: http.StatusInternalServerError, Message: "获取余额失败: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, AccountBalance{
		Address: address,
		Balance: balance.String(),
	})
}

// SetAccountBalance godoc
// @Summary      设置账户余额
// @Description  设置指定以太坊地址的账户余额
// @Tags         账户
// @Accept       application/json
// @Produce      application/json
// @Param        address  path     string  true  "以太坊地址"
// @Param        request  body     AccountBalance  true  "余额请求"
// @Success      200      {object}  map[string]string
// @Failure      400      {object}  ErrorResponse
// @Failure      500      {object}  ErrorResponse
// @Router       /accounts/{address}/balance [put]
func SetAccountBalance(c *gin.Context) {
	address := c.Param("address")
	if address == "" {
		c.JSON(http.StatusBadRequest, ErrorResponse{Code: http.StatusBadRequest, Message: "地址不能为空"})
		return
	}

	var req AccountBalance
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Code: http.StatusBadRequest, Message: "参数错误: " + err.Error()})
		return
	}

	client, err := utils.GetEthClientHTTP()
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Code: http.StatusInternalServerError, Message: "初始化以太坊客户端失败: " + err.Error()})
		return
	}

	account := common.HexToAddress(address)
	balance, err := client.BalanceAt(c.Request.Context(), account, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Code: http.StatusInternalServerError, Message: "获取余额失败: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, map[string]string{
		"address": address,
		"balance": balance.String(),
	})
}

// CreateWallet godoc
// @Summary      创建钱包
// @Description  创建新的以太坊账户钱包
// @Tags         账户
// @Accept       application/json
// @Produce      application/json
// @Success      200  {object}  AccountResponse
// @Failure      500  {object}  ErrorResponse
// @Router       /accounts/wallet [post]
func CreateWallet(c *gin.Context) {
	account, err := utils.CreateAccount()
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Code: http.StatusInternalServerError, Message: "创建钱包失败: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, AccountResponse{
		Address: account.Address.Hex(),
	})
}

// CreateKeystore godoc
// @Summary      创建密钥库
// @Description  创建新的以太坊账户密钥库文件
// @Tags         账户
// @Accept       application/json
// @Produce      application/json
// @Success      200  {object}  AccountResponse
// @Failure      500  {object}  ErrorResponse
// @Router       /accounts/keystore [post]
func CreateKeystore(c *gin.Context) {
	account, err := utils.CreateKeystore()
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Code: http.StatusInternalServerError, Message: "创建密钥库失败: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, AccountResponse{
		Address: account.Address.Hex(),
	})
}

// CreateHDWallet godoc
// @Summary      创建 HD 钱包
// @Description  创建新的以太坊分层确定性钱包
// @Tags         账户
// @Accept       application/json
// @Produce      application/json
// @Success      200  {object}  map[string]string
// @Failure      500  {object}  ErrorResponse
// @Router       /accounts/hd-wallet [post]
func CreateHDWallet(c *gin.Context) {
	if err := utils.Chdwallet(); err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Code: http.StatusInternalServerError, Message: "创建 HD 钱包失败: " + err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "HD 钱包创建成功",
	})
}

// GetAccountTransactions godoc
// @Summary      获取账户交易列表
// @Description  获取指定以太坊地址的交易历史记录
// @Tags         账户
// @Accept       application/json
// @Produce      application/json
// @Param        address  path     string  true  "以太坊地址"
// @Success      200      {array}   TransactionResponse
// @Failure      400      {object}  ErrorResponse
// @Failure      500      {object}  ErrorResponse
// @Router       /accounts/{address}/transactions [get]
func GetAccountTransactions(c *gin.Context) {
	address := c.Param("address")
	if address == "" {
		c.JSON(http.StatusBadRequest, ErrorResponse{Code: http.StatusBadRequest, Message: "地址不能为空"})
		return
	}

	// 检查地址格式
	if !common.IsHexAddress(address) {
		c.JSON(http.StatusBadRequest, ErrorResponse{Code: http.StatusBadRequest, Message: "无效的以太坊地址"})
		return
	}

	client, err := utils.GetEthClientHTTP()
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Code: http.StatusInternalServerError, Message: "初始化以太坊客户端失败: " + err.Error()})
		return
	}

	// 获取最新区块号
	latestBlock, err := client.BlockByNumber(c.Request.Context(), nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Code: http.StatusInternalServerError, Message: "获取最新区块失败: " + err.Error()})
		return
	}

	// 仅查询最近的几个区块
	maxBlocks := 10
	startBlock := new(big.Int).Sub(latestBlock.Number(), big.NewInt(int64(maxBlocks)))
	if startBlock.Cmp(big.NewInt(0)) < 0 {
		startBlock = big.NewInt(0)
	}

	var transactions []TransactionResponse
	ethAddress := common.HexToAddress(address)

	// 遍历区块查找与地址相关的交易
	for blockNum := new(big.Int).Set(startBlock); blockNum.Cmp(latestBlock.Number()) <= 0; blockNum.Add(blockNum, big.NewInt(1)) {
		block, err := client.BlockByNumber(c.Request.Context(), blockNum)
		if err != nil {
			continue
		}

		for _, tx := range block.Transactions() {
			// 检查交易发送方
			from, err := client.TransactionSender(c.Request.Context(), tx, block.Hash(), uint(0))
			if err != nil {
				continue
			}

			// 检查是否与目标地址相关（发送方或接收方）
			if from == ethAddress || (tx.To() != nil && *tx.To() == ethAddress) {
				receipt, err := client.TransactionReceipt(c.Request.Context(), tx.Hash())
				var status string
				if err == nil {
					if receipt.Status == 1 {
						status = "成功"
					} else {
						status = "失败"
					}
				} else {
					status = "未知"
				}

				var to string
				if tx.To() != nil {
					to = tx.To().Hex()
				} else {
					to = "合约创建"
				}

				transactions = append(transactions, TransactionResponse{
					Hash:        tx.Hash().Hex(),
					From:        from.Hex(),
					To:          to,
					Value:       tx.Value().String(),
					Gas:         tx.Gas(),
					GasPrice:    tx.GasPrice().String(),
					BlockHash:   block.Hash().Hex(),
					BlockNumber: block.Number().String(),
					Timestamp:   block.Time(),
					Status:      status,
				})
			}
		}
	}

	c.JSON(http.StatusOK, transactions)
}

// GetAccountNonce godoc
// @Summary      获取账户Nonce值
// @Description  获取指定以太坊地址的账户当前Nonce值
// @Tags         账户
// @Accept       application/json
// @Produce      application/json
// @Param        address  path     string  true  "以太坊地址"
// @Success      200      {object}  map[string]uint64
// @Failure      400      {object}  ErrorResponse
// @Failure      500      {object}  ErrorResponse
// @Router       /accounts/{address}/nonce [get]
func GetAccountNonce(c *gin.Context) {
	address := c.Param("address")
	if address == "" {
		c.JSON(http.StatusBadRequest, ErrorResponse{Code: http.StatusBadRequest, Message: "地址不能为空"})
		return
	}

	// 检查地址格式
	if !common.IsHexAddress(address) {
		c.JSON(http.StatusBadRequest, ErrorResponse{Code: http.StatusBadRequest, Message: "无效的以太坊地址"})
		return
	}

	client, err := utils.GetEthClientHTTP()
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Code: http.StatusInternalServerError, Message: "初始化以太坊客户端失败: " + err.Error()})
		return
	}

	// 获取账户的nonce值
	nonce, err := client.NonceAt(c.Request.Context(), common.HexToAddress(address), nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Code: http.StatusInternalServerError, Message: "获取账户Nonce失败: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, map[string]uint64{"nonce": nonce})
}

// GetAccountCode godoc
// @Summary      获取账户合约代码
// @Description  获取指定以太坊地址的合约代码，如果地址是合约地址则返回其代码
// @Tags         账户
// @Accept       application/json
// @Produce      application/json
// @Param        address  path     string  true  "以太坊地址"
// @Success      200      {object}  map[string]string
// @Failure      400      {object}  ErrorResponse
// @Failure      500      {object}  ErrorResponse
// @Router       /accounts/{address}/code [get]
func GetAccountCode(c *gin.Context) {
	address := c.Param("address")
	if address == "" {
		c.JSON(http.StatusBadRequest, ErrorResponse{Code: http.StatusBadRequest, Message: "地址不能为空"})
		return
	}

	// 检查地址格式
	if !common.IsHexAddress(address) {
		c.JSON(http.StatusBadRequest, ErrorResponse{Code: http.StatusBadRequest, Message: "无效的以太坊地址"})
		return
	}

	client, err := utils.GetEthClientHTTP()
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Code: http.StatusInternalServerError, Message: "初始化以太坊客户端失败: " + err.Error()})
		return
	}

	// 获取合约代码
	code, err := client.CodeAt(c.Request.Context(), common.HexToAddress(address), nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Code: http.StatusInternalServerError, Message: "获取合约代码失败: " + err.Error()})
		return
	}

	var message string
	if len(code) == 0 {
		message = "该地址不是合约地址或合约代码为空"
	} else {
		message = "成功获取合约代码"
	}

	c.JSON(http.StatusOK, map[string]string{
		"address": address,
		"code":    common.Bytes2Hex(code),
		"message": message,
	})
}
