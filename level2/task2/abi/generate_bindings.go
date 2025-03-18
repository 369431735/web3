package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

// ContractJSON 表示合约的JSON文件结构
type ContractJSON struct {
	ABI      json.RawMessage `json:"abi"`
	Bytecode string          `json:"bytecode"`
}

// GenerateBindings 生成所有合约的Go绑定代码
func GenerateBindings() error {
	// 设置ABI和字节码文件路径
	abiPath := os.Getenv("ABI_PATH")
	if abiPath == "" {
		abiPath = "D:/work/gitspace/web3/web3/hardhat-project/artifacts/contracts" // 默认路径
	}
	bindingsPath := "./bindings"

	// 创建bindings目录
	if err := os.MkdirAll(bindingsPath, 0755); err != nil {
		return fmt.Errorf("创建目录失败: %v", err)
	}

	// 合约列表及其文件名
	contracts := map[string]string{
		"ERC20MinerReward": "ERC20MinerReward",
		"SimpleStorage":    "SimpleStorage",
		"Lock":             "Lock",
		"Shipping":         "Shipping",
		"Lottery":          "Lottery",
		"SimpleAuction":    "SimpleAuction",
		"Purchase":         "Purchase",
		"ArrayDemo":        "ArrayDemo",
		"Ballot":           "Ballot",
		"Index":            "Index",
	}

	// 生成每个合约的绑定
	for contractName, fileName := range contracts {
		jsonFile := filepath.Join(abiPath, fileName+".sol", fileName+".json")
		outputFile := filepath.Join(bindingsPath, contractName+".go")

		// 读取JSON文件
		jsonData, err := os.ReadFile(jsonFile)
		if err != nil {
			fmt.Printf("读取%s的JSON文件失败: %v\n", contractName, err)
			continue
		}

		// 解析JSON文件
		var contractJSON ContractJSON
		if err := json.Unmarshal(jsonData, &contractJSON); err != nil {
			fmt.Printf("解析%s的JSON文件失败: %v\n", contractName, err)
			continue
		}

		// 创建临时文件
		tempAbiFile := filepath.Join(os.TempDir(), contractName+".abi")
		tempBinFile := filepath.Join(os.TempDir(), contractName+".bin")

		// 写入ABI文件
		if err := os.WriteFile(tempAbiFile, contractJSON.ABI, 0644); err != nil {
			fmt.Printf("创建临时ABI文件失败: %v\n", err)
			continue
		}

		// 写入字节码文件
		if err := os.WriteFile(tempBinFile, []byte(contractJSON.Bytecode), 0644); err != nil {
			fmt.Printf("创建临时字节码文件失败: %v\n", err)
			continue
		}

		cmd := exec.Command("abigen",
			"--abi", tempAbiFile,
			"--bin", tempBinFile,
			"--pkg", "bindings",
			"--type", contractName,
			"--out", outputFile,
		)

		if output, err := cmd.CombinedOutput(); err != nil {
			fmt.Printf("生成%s的绑定失败: %v\n%s\n", contractName, err, output)
			continue
		}

		// 清理临时文件
		os.Remove(tempAbiFile)
		os.Remove(tempBinFile)

		fmt.Printf("成功生成%s的绑定\n", contractName)
	}

	fmt.Println("所有合约绑定生成完成")
	return nil
}

func main() {
	if err := GenerateBindings(); err != nil {
		fmt.Printf("生成绑定失败: %v\n", err)
		os.Exit(1)
	}
}
