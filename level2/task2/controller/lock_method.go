package controller

import (
	"net/http"
	bindings2 "task2/contracts/bindings"
	"task2/utils"

	"github.com/gin-gonic/gin"
)

// ----- Lock合约方法 -----

// LockWithdraw godoc
// @Summary      执行Lock合约的Withdraw方法
// @Description  调用Lock合约的Withdraw方法，提取合约中的资金
// @Tags         锁定合约
// @Accept       application/json
// @Produce      application/json
// @Param        request body LockWithdrawRequest true "提取资金请求参数"
// @Success      200  {object}  ContractResponse
// @Failure      400  {object}  ErrorResponse
// @Failure      500  {object}  ErrorResponse
// @Router       /contract/lock/withdraw [post]
func LockWithdraw(c *gin.Context) {
	// 创建以太坊客户端连接
	client, err := utils.GetEthClientHTTP()
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "创建以太坊客户端连接失败: " + err.Error(),
		})
		return
	}

	// 获取合约交易选项
	auth, err := getTransactOpts(client)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "获取合约交易选项失败: " + err.Error(),
		})
		return
	}

	// 获取合约地址
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

	// 执行Withdraw方法
	tx, err := instance.Withdraw(auth)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "执行Withdraw方法失败: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, ContractResponse{
		Address: contractAddress.Hex(),
		TxHash:  tx.Hash().Hex(),
	})
}
