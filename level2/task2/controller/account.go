package controller

import (
	"net/http"
	"task2/utils"

	"github.com/gin-gonic/gin"
)

// SetAccountBalance 设置账户余额
func SetAccountBalance(c *gin.Context) {
	if err := utils.SetAccountBalance(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "账户余额设置成功",
	})
}

// GetBalance 获取账户余额
func GetBalance(c *gin.Context) {
	address := c.Query("address")
	if address == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "地址参数不能为空",
		})
		return
	}

	balance, err := utils.GetBalance(address)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"address": address,
		"balance": balance.String(),
	})
}

// CreateWallet 创建新钱包
func CreateWallet(c *gin.Context) {
	if err := utils.NewWallet(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "钱包创建成功",
	})
}

// CreateKeystore 创建 Keystore
func CreateKeystore(c *gin.Context) {
	if err := utils.CreateKs(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Keystore 创建成功",
	})
}

// CreateHDWallet 创建 HD 钱包
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
