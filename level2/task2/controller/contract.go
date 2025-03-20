package controller

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"net/http"
	"strings"
	"time"

	"task2/config"
	"task2/contracts"
	"task2/storage"
	"task2/types"
	"task2/utils"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
)

// getContractTransactOpts 获取合约交易选项
func getContractTransactOpts(client *ethclient.Client) (*bind.TransactOpts, error) {
	// 获取网络配置
	network := config.GetCurrentNetwork()
	if network == nil {
		return nil, fmt.Errorf("未找到网络配置")
	}

	// 获取默认账户
	defaultAccount, ok := network.Accounts["default"]
	if !ok {
		return nil, fmt.Errorf("未找到默认账户")
	}

	// 获取私钥
	privateKey, err := utils.GetPrivateKey(defaultAccount.PrivateKey)
	if err != nil {
		return nil, fmt.Errorf("解析私钥失败: %v", err)
	}

	// 创建交易选项
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(network.ChainID))
	if err != nil {
		return nil, fmt.Errorf("创建交易选项失败: %v", err)
	}

	// 设置 gas 限制和价格
	auth.GasLimit = uint64(3000000)
	auth.GasPrice = big.NewInt(20000000000) // 20 Gwei

	return auth, nil
}

// DeployContracts 部署单个合约
// @Summary 部署单个合约
// @Description 部署指定的智能合约
// @Tags contracts
// @Accept json
// @Produce json
// @Param request body types.ContractDeployRequest true "部署请求参数"
// @Success 200 {object} types.ContractResponse
// @Failure 400 {object} ErrorResponse
// @Router /contracts/deploy [post]
func DeployContracts(c *gin.Context) {
	var request types.ContractDeployRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, ErrorResponse{Code: 400, Message: fmt.Sprintf("无效的请求参数: %v", err)})
		return
	}

	// 初始化以太坊客户端
	client, err := utils.GetEthClientHTTP()
	if err != nil {
		c.JSON(400, ErrorResponse{Code: 400, Message: fmt.Sprintf("初始化以太坊客户端失败: %v", err)})
		return
	}
	// 不需要defer client.Close()，因为我们使用的是单例客户端

	// 获取交易选项
	opts, err := getContractTransactOpts(client)
	if err != nil {
		c.JSON(400, ErrorResponse{Code: 400, Message: fmt.Sprintf("获取交易选项失败: %v", err)})
		return
	}

	// 根据合约类型部署合约
	var address common.Address
	var tx *ethTypes.Transaction

	switch request.ContractType {
	case "SimpleStorage":
		address, tx, _, err = contracts.DeploySimpleStorage(opts, client)
	case "Lock":
		unlockTime := time.Now().Add(24 * time.Hour).Unix()
		address, tx, _, err = contracts.DeployLock(opts, client, big.NewInt(unlockTime))
	default:
		c.JSON(400, ErrorResponse{Code: 400, Message: "不支持的合约类型"})
		return
	}

	if err != nil {
		c.JSON(400, ErrorResponse{Code: 400, Message: fmt.Sprintf("部署合约失败: %v", err)})
		return
	}

	// 保存合约地址到存储
	contractStorage := storage.GetInstance()
	if err := contractStorage.SetAddress(request.ContractType, address.Hex()); err != nil {
		log.Printf("保存合约地址失败: %v", err)
	}

	c.JSON(200, types.ContractResponse{
		ContractType: request.ContractType,
		Address:      address.Hex(),
		TxHash:       tx.Hash().Hex(),
	})
}

