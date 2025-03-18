package controller

import (
	"net/http"
	"task2/utils"

	"github.com/gin-gonic/gin"
)

// DeployContracts 部署所有合约
func DeployContracts(c *gin.Context) {
	if err := utils.DeployContracts(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "合约部署成功",
	})
}
