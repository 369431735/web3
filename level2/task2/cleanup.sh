#!/bin/bash

# 确保脚本在出错时停止执行
set -e

echo "开始清理并更新合约文件..."

# 替换Lock合约文件
echo "更新Lock合约文件..."
mv contracts/lock_new.go contracts/Lock.go

# 替换SimpleStorage合约文件
echo "更新SimpleStorage合约文件..."
mv contracts/simple_storage_new.go contracts/SimpleStorage.go

# 替换部署文件
echo "更新部署文件..."
mv contracts/deploy_new.go contracts/deploy.go

# 替换事件控制器
echo "更新事件控制器..."
mv controller/events_new.go controller/events.go

# 删除旧的绑定文件夹
echo "删除旧的绑定文件夹..."
rm -rf task2/contracts/bindings 2>/dev/null || true
rm -rf task2/contracts 2>/dev/null || true

echo "清理完成！现在项目结构已经更新，重复的声明已被移除。"
echo "你现在可以正常编译和运行项目了。" 