// DeployAllContracts 部署所有合约
// @Summary 部署所有合约
// @Description 部署所有支持的智能合约
// @Tags contracts
// @Produce json
// @Success 200 {object} map[string]types.ContractResponse
// @Failure 400 {object} ErrorResponse
// @Router /contracts/deploy-all [post]
func DeployAllContracts(c *gin.Context) {
	log.Printf("开始部署所有合约...")

	// 初始化以太坊客户端
	client, err := utils.GetEthClientHTTP()
	if err != nil {
		errMsg := fmt.Sprintf("初始化以太坊客户端失败: %v", err)
		log.Printf("错误: %s", errMsg)
		c.JSON(400, ErrorResponse{Code: 400, Message: errMsg})
		return
	}
	// 不需要defer client.Close()，因为我们使用的是单例客户端
	log.Printf("以太坊客户端初始化成功")

	// 获取交易选项
	opts, err := getContractTransactOpts(client)
	if err != nil {
		errMsg := fmt.Sprintf("获取交易选项失败: %v", err)
		log.Printf("错误: %s", errMsg)
		c.JSON(400, ErrorResponse{Code: 400, Message: errMsg})
		return
	}
	log.Printf("交易选项获取成功")

	// 定义要部署的合约列表及其部署函数
	type deployFunc func(*bind.TransactOpts, *ethclient.Client) (common.Address, *ethTypes.Transaction, error)

	deployFuncs := map[string]deployFunc{
		"SimpleStorage": contracts.DeploySimpleStorageFromBindings,
		"Lock":          contracts.DeployLockFromBindings,
		"Shipping":      contracts.DeployShippingFromBindings,
		"SimpleAuction": contracts.DeploySimpleAuctionFromBindings,
		"ArrayDemo":     contracts.DeployArrayDemoFromBindings,
		"Ballot":        contracts.DeployBallotFromBindings,
		"Lottery":       contracts.DeployLotteryFromBindings,
		"Purchase":      contracts.DeployPurchaseFromBindings,
	}

	results := make(map[string]types.ContractResponse)
	contractStorage := storage.GetInstance()
	log.Printf("准备部署 %d 个合约", len(deployFuncs))

	// 部署每个合约
	for contractType, deployFn := range deployFuncs {
		log.Printf("开始部署合约: %s", contractType)

		// 调用部署函数
		address, tx, err := deployFn(opts, client)

		if err != nil {
			errMsg := fmt.Sprintf("部署失败: %v", err)
			log.Printf("合约 %s 部署失败: %v", contractType, err)
			results[contractType] = types.ContractResponse{
				ContractType: contractType,
				Address:      "",
				TxHash:       "",
				Error:        errMsg,
			}
			continue
		}

		// 等待交易确认
		log.Printf("等待交易确认: %s", tx.Hash().Hex())
		receipt, err := bind.WaitMined(context.Background(), client, tx)
		if err != nil {
			errMsg := fmt.Sprintf("交易确认失败: %v", err)
			log.Printf("合约 %s 交易确认失败: %v", contractType, err)
			results[contractType] = types.ContractResponse{
				ContractType: contractType,
				Address:      address.Hex(),
				TxHash:       tx.Hash().Hex(),
				Error:        errMsg,
			}
			continue
		}

		if receipt.Status == 0 {
			errMsg := "交易失败"
			log.Printf("合约 %s 交易失败", contractType)
			results[contractType] = types.ContractResponse{
				ContractType: contractType,
				Address:      address.Hex(),
				TxHash:       tx.Hash().Hex(),
				Error:        errMsg,
			}
			continue
		}

		log.Printf("合约 %s 部署成功, 地址: %s, 交易哈希: %s", contractType, address.Hex(), tx.Hash().Hex())

		// 保存合约地址到存储
		if err := contractStorage.SetAddress(contractType, address.Hex()); err != nil {
			errMsg := fmt.Sprintf("保存合约地址失败: %v", err)
			log.Printf("保存合约地址到文件失败: %v", err)
			results[contractType] = types.ContractResponse{
				ContractType: contractType,
				Address:      address.Hex(),
				TxHash:       tx.Hash().Hex(),
				Error:        errMsg,
			}
			continue
		}
		log.Printf("合约地址已保存到文件: %s => %s", contractType, address.Hex())

		// 同时更新内存中的映射
		inMemoryName := strings.ToLower(contractType)
		types.DeployedContracts[inMemoryName] = address
		log.Printf("更新内存中的合约地址映射: %s => %s", inMemoryName, address.Hex())

		results[contractType] = types.ContractResponse{
			ContractType: contractType,
			Address:      address.Hex(),
			TxHash:       tx.Hash().Hex(),
		}
	}

	log.Printf("所有合约部署完成，共 %d 个合约", len(results))
	c.JSON(200, results)
}

// GetContractAddresses 获取所有已部署合约的地址
// @Summary 获取合约地址
// @Description 获取所有已部署合约的地址
// @Tags contracts
// @Produce json
// @Success 200 {object} storage.ContractAddresses
// @Failure 400 {object} ErrorResponse
// @Router /contracts [get]
func GetContractAddresses(c *gin.Context) {
	contractStorage := storage.GetInstance()
	addresses := contractStorage.GetAllAddresses()
	c.JSON(200, addresses)
}

