package controller

import (
	"math/big"
	"net/http"
	"time"

	"task2/contracts"
	"task2/types"
	"task2/utils"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/gin-gonic/gin"
)

// 存储已部署合约的地址
var DeployedContracts = make(map[string]common.Address)

// 注册合约地址到映射表
func registerContract(name string, address common.Address) {
	DeployedContracts[name] = address
}

// DeployContracts godoc
// @Summary      部署合约
// @Description  部署智能合约
// @Tags         合约
// @Accept       json
// @Produce      json
// @Param        request  body      types.ContractDeployRequest  true  "合约部署参数"
// @Success      200     {object}  types.ContractResponse
// @Failure      500     {object}  ErrorResponse
// @Router       /contracts/deploy [post]
func DeployContracts(c *gin.Context) {
	var req types.ContractDeployRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Code: http.StatusBadRequest, Message: "无效的请求参数: " + err.Error()})
		return
	}

	// 初始化以太坊客户端
	client, err := utils.InitClient()
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Code: http.StatusInternalServerError, Message: "初始化以太坊客户端失败: " + err.Error()})
		return
	}

	// 获取网络配置
	network := utils.GetCurrentNetwork()
	if network == nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Code: http.StatusInternalServerError, Message: "未找到网络配置"})
		return
	}

	// 获取默认账户
	defaultAccount, ok := network.Accounts["default"]
	if !ok {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Code: http.StatusInternalServerError, Message: "未找到默认账户"})
		return
	}

	// 获取私钥
	privateKey, err := crypto.HexToECDSA(defaultAccount.PrivateKey[2:]) // 移除 "0x" 前缀
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Code: http.StatusInternalServerError, Message: "解析私钥失败: " + err.Error()})
		return
	}

	// 创建交易选项
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(network.ChainID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Code: http.StatusInternalServerError, Message: "创建交易选项失败: " + err.Error()})
		return
	}

	// 根据合约名称部署对应的合约
	var address common.Address
	var tx *ethTypes.Transaction

	switch req.ContractName {
	case "SimpleStorage":
		address, tx, _, err = contracts.DeploySimpleStorage(auth, client)
	case "Lock":
		unlockTime := time.Now().Add(24 * time.Hour).Unix()
		address, tx, _, err = contracts.DeployLock(auth, client, big.NewInt(unlockTime))
	case "Shipping":
		address, tx, _, err = contracts.DeployShipping(auth, client)
	case "SimpleAuction":
		beneficiary := common.HexToAddress(defaultAccount.Address)
		biddingTime := big.NewInt(3600) // 1小时
		address, tx, _, err = contracts.DeploySimpleAuction(auth, client, biddingTime, beneficiary)
	case "ArrayDemo":
		address, tx, _, err = contracts.DeployArrayDemo(auth, client)
	default:
		c.JSON(http.StatusBadRequest, ErrorResponse{Code: http.StatusBadRequest, Message: "不支持的合约名称: " + req.ContractName})
		return
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Code: http.StatusInternalServerError, Message: "部署合约失败: " + err.Error()})
		return
	}

	// 注册合约地址
	types.RegisterContract(req.ContractName, address)

	c.JSON(http.StatusOK, types.ContractResponse{
		Address: address.Hex(),
		TxHash:  tx.Hash().Hex(),
	})
}

