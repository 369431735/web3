package controller

import (
	"math/big"
	"net/http"
	bindings2 "task2/contracts/bindings"
	"task2/utils"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/gin-gonic/gin"
)

// ----- Ballot投票合约相关操作 -----

// BallotVote godoc
// @Summary      投票功能
// @Description  调用Ballot合约的vote方法进行投票
// @Tags         Ballot
// @Accept       application/json
// @Produce      application/json
// @Param        request body BallotVoteRequest true "投票请求参数"
// @Success      200  {object}  SuccessResponse
// @Failure      400  {object}  ErrorResponse
// @Failure      500  {object}  ErrorResponse
// @Router       /contracts/Ballot/vote [post]
func BallotVote(c *gin.Context) {
	var req BallotVoteRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":   http.StatusBadRequest,
			"errMsg": "参数验证失败: " + err.Error(),
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

	// 获取合约地址
	contractAddress, err := getContractAddress("Ballot")
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":   http.StatusInternalServerError,
			"errMsg": "获取Ballot合约地址失败: " + err.Error(),
			"data":   nil,
		})
		return
	}

	// 转换提案ID
	proposalId := new(big.Int)
	_, ok := proposalId.SetString(req.ProposalId, 10)
	if !ok {
		c.JSON(http.StatusOK, gin.H{
			"code":   http.StatusBadRequest,
			"errMsg": "无效的提案ID格式",
			"data":   nil,
		})
		return
	}

	// 创建合约实例
	instance, err := bindings2.NewBallot(contractAddress, client)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":   http.StatusInternalServerError,
			"errMsg": "创建合约实例失败: " + err.Error(),
			"data":   nil,
		})
		return
	}

	// 调用vote方法
	tx, err := instance.Vote(auth, proposalId)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":   http.StatusInternalServerError,
			"errMsg": "调用vote方法失败: " + err.Error(),
			"data":   nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": gin.H{
			"message":          "投票成功",
			"transaction_hash": tx.Hash().Hex(),
			"contract_address": contractAddress.Hex(),
			"proposal_id":      req.ProposalId,
		},
		"errMsg": "",
	})
}

// BallotWinningProposal godoc
// @Summary      获取获胜提案ID
// @Description  调用Ballot合约的winningProposal方法获取当前得票最多的提案ID
// @Tags         Ballot
// @Accept       application/json
// @Produce      application/json
// @Success      200  {object}  SuccessResponse
// @Failure      500  {object}  ErrorResponse
// @Router       /contracts/Ballot/winner [get]
func BallotWinningProposal(c *gin.Context) {
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
	contractAddress, err := getContractAddress("Ballot")
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":   http.StatusInternalServerError,
			"errMsg": "获取Ballot合约地址失败: " + err.Error(),
			"data":   nil,
		})
		return
	}

	// 创建合约实例
	instance, err := bindings2.NewBallot(contractAddress, client)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":   http.StatusInternalServerError,
			"errMsg": "创建合约实例失败: " + err.Error(),
			"data":   nil,
		})
		return
	}

	// 调用winningProposal方法
	winningId, err := instance.WinningProposal(&bind.CallOpts{})
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":   http.StatusInternalServerError,
			"errMsg": "调用winningProposal方法失败: " + err.Error(),
			"data":   nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": gin.H{
			"winning_proposal": winningId.String(),
			"contract_address": contractAddress.Hex(),
		},
		"errMsg": "",
	})
}

// BallotWinnerName godoc
// @Summary      获取获胜提案名称
// @Description  调用Ballot合约的winnerName方法获取当前得票最多的提案名称
// @Tags         Ballot
// @Accept       application/json
// @Produce      application/json
// @Success      200  {object}  SuccessResponse
// @Failure      500  {object}  ErrorResponse
// @Router       /contracts/Ballot/winner-name [get]
func BallotWinnerName(c *gin.Context) {
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
	contractAddress, err := getContractAddress("Ballot")
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":   http.StatusInternalServerError,
			"errMsg": "获取Ballot合约地址失败: " + err.Error(),
			"data":   nil,
		})
		return
	}

	// 创建合约实例
	instance, err := bindings2.NewBallot(contractAddress, client)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":   http.StatusInternalServerError,
			"errMsg": "创建合约实例失败: " + err.Error(),
			"data":   nil,
		})
		return
	}

	// 调用winnerName方法
	winnerNameBytes, err := instance.WinnerName(&bind.CallOpts{})
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":   http.StatusInternalServerError,
			"errMsg": "调用winnerName方法失败: " + err.Error(),
			"data":   nil,
		})
		return
	}

	// 将字节数组转换为字符串，去除尾部的零值字节
	var winnerName string
	for _, b := range winnerNameBytes {
		if b == 0 {
			break
		}
		winnerName += string(b)
	}

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": gin.H{
			"winner_name":      winnerName,
			"contract_address": contractAddress.Hex(),
		},
		"errMsg": "",
	})
}

// BallotVoteRequest 投票请求参数
type BallotVoteRequest struct {
	ProposalId string `json:"proposal_id" binding:"required"` // 提案ID
}
