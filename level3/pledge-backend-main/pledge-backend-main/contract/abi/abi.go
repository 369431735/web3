package abi

import (
	"io/ioutil"
	"path"
	"path/filepath"
	"runtime"
)

// GetAbiByToken 根据代币标识获取对应的ABI文件内容
// ABI (Application Binary Interface) 文件定义了智能合约的接口，
// 包括函数签名、事件定义和参数类型，用于与区块链上的智能合约交互
// 参数：
//   - token: 代币标识或合约标识，用于定位具体的ABI文件
//
// 返回：
//   - string: ABI文件内容，格式为JSON字符串，描述了合约接口
//   - error: 错误信息，nil表示操作成功，非nil表示读取ABI文件失败
func GetAbiByToken(token string) (string, error) {
	// 获取当前目录的绝对路径
	currentAbPath := getCurrentAbPathByCaller()

	// 根据代币标识构建完整的ABI文件路径
	// 例如：/path/to/contract/abi/pledge_pool.abi
	abiFile, err := filepath.Abs(currentAbPath + "/" + token + ".abi")
	if err != nil {
		return "", err
	}

	// 读取ABI文件的内容
	abiBytes, err := ioutil.ReadFile(abiFile)
	if err != nil {
		return "", err
	}

	// 将ABI文件内容转换为字符串并返回
	return string(abiBytes), nil
}

// getCurrentAbPathByCaller 获取当前执行文件的绝对路径
// 用于确定ABI文件的存储位置，支持在不同环境下（如开发和生产）正确定位文件
//
// 返回：
//   - string: 当前ABI文件所在目录的绝对路径
//
// 说明：
//   - 该函数通过runtime.Caller获取调用栈信息，提取出当前文件的路径
//   - 支持在go run模式和编译后的二进制文件中正确工作
//   - 对于项目中移动文件位置的情况，只需更新此函数即可维护路径正确性
func getCurrentAbPathByCaller() string {
	// 获取调用者的文件路径
	// 参数0表示获取当前函数的调用信息
	_, filename, _, ok := runtime.Caller(0)

	// 提取目录部分
	var abPath string
	if ok {
		// 从完整文件路径中提取目录部分
		abPath = path.Dir(filename)
	}

	return abPath
}