// DeployAllContracts godoc
// @Summary      部署所有合约
// @Description  一次性部署所有支持的智能合约
// @Tags         合约
// @Accept       json
// @Produce      json
// @Success      200  {object}  map[string]types.ContractResponse
// @Failure      500  {object}  ErrorResponse
// @Router       /contracts/deploy-all [post]
func DeployAllContracts(c *gin.Context) {
	// 初始化以太坊客户端
	client, err := utils.InitClient()
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Code: http.StatusInternalServerError, Message: "初始化以太坊客户端失败: " + err.Error()})
		return
	}

	// 获取网络配置
	network := utils.GetCurrentNetwork()
	if network == nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Code: http.StatusInternalServerError, Message: "未找到网络配置"})
		return
	}

	// 获取默认账户
	defaultAccount, ok := network.Accounts["default"]
	if !ok {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Code: http.StatusInternalServerError, Message: "未找到默认账户"})
		return
	}

	// 获取私钥
	privateKey, err := crypto.HexToECDSA(defaultAccount.PrivateKey[2:]) // 移除 "0x" 前缀
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Code: http.StatusInternalServerError, Message: "解析私钥失败: " + err.Error()})
		return
	}

	// 创建交易选项
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(network.ChainID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Code: http.StatusInternalServerError, Message: "创建交易选项失败: " + err.Error()})
		return
	}

	// 定义要部署的合约列表
	contractList := []string{
		"SimpleStorage",
		"Lock",
		"Shipping",
		"SimpleAuction",
		"ArrayDemo",
	}

	// 存储部署结果
	results := make(map[string]types.ContractResponse)
	hasSuccess := false

	// 遍历部署所有合约
	for _, contractName := range contractList {
		var address common.Address
		var tx *ethTypes.Transaction

		switch contractName {
		case "SimpleStorage":
			address, tx, _, err = contracts.DeploySimpleStorage(auth, client)
		case "Lock":
			unlockTime := time.Now().Add(24 * time.Hour).Unix()
			address, tx, _, err = contracts.DeployLock(auth, client, big.NewInt(unlockTime))
		case "Shipping":
			address, tx, _, err = contracts.DeployShipping(auth, client)
		case "SimpleAuction":
			beneficiary := common.HexToAddress(defaultAccount.Address)
			biddingTime := big.NewInt(3600) // 1小时
			address, tx, _, err = contracts.DeploySimpleAuction(auth, client, biddingTime, beneficiary)
		case "ArrayDemo":
			address, tx, _, err = contracts.DeployArrayDemo(auth, client)
		}

		if err != nil {
			utils.LogError("部署合约失败: "+contractName, err)
			continue
		}

		// 注册合约地址
		types.RegisterContract(contractName, address)
		hasSuccess = true

		// 记录部署结果
		txHash := "pending"
		if tx != nil {
			txHash = tx.Hash().Hex()
		}

		results[contractName] = types.ContractResponse{
			Address: address.Hex(),
			TxHash:  txHash,
		}
	}

	if !hasSuccess {
		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "所有合约部署失败",
		})
		return
	}

	c.JSON(http.StatusOK, results)
}

// GetContractAddresses godoc
// @Summary      获取合约地址
// @Description  获取所有已部署合约的地址
// @Tags         合约
// @Accept       json
// @Produce      json
// @Success      200  {object}  map[string]string
// @Router       /contracts [get]
func GetContractAddresses(c *gin.Context) {
	c.JSON(200, types.GetDeployedContracts())
}

// GetContractBytecode godoc
// @Summary      获取智能合约字节码
// @Description  根据合约地址获取智能合约的字节码
// @Tags         合约
// @Accept       json
// @Produce      json
// @Param        request body types.ContractBytecodeRequest true "合约地址"
// @Success      200  {object}  types.ContractBytecodeResponse
// @Failure      400  {object}  ErrorResponse
// @Failure      500  {object}  ErrorResponse
// @Router       /contract/bytecode [post]
func GetContractBytecode(c *gin.Context) {
	var req types.ContractBytecodeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "无效的请求参数: " + err.Error(),
		})
		return
	}

	// 初始化以太坊客户端
	client, err := utils.InitClient()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "初始化以太坊客户端失败: " + err.Error(),
		})
		return
	}

	// 获取合约字节码
	bytecode, err := client.CodeAt(c.Request.Context(), common.HexToAddress(req.Address), nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	// 将字节码转换为十六进制字符串
	c.JSON(http.StatusOK, types.ContractBytecodeResponse{
		Bytecode: "0x" + common.Bytes2Hex(bytecode),
	})
}
