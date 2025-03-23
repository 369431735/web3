package controller

import (
	"net/http"
	"task2/utils"

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
	// 待实现...
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
	// 待实现...
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
	// 待实现...
}
