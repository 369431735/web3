package config

// Config 全局配置对象
// 保存应用程序的所有配置信息
var Config *Conf

// Conf 配置结构体
// 包含应用程序的所有配置项
type Conf struct {
	Mysql        MysqlConfig        // MySQL数据库配置
	Redis        RedisConfig        // Redis配置
	TestNet      TestNetConfig      // 测试网络配置
	MainNet      MainNetConfig      // 主网配置
	Token        TokenConfig        // 代币配置
	Email        EmailConfig        // 邮件配置
	DefaultAdmin DefaultAdminConfig // 默认管理员配置
	Threshold    ThresholdConfig    // 阈值配置
	Jwt          JwtConfig          // JWT令牌配置
	Env          EnvConfig          // 环境配置
}

// EnvConfig 环境配置
// 包含应用程序的运行环境、版本和网络设置
type EnvConfig struct {
	Port               string `toml:"port"`                 // 服务端口
	Version            string `toml:"version"`              // 应用版本
	Protocol           string `toml:"protocol"`             // 通信协议
	DomainName         string `toml:"domain_name"`          // 域名
	TaskDuration       int64  `toml:"task_duration"`        // 任务持续时间（秒）
	WssTimeoutDuration int64  `toml:"wss_timeout_duration"` // WebSocket超时时间（秒）
	TaskExtendDuration int64  `toml:"task_extend_duration"` // 任务延长时间（秒）
}

// ThresholdConfig 阈值配置
// 定义各种操作的阈值限制
type ThresholdConfig struct {
	PledgePoolTokenThresholdBnb string `toml:"pledge_pool_token_threshold_bnb"` // 质押池BNB代币阈值
}

// EmailConfig 邮件配置
// 定义发送系统邮件所需的参数
type EmailConfig struct {
	Username string   `toml:"username"` // 邮箱用户名
	Pwd      string   `toml:"pwd"`      // 邮箱密码
	Host     string   `toml:"host"`     // 邮件服务器主机
	Port     string   `toml:"port"`     // 邮件服务器端口
	From     string   `toml:"from"`     // 发件人地址
	Subject  string   `toml:"subject"`  // 邮件主题
	To       []string `toml:"to"`       // 收件人列表
	Cc       []string `toml:"cc"`       // 抄送人列表
}

// DefaultAdminConfig 默认管理员配置
// 系统初始管理员账号的用户名和密码
type DefaultAdminConfig struct {
	Username string `toml:"username"` // 管理员用户名
	Password string `toml:"password"` // 管理员密码
}

// JwtConfig JWT令牌配置
// 用于生成和验证JWT身份令牌
type JwtConfig struct {
	SecretKey  string `toml:"secret_key"`  // JWT密钥
	ExpireTime int    `toml:"expire_time"` // 过期时间（秒）
}

// TokenConfig 代币配置
// 定义系统中使用的代币相关设置
type TokenConfig struct {
	LogoUrl string `toml:"logo_url"` // 代币Logo图片URL
}

// MysqlConfig MySQL数据库配置
// 定义MySQL数据库连接和连接池参数
type MysqlConfig struct {
	Address      string `toml:"address"`        // 数据库服务器地址
	Port         string `toml:"port"`           // 数据库端口
	DbName       string `toml:"db_name"`        // 数据库名称
	UserName     string `toml:"user_name"`      // 数据库用户名
	Password     string `toml:"password"`       // 数据库密码
	MaxOpenConns int    `toml:"max_open_conns"` // 最大连接数
	MaxIdleConns int    `toml:"max_idle_conns"` // 最大空闲连接数
	MaxLifeTime  int    `toml:"max_life_time"`  // 连接最大生存时间
}

// TestNetConfig 测试网络配置
// 包含区块链测试网络的设置和合约地址
type TestNetConfig struct {
	ChainId              string `toml:"chain_id"`                // 区块链网络ID
	NetUrl               string `toml:"net_url"`                 // 网络RPC URL
	PlgrAddress          string `toml:"plgr_address"`            // PLGR代币合约地址
	PledgePoolToken      string `toml:"pledge_pool_token"`       // 质押池代币合约地址
	BscPledgeOracleToken string `toml:"bsc_pledge_oracle_token"` // BSC质押预言机合约地址
}

// MainNetConfig 主网配置
// 包含区块链主网络的设置和合约地址
type MainNetConfig struct {
	ChainId              string `toml:"chain_id"`                // 区块链网络ID
	NetUrl               string `toml:"net_url"`                 // 网络RPC URL
	PlgrAddress          string `toml:"plgr_address"`            // PLGR代币合约地址
	PledgePoolToken      string `toml:"pledge_pool_token"`       // 质押池代币合约地址
	BscPledgeOracleToken string `toml:"bsc_pledge_oracle_token"` // BSC质押预言机合约地址
}

// RedisConfig Redis配置
// 定义Redis连接和连接池参数
type RedisConfig struct {
	Address     string `toml:"address"`      // Redis服务器地址
	Port        string `toml:"port"`         // Redis端口
	Db          int    `toml:"db"`           // Redis数据库索引
	Password    string `toml:"password"`     // Redis密码
	MaxIdle     int    `toml:"max_idle"`     // 最大空闲连接数
	MaxActive   int    `toml:"max_active"`   // 最大活跃连接数
	IdleTimeout int    `toml:"idle_timeout"` // 空闲连接超时时间
}
