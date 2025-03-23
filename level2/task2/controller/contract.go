package controller

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"task2/config"
	"task2/contracts/deploy"
	"task2/storage"
	"task2/utils"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
)

// getContractTransactOpts 获取合约交易所需的交易选项
func getContractTransactOpts(client *ethclient.Client) (*bind.TransactOpts, error) {
	// 获取当前网络配置
	network := config.GetCurrentNetwork()
	if network == nil {
		return nil, fmt.Errorf("获取当前网络配置失败")
	}

	// 获取默认账户信息
	defaultAccount, ok := network.Accounts["default"]
	if !ok {
		return nil, fmt.Errorf("获取默认账户信息失败")
	}

	// 获取私钥
	privateKey, err := utils.GetPrivateKey(defaultAccount.PrivateKey)
	if err != nil {
		return nil, fmt.Errorf("获取私钥失败: %v", err)
	}

	// 创建交易选项
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(network.ChainID))
	if err != nil {
		return nil, fmt.Errorf("创建交易选项失败: %v", err)
	}

	// 设置 gas 参数
	auth.GasLimit = uint64(3000000)
	auth.GasPrice = big.NewInt(20000000000) // 20 Gwei

	return auth, nil
}

// DeployAllContracts godoc
// @Summary      部署所有合约
// @Description  部署所有合约，包括SimpleStorage、SimpleAuction、Purchase、Lottery等智能合约
// @Tags         合约
// @Accept       application/json
// @Produce      application/json
// @Success      200  {object}  map[string]string
// @Failure      500  {object}  ErrorResponse
// @Router       /contracts/deploy-all [post]
func DeployAllContracts(c *gin.Context) {
	log.Printf("开始部署所有合约，包括SimpleStorage、SimpleAuction、Purchase、Lottery等智能合约..")

	// 创建以太坊客户端
	client, err := utils.GetEthClientHTTP()
	if err != nil {
		errMsg := fmt.Sprintf("创建以太坊客户端失败: %v", err)
		log.Printf("错误: %s", errMsg)
		c.JSON(400, ErrorResponse{Code: 400, Message: errMsg})
		return
	}
	// 关闭客户端连接，defer client.Close()在函数结束时执行
	log.Printf("以太坊客户端连接成功")

	// 获取合约交易选项
	opts, err := getContractTransactOpts(client)
	if err != nil {
		errMsg := fmt.Sprintf("获取合约交易选项失败: %v", err)
		log.Printf("错误: %s", errMsg)
		c.JSON(400, ErrorResponse{Code: 400, Message: errMsg})
		return
	}
	log.Printf("合约交易选项获取成功")

	// 定义部署函数类型，用于统一处理各种合约的部署
	type deployFunc func(*bind.TransactOpts, *ethclient.Client) (common.Address, *ethTypes.Transaction, error)

	deployFuncs := map[string]deployFunc{
		"SimpleStorage": deploy.DeploySimpleStorageFromBindings,
		"Lock":          deploy.DeployLockFromBindings,
		"Shipping":      deploy.DeployShippingFromBindings,
		"SimpleAuction": deploy.DeploySimpleAuctionFromBindings,
		"ArrayDemo":     deploy.DeployArrayDemoFromBindings,
		"Ballot":        deploy.DeployBallotFromBindings,
		"Lottery":       deploy.DeployLotteryFromBindings,
		"Purchase":      deploy.DeployPurchaseFromBindings,
	}

	results := make(map[string]ContractResponse)
	contractStorage := storage.GetInstance()
	log.Printf("将部署 %d 个合约", len(deployFuncs))

	// 遍历所有合约类型，依次部署合约
	for contractType, deployFn := range deployFuncs {
		log.Printf("开始部署合约 %s", contractType)

		// 部署合约
		address, tx, err := deployFn(opts, client)
		if err != nil {
			errMsg := fmt.Sprintf("部署合约失败: %v", err)
			log.Printf("错误: %s %s %v", contractType, errMsg, err)
			results[contractType] = ContractResponse{
				ContractType: contractType,
				Address:      "",
				TxHash:       "",
				Error:        errMsg,
			}
			continue
		}

		// 等待交易确认
		log.Printf("等待交易确认 %s", tx.Hash().Hex())
		receipt, err := bind.WaitMined(context.Background(), client, tx)
		if err != nil {
			errMsg := fmt.Sprintf("等待交易确认失败: %v", err)
			log.Printf("错误: %s %s %v", contractType, errMsg, err)
			results[contractType] = ContractResponse{
				ContractType: contractType,
				Address:      address.Hex(),
				TxHash:       tx.Hash().Hex(),
				Error:        errMsg,
			}
			continue
		}

		if receipt.Status == 0 {
			errMsg := "交易失败，合约部署未完成"
			log.Printf("错误: %s %s", contractType, errMsg)
			results[contractType] = ContractResponse{
				ContractType: contractType,
				Address:      address.Hex(),
				TxHash:       tx.Hash().Hex(),
				Error:        errMsg,
			}
			continue
		}

		log.Printf("合约 %s 部署成功，合约地址: %s, 交易哈希: %s", contractType, address.Hex(), tx.Hash().Hex())

		// 保存合约地址到存储
		if err := contractStorage.SetAddress(contractType, address.Hex()); err != nil {
			errMsg := fmt.Sprintf("保存合约地址失败: %v", err)
			log.Printf("错误: %s %v", contractType, err)
			results[contractType] = ContractResponse{
				ContractType: contractType,
				Address:      address.Hex(),
				TxHash:       tx.Hash().Hex(),
				Error:        errMsg,
			}
			continue
		}
		log.Printf("合约 %s 部署成功，合约地址: %s, 交易哈希: %s", contractType, address.Hex(), tx.Hash().Hex())

		results[contractType] = ContractResponse{
			ContractType: contractType,
			Address:      address.Hex(),
			TxHash:       tx.Hash().Hex(),
		}
	}

	log.Printf("所有合约部署完成，共部署 %d 个合约", len(results))
	c.JSON(200, results)
}

