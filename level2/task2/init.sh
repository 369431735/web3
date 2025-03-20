#!/bin/bash

# 创建必要的目录
mkdir -p contracts

# 删除旧文件和目录
echo "正在删除旧文件..."
rm -rf contracts/bindings
rm -rf abi

# 编译项目
echo "正在编译项目..."
go build -v ./...

echo "初始化完成！" 