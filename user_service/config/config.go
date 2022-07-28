package config

type ServiceConfig struct {
	Name       string       `mapstructure:"name"`
	MysqlInfo  MysqlConfig  `mapstructure:"mysql_config"`
	ConsulInfo ConsulConfig `mapstructure:"consul_config"`
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

type FilePathConfig struct {
	ConfigFile string
	LogFile    string
}
