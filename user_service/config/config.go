package config

// NacosConfig
// @Description: Nacos连接配置
//
type NacosConfig struct {
	Host      string `mapstructure:"host"`
	Port      int    `mapstructure:"port"`
	Namespace string `mapstructure:"namespace"`
	User      string `mapstructure:"user"`
	Password  string `mapstructure:"password"`
	Dataid    string `mapstructure:"dataid"`
	Group     string `mapstructure:"group"`
}

// FilePathConfig
// @Description: 文件路径配置
//
type FilePathConfig struct {
	ConfigFile string
	LogFile    string
}

// ServiceConfig
// @Description: 服务配置
//
type ServiceConfig struct {
	Name       string       `json:"name"`
	MysqlInfo  MysqlConfig  `json:"mysql"`
	ConsulInfo ConsulConfig `json:"consul"`
	RedisInfo  RedisConfig  `json:"redis"`
}

// MysqlConfig
// @Description: 数据库配置
//
type MysqlConfig struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Name     string `json:"name"`
	User     string `json:"user"`
	Password string `json:"password"`
}

// ConsulConfig
// @Description: consul配置
//
type ConsulConfig struct {
	Host string `json:"host"`
	Port int    `json:"port"`
}

// RedisConfig
// @Description: Redis配置
//
type RedisConfig struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Password string `json:"password"`
}