// GetContractBytecode 获取合约字节码
// @Summary 获取合约字节码
// @Description 获取指定合约的字节码
// @Tags contracts
// @Accept json
// @Produce json
// @Param request body types.ContractBytecodeRequest true "合约字节码请求参数"
// @Success 200 {object} types.ContractBytecodeResponse
// @Failure 400 {object} ErrorResponse
// @Router /contract/bytecode [post]
func GetContractBytecode(c *gin.Context) {
	var request types.ContractBytecodeRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, ErrorResponse{Code: 400, Message: fmt.Sprintf("无效的请求参数: %v", err)})
		return
	}

	// 从存储中获取合约地址
	contractStorage := storage.GetInstance()
	address, err := contractStorage.GetAddress(request.ContractType)
	if err != nil {
		c.JSON(400, ErrorResponse{Code: 400, Message: fmt.Sprintf("获取合约地址失败: %v", err)})
		return
	}

	// 获取以太坊客户端
	client, err := utils.GetEthClientHTTP()
	if err != nil {
		c.JSON(400, ErrorResponse{Code: 400, Message: fmt.Sprintf("初始化以太坊客户端失败: %v", err)})
		return
	}
	// 不需要defer client.Close()，因为使用的是单例客户端

	// 获取合约字节码
	bytecode, err := client.CodeAt(c.Request.Context(), common.HexToAddress(address), nil)
	if err != nil {
		c.JSON(400, ErrorResponse{Code: 400, Message: fmt.Sprintf("获取合约字节码失败: %v", err)})
		return
	}

	c.JSON(200, types.ContractBytecodeResponse{
		ContractType: request.ContractType,
		Address:      address,
		Bytecode:     "0x" + common.Bytes2Hex(bytecode),
	})
}

// ContractController 处理与智能合约相关的请求
type ContractController struct{}

// DeployLockRequest Lock合约部署请求参数
type DeployLockRequest struct {
	UnlockTime int64 `json:"unlockTime" binding:"required"`
}

// DeployLock 部署Lock合约
func (c *ContractController) DeployLock(ctx *gin.Context) {
	var req DeployLockRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "无效的请求参数: " + err.Error()})
		return
	}

	// I连接到以太坊节点
	client, err := utils.GetEthClientHTTP()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "无法连接到以太坊网络: " + err.Error()})
		return
	}
	// 不需要defer client.Close()，因为使用的是单例客户端

	// 获取交易选项
	opts, err := getContractTransactOpts(client)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "获取交易选项失败: " + err.Error()})
		return
	}

	// 使用default账户进行部署
	network := config.GetCurrentNetwork()
	if network == nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "未找到网络配置"})
		return
	}

	defaultAccount, ok := network.Accounts["default"]
	if !ok {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "未找到默认账户"})
		return
	}

	log.Printf("使用默认账户 %s 部署Lock合约", defaultAccount.Address)

	// 部署合约
	address, tx, _, err := contracts.DeployLock(opts, client, big.NewInt(req.UnlockTime))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "部署Lock合约失败: " + err.Error()})
		return
	}

	// 保存合约地址到存储
	contractStorage := storage.GetInstance()
	if err := contractStorage.SetAddress("Lock", address.Hex()); err != nil {
		log.Printf("保存Lock合约地址失败: %v", err)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Lock合约部署成功",
		"address": address.Hex(),
		"txHash":  tx.Hash().Hex(),
	})
}

// DeploySimpleStorage 部署SimpleStorage合约
func (c *ContractController) DeploySimpleStorage(ctx *gin.Context) {
	// 获取以太坊客户端
	client, err := utils.GetEthClientHTTP()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "无法连接到以太坊网络: " + err.Error()})
		return
	}
	// 不需要defer client.Close()，因为使用的是单例客户端

	// 获取交易选项
	opts, err := getContractTransactOpts(client)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "获取交易选项失败: " + err.Error()})
		return
	}

	// 使用default账户进行部署
	network := config.GetCurrentNetwork()
	if network == nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "未找到网络配置"})
		return
	}

	defaultAccount, ok := network.Accounts["default"]
	if !ok {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "未找到默认账户"})
		return
	}

	log.Printf("使用默认账户 %s 部署SimpleStorage合约", defaultAccount.Address)

	// 部署合约
	address, tx, _, err := contracts.DeploySimpleStorage(opts, client)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "部署SimpleStorage合约失败: " + err.Error()})
		return
	}

	// 保存合约地址到存储
	contractStorage := storage.GetInstance()
	if err := contractStorage.SetAddress("SimpleStorage", address.Hex()); err != nil {
		log.Printf("保存SimpleStorage合约地址失败: %v", err)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "SimpleStorage合约部署成功",
		"address": address.Hex(),
		"txHash":  tx.Hash().Hex(),
	})
}
