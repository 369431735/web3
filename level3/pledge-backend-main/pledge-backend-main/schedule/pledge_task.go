package main

import (
	"pledge-backend/db"
	"pledge-backend/schedule/models"
	"pledge-backend/schedule/tasks"
)

// main 定时任务服务的主入口函数，初始化数据库连接并启动任务调度
func main() {

	// 初始化MySQL数据库连接
	db.InitMysql()

	// 初始化Redis连接
	db.InitRedis()

	// 创建并初始化数据表结构
	models.InitTable()

	// 启动质押池相关的定时任务
	tasks.Task()

}

/*
 If you change the version, you need to modify the following files'
 config/init.go
*/
