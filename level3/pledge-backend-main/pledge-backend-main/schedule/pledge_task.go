package main

import (
	"pledge-backend/db"
	"pledge-backend/schedule/models"
	"pledge-backend/schedule/tasks"
)

// main 定时任务服务的主入口函数
// 负责初始化系统所需的数据库连接和数据表结构，然后启动定时任务调度系统
// 定时任务服务是系统的核心组件之一，负责自动更新质押池信息、代币价格等数据
func main() {

	// 初始化MySQL数据库连接
	// 创建并配置系统所需的MySQL数据库连接，用于持久化存储质押池和代币信息
	db.InitMysql()

	// 初始化Redis连接
	// 创建并配置Redis连接池，用于缓存频繁访问的数据和存储临时状态
	db.InitRedis()

	// 创建并初始化数据表结构
	// 确保数据库中存在所有必要的表，以及这些表具有正确的结构
	models.InitTable()

	// 启动质押池相关的定时任务
	// 配置并开始执行所有定期任务，如更新质押池信息、监控代币价格等
	tasks.Task()

}

/*
 注意：如果更改版本，需要同时修改以下文件：
 config/init.go

 这确保配置文件和任务调度系统保持版本一致性，避免兼容性问题。
*/
