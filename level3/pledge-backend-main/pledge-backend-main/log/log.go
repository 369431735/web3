package log

import (
	"os"
	"path"
	"runtime"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

// Logger 全局日志对象
// 使用zap日志库的实例，提供高性能的结构化日志记录
var Logger *zap.Logger

// init 初始化日志系统
// 在程序启动时自动执行，配置并创建全局日志对象
// 设置日志输出格式、日志级别和日志文件归档策略
func init() {

	// zap 不支持文件归档，如果要支持文件按大小或者时间归档，需要使用lumberjack，lumberjack也是zap官方推荐的。
	// https://github.com/uber-go/zap/blob/master/FAQ.md
	hook := lumberjack.Logger{
		Filename:   getCurrentAbPathByCaller() + "/logs/log.log", // 日志文件路径，使用当前目录的logs文件夹
		MaxSize:    50,                                           // 每个日志文件保存的最大尺寸（单位：MB）
		MaxBackups: 20,                                           // 日志文件最多保存多少个备份
		MaxAge:     7,                                            // 文件最多保存多少天
		Compress:   true,                                         // 是否压缩旧日志文件
	}

	// 定义日志内容的编码配置
	// 指定日志中各字段的名称和编码方式
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",                         // 时间字段的键名
		LevelKey:       "level",                        // 日志级别字段的键名
		NameKey:        "logger",                       // 记录器名称字段的键名
		CallerKey:      "line",                         // 调用者信息字段的键名
		MessageKey:     "msg",                          // 消息体字段的键名
		StacktraceKey:  "stacktrace",                   // 堆栈跟踪字段的键名
		LineEnding:     zapcore.DefaultLineEnding,      // 行尾符号
		EncodeLevel:    zapcore.LowercaseLevelEncoder,  // 日志级别使用小写字母编码
		EncodeTime:     zapcore.ISO8601TimeEncoder,     // 时间格式使用ISO8601标准的UTC时间格式
		EncodeDuration: zapcore.SecondsDurationEncoder, // 持续时间使用秒为单位
		EncodeCaller:   zapcore.FullCallerEncoder,      // 调用者信息使用完整路径编码
		EncodeName:     zapcore.FullNameEncoder,        // 记录器名称使用完整名称
	}

	// 设置日志级别
	// 创建原子级别，允许在运行时动态修改日志级别
	atomicLevel := zap.NewAtomicLevel()
	atomicLevel.SetLevel(zap.InfoLevel) // 设置为Info级别，将只记录Info级别及以上的日志

	// 创建核心日志组件
	// 组合编码器配置、输出目标和日志级别
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),                                           // 使用JSON编码器，根据上面的编码器配置
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(&hook)), // 同时输出到标准输出和日志文件
		atomicLevel, // 使用上面设置的日志级别
	)

	// 增强日志功能的选项
	caller := zap.AddCaller()                                // 开启开发模式，记录调用者信息
	development := zap.Development()                         // 开启文件及行号，方便调试
	filed := zap.Fields(zap.String("serviceName", "pledge")) // 设置初始化字段，记录服务名称

	// 构造全局日志对象
	// 使用上面配置的核心组件和增强选项
	Logger = zap.New(core, caller, development, filed)
}

// getCurrentAbPathByCaller 获取当前执行文件的绝对路径
// 用于确定日志文件的存储位置，支持go run模式下的路径解析
// 返回：
//   - string: 当前执行文件所在目录的绝对路径
func getCurrentAbPathByCaller() string {
	var abPath string
	_, filename, _, ok := runtime.Caller(0) // 获取调用栈信息
	if ok {
		abPath = path.Dir(filename) // 提取文件所在目录
	}
	return abPath
}
