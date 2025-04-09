package tasks

import (
	"pledge-backend/db"
	"pledge-backend/schedule/common"
	"pledge-backend/schedule/services"
	"time"

	"github.com/jasonlvhit/gocron"
)

// Task 启动所有定时任务
// 初始化并配置所有系统定时任务，包括质押池数据更新、代币价格监控等
// 使用gocron库实现定时调度，所有任务在UTC时区执行
func Task() {

	// 获取环境变量
	// 读取系统环境配置，用于定时任务的运行参数设置
	common.GetEnv()

	// 清空Redis数据库
	// 在任务启动时清除之前可能残留的缓存数据，确保数据一致性
	err := db.RedisFlushDB()
	if err != nil {
		panic("清除Redis数据库失败: " + err.Error())
	}

	// 初始化任务执行
	// 在启动定时任务前，先执行一次所有任务，确保系统有初始数据
	services.NewPool().UpdateAllPoolInfo()           // 更新所有质押池信息
	services.NewTokenPrice().UpdateContractPrice()   // 更新代币合约价格
	services.NewTokenSymbol().UpdateContractSymbol() // 更新代币合约符号
	services.NewTokenLogo().UpdateTokenLogo()        // 更新代币Logo
	services.NewBalanceMonitor().Monitor()           // 监控余额变化
	// services.NewTokenPrice().SavePlgrPrice()        // 保存PLGR代币主网价格（已注释）
	services.NewTokenPrice().SavePlgrPriceTestNet() // 保存PLGR代币测试网价格

	// 配置并启动周期性任务
	// 创建调度器并设置为UTC时区
	s := gocron.NewScheduler()
	s.ChangeLoc(time.UTC)

	// 每2分钟更新一次所有质押池信息
	// 确保前端显示的质押池数据（如余额、利率）保持最新
	_ = s.Every(2).Minutes().From(gocron.NextTick()).Do(services.NewPool().UpdateAllPoolInfo)

	// 每1分钟更新一次代币价格
	// 由于加密货币价格波动频繁，需要较高频率更新以保证价格数据准确性
	_ = s.Every(1).Minute().From(gocron.NextTick()).Do(services.NewTokenPrice().UpdateContractPrice)

	// 每2小时更新一次代币符号
	// 代币符号变化较少，低频率更新即可
	_ = s.Every(2).Hours().From(gocron.NextTick()).Do(services.NewTokenSymbol().UpdateContractSymbol)

	// 每2小时更新一次代币Logo
	// 代币Logo变化较少，低频率更新即可
	_ = s.Every(2).Hours().From(gocron.NextTick()).Do(services.NewTokenLogo().UpdateTokenLogo)

	// 每30分钟执行一次余额监控
	// 检查系统关键账户余额，在余额低于阈值时触发告警
	_ = s.Every(30).Minutes().From(gocron.NextTick()).Do(services.NewBalanceMonitor().Monitor)

	// 每30分钟更新一次PLGR代币价格（主网，已注释）
	//_ = s.Every(30).Minutes().From(gocron.NextTick()).Do(services.NewTokenPrice().SavePlgrPrice)

	// 每30分钟更新一次PLGR代币价格（测试网）
	// 保存PLGR代币在测试网络的最新价格数据
	_ = s.Every(30).Minutes().From(gocron.NextTick()).Do(services.NewTokenPrice().SavePlgrPriceTestNet)

	// 启动所有待执行的任务并阻塞主线程
	// 确保定时任务持续运行，直到程序被终止
	<-s.Start() // 启动所有待执行的任务

}