// GetContractAddresses godoc
// @Summary      获取所有合约地址
// @Description  获取所有合约地址，包括SimpleStorage、SimpleAuction、Purchase、Lottery等智能合约
// @Tags         合约
// @Accept       application/json
// @Produce      application/json
// @Success      200  {object}  map[string]string
// @Failure      500  {object}  ErrorResponse
// @Router       /contracts/allAddresses [get]
func GetContractAddresses(c *gin.Context) {
	contractStorage := storage.GetInstance()
	addresses := contractStorage.GetAllAddresses()
	c.JSON(200, addresses)
}

// GetContractBytecode 获取合约字节码
// @Summary      获取合约字节码
// @Description  获取指定合约的字节码
// @Tags         合约
// @Accept       application/json
// @Produce      application/json
// @Param request body ContractBytecodeRequest true "合约类型和合约地址"
// @Success      200  {object}  ContractBytecodeResponse
// @Failure      400  {object}  ErrorResponse
// @Router       /contracts/bytecode [post]
func GetContractBytecode(c *gin.Context) {
	var request ContractBytecodeRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, ErrorResponse{Code: 400, Message: fmt.Sprintf("绑定请求失败: %v", err)})
		return
	}

	// 获取合约存储实例
	contractStorage := storage.GetInstance()
	address, err := contractStorage.GetAddress(request.ContractType)
	if err != nil {
		c.JSON(400, ErrorResponse{Code: 400, Message: fmt.Sprintf("获取合约地址失败: %v", err)})
		return
	}

	// 创建以太坊客户端
	client, err := utils.GetEthClientHTTP()
	if err != nil {
		c.JSON(400, ErrorResponse{Code: 400, Message: fmt.Sprintf("创建以太坊客户端失败: %v", err)})
		return
	}

	// 获取合约字节码
	bytecode, err := client.CodeAt(c.Request.Context(), common.HexToAddress(address), nil)
	if err != nil {
		c.JSON(400, ErrorResponse{Code: 400, Message: fmt.Sprintf("获取合约字节码失败: %v", err)})
		return
	}

	c.JSON(200, ContractBytecodeResponse{
		ContractType: request.ContractType,
		Address:      address,
		Bytecode:     "0x" + common.Bytes2Hex(bytecode),
	})
}
