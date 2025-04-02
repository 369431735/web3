package utils

import (
	"errors"
	"os"
	"path/filepath"
)

// IsDir 判断给定路径是否为目录
// 检查指定路径是否存在且是一个目录
// path: 要检查的路径
// 返回: 如果是目录则返回true，否则返回false
func IsDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}

// IsFile 判断给定路径是否为文件
// 检查指定路径是否存在且是一个文件
// path: 要检查的路径
// 返回: 如果是文件则返回true，否则返回false
func IsFile(path string) bool {
	return !IsDir(path)
}

// 创建目录
//func MkDir(path string) error {
//	return os.MkdirAll(path, os.ModePerm)
//}

// MkDir 创建目录
// 在当前工作目录下创建指定的目录结构，如果目录已存在则不做任何操作
// dir: 要创建的目录路径（相对于项目根目录）
// 返回: 可能的错误信息
func MkDir(dir string) error {
	// 获取项目根目录的绝对路径
	rootDir, err := filepath.Abs("./")
	if err != nil {
		return errors.New("获取根目录失败，" + err.Error())
	}
	// 构建完整的目录路径
	dir = rootDir + "/" + dir
	// 检查目录是否已存在
	_, err = os.Stat(dir)
	if os.IsNotExist(err) {
		// 如果目录不存在，则创建目录及所有必要的父目录
		err = os.MkdirAll(dir, os.ModePerm)
		if err != nil {
			return err
		}
	}
	return nil
}
