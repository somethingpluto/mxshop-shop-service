package config

// FilePathConfig
// @Description: 文件路劲配置
//
type FilePathConfig struct {
	ConfigFile string
	LogFile    string
}

type NacosConfig struct {
	Host      string `mapstructure:"host"`
	Port      int    `mapstructure:"port"`
	Namespace string `mapstructure:"namespace"`
	User      string `mapstructure:"user"`
	Password  string `mapstructure:"password"`
	Dataid    string `mapstructure:"dataid"`
	Group     string `mapstructure:"group"`
}

type ServiceConfig struct {
	Name   string       `json:"name"`
	Mode   string       `json:"mode"`
	Mysql  MysqlConfig  `json:"mysql"`
	Consul ConsulConfig `json:"consul"`
}

type MysqlConfig struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

type ConsulConfig struct {
	Host string `json:"host"`
	Port int    `json:"port"`
}
