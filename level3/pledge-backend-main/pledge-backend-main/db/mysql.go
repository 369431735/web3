package db

import (
	"fmt"
	"pledge-backend/config"
	"pledge-backend/log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

// InitMysql 初始化MySQL数据库连接
// 配置并创建与MySQL数据库的连接，设置连接池参数，注册GORM回调函数
// 该函数在应用启动时被调用，初始化全局MySQL连接对象
func InitMysql() {
	mysqlConf := config.Config.Mysql
	log.Logger.Info("Init Mysql")
	// 构建数据库连接字符串
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		mysqlConf.UserName, // 数据库用户名
		mysqlConf.Password, // 数据库密码
		mysqlConf.Address,  // 数据库服务器地址
		mysqlConf.Port,     // 数据库端口
		mysqlConf.DbName)   // 数据库名称

	// 使用GORM打开数据库连接
	// 配置MySQL驱动程序选项以适应不同版本的MySQL数据库
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,   // DSN数据源名称
		DefaultStringSize:         256,   // string类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用datetime精度，MySQL 5.6之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7之前的数据库和MariaDB不支持重命名索引
		DontSupportRenameColumn:   true,  // 用`change`重命名列，MySQL 8之前的数据库和MariaDB不支持重命名列
		SkipInitializeWithVersion: false, // 根据当前MySQL版本自动配置
	}), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 关闭复数表名（不在表名后加s）, 例如：user表而不是users表
		},
		SkipDefaultTransaction: true, // 跳过默认事务，提高性能
	})
	if err != nil {
		log.Logger.Panic(fmt.Sprintf("mysql connention error ==>  %+v", err))
	}

	// 注册GORM的回调函数，用于数据库操作完成后的处理
	// 这些回调可用于记录SQL日志、性能监控或其他自定义操作
	_ = db.Callback().Create().After("gorm:after_create").Register("after_create", After)
	_ = db.Callback().Query().After("gorm:after_query").Register("after_query", After)
	_ = db.Callback().Delete().After("gorm:after_delete").Register("after_delete", After)
	_ = db.Callback().Update().After("gorm:after_update").Register("after_update", After)
	_ = db.Callback().Row().After("gorm:row").Register("after_row", After)
	_ = db.Callback().Raw().After("gorm:raw").Register("after_raw", After)

	// 自动迁移功能（已注释）
	// 自动迁移为给定模型运行自动迁移，只会添加缺失的字段，不会删除/更改当前数据
	// db.AutoMigrate(&TestTable{})

	// 获取底层的SQL数据库连接
	sqlDB, err := db.DB()
	if err != nil {
		log.Logger.Error("db.DB() err:" + err.Error())
	}

	// 设置连接池参数
	// 配置合适的连接池参数对于应用性能至关重要
	// 参考技术文档: https://colobu.com/2019/05/27/configuring-sql-DB-for-better-performance/
	sqlDB.SetMaxIdleConns(mysqlConf.MaxIdleConns)                                // 空闲连接数 - 保持在连接池中的最大空闲连接数
	sqlDB.SetMaxOpenConns(mysqlConf.MaxOpenConns)                                // 最大连接数 - 与数据库的最大打开连接数
	sqlDB.SetConnMaxLifetime(time.Duration(mysqlConf.MaxLifeTime) * time.Second) // 连接最大生存时间 - 连接可重用的最长时间
	Mysql = db                                                                   // 将数据库连接赋值给全局变量
}

// After GORM框架的回调函数，用于处理数据库操作后的逻辑
// 该函数在每次数据库操作完成后被调用
// 可用于打印SQL语句、记录执行时间、监控数据库操作等
// 参数：
//   - db: GORM数据库实例，包含当前SQL操作的相关信息
func After(db *gorm.DB) {
	// 解释执行的SQL语句，但不打印
	db.Dialector.Explain(db.Statement.SQL.String(), db.Statement.Vars...)

	// 如需记录SQL语句，可取消以下注释
	// sql := db.Dialector.Explain(db.Statement.SQL.String(), db.Statement.Vars...)
	// log.Logger.Info(sql)
}
