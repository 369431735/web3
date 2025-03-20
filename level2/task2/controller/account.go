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
// @Description  获取指定账户的以太坊余额
// @Tags         账户
// @Accept       json
// @Produce      json
// @Param        address  query     string  true  "账户地址"
// @Success      200     {object}  AccountResponse
// @Failure      500     {object}  ErrorResponse
// @Router       /account/balance [get]
func GetAccountBalance(c *gin.Context) {
	address := c.Query("address")
	if address == "" {
		c.JSON(http.StatusBadRequest, ErrorResponse{Code: http.StatusBadRequest, Message: "地址不能为空"})
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

	c.JSON(http.StatusOK, AccountResponse{
		Address: address,
		Balance: balance.String(),
	})
}

// SetAccountBalance godoc
// @Summary      设置账户余额
// @Description  设置指定账户的以太坊余额（仅用于测试网络）
// @Tags         账户
// @Accept       json
// @Produce      json
// @Param        request  body      AccountBalance  true  "账户余额信息"
// @Success      200     {object}  AccountResponse
// @Failure      500     {object}  ErrorResponse
// @Router       /account/balance [post]
func SetAccountBalance(c *gin.Context) {
	var req AccountBalance
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Code: http.StatusBadRequest, Message: "无效的请求参数: " + err.Error()})
		return
	}

	client, err := utils.GetEthClientHTTP()
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Code: http.StatusInternalServerError, Message: "初始化以太坊客户端失败: " + err.Error()})
		return
	}

	account := common.HexToAddress(req.Address)
	balance, err := client.BalanceAt(c.Request.Context(), account, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Code: http.StatusInternalServerError, Message: "获取余额失败: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, AccountResponse{
		Address: req.Address,
		Balance: balance.String(),
	})
}

// CreateWallet godoc
// @Summary      创建钱包
// @Description  创建新的以太坊钱包
// @Tags         账户
// @Accept       json
// @Produce      json
// @Success      200  {object}  AccountResponse
// @Failure      500  {object}  ErrorResponse
// @Router       /account/wallet [post]
func CreateWallet(c *gin.Context) {
	account, err := utils.CreateAccount()
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Code: http.StatusInternalServerError, Message: "创建钱包失败: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, AccountResponse{
		Address: account.Address.Hex(),
		Balance: "0",
	})
}

// CreateKeystore godoc
// @Summary      创建密钥库
// @Description  创建新的以太坊密钥库文件
// @Tags         账户
// @Accept       json
// @Produce      json
// @Success      200  {object}  AccountResponse
// @Failure      500  {object}  ErrorResponse
// @Router       /account/keystore [post]
func CreateKeystore(c *gin.Context) {
	account, err := utils.CreateKeystore()
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Code: http.StatusInternalServerError, Message: "创建密钥库失败: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, AccountResponse{
		Address: account.Address.Hex(),
		Balance: "0",
	})
}

// CreateHDWallet godoc
// @Summary      创建 HD 钱包
// @Description  创建新的分层确定性钱包
// @Tags         账户
// @Accept       json
// @Produce      json
// @Success      200  {object}  map[string]string
// @Failure      500  {object}  ErrorResponse
// @Router       /account/hdwallet [post]
func CreateHDWallet(c *gin.Context) {
	if err := utils.Chdwallet(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "HD 钱包创建成功",
	})
}
