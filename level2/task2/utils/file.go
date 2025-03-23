package utils

import (
	"fmt"
	"os"
	"path/filepath"
)

// ReadFile 读取指定路径的文件内容
func ReadFile(path string) ([]byte, error) {
	// 规范化路径
	absPath, err := filepath.Abs(path)
	if err != nil {
		return nil, fmt.Errorf("获取绝对路径失败: %w", err)
	}

	// 检查文件是否存在
	_, err = os.Stat(absPath)
	if os.IsNotExist(err) {
		return nil, fmt.Errorf("文件不存在: %s", absPath)
	} else if err != nil {
		return nil, fmt.Errorf("检查文件状态失败: %w", err)
	}

	// 读取文件内容
	data, err := os.ReadFile(absPath)
	if err != nil {
		return nil, fmt.Errorf("读取文件失败: %w", err)
	}

	return data, nil
}
