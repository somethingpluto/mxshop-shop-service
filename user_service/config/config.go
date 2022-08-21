package config

// ServiceConfig
// @Description: 服务配置
//
type ServiceConfig struct {
	Name        string        `mapstructure:"name" json:"name"`
	MysqlInfo   MysqlConfig   `mapstructure:"mysql_config" json:"mysqlInfo"`
	ConsulInfo  ConsulConfig  `mapstructure:"consul_config"`
	RuntimeInfo RunTimeConfig `mapstructure:"runtime_config"`
}

// MysqlConfig
// @Description: 数据库配置
//
type MysqlConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Name     string `mapstructure:"name"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
}

type ConsulConfig struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}

type RunTimeConfig struct {
	Mode string `mapstructure:"mode"`
}

type FilePathConfig struct {
	ConfigFile string
	LogFile    string
}
