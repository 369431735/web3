package controller

import (
	"net/http"
	"task2/utils"

	"github.com/gin-gonic/gin"
)

// CheckAddress 检查以太坊地址格式
func CheckAddress(c *gin.Context) {
	if err := utils.AddressCheck(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "以太坊地址检查成功",
	})
}
