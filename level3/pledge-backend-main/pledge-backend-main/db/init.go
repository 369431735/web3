package db

import (
	"github.com/gomodule/redigo/redis"
	"gorm.io/gorm"
)

// Mysql 全局MySQL数据库连接对象
// 用于在整个应用程序中共享MySQL数据库连接
var Mysql *gorm.DB

// RedisConn 全局Redis连接池对象
// 用于在整个应用程序中共享Redis连接
var RedisConn *redis.Pool
