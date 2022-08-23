package config

// NacosConfig
// @Description: nacos 连接配置
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
	Name       string       `mapstructure:"name" json:"name"`
	Mode       string       `json:"mode"`
	MySqlInfo  MySqlConfig  `mapstructure:"mysql_config" json:"mysql"`
	ConsulInfo ConsulConfig `json:"consul"`
}

// MySqlConfig
// @Description: Mysql连接信息
//
type MySqlConfig struct {
	Host     string `mapstructure:"host" json:"host"`
	Port     int    `mapstructure:"port" json:"port"`
	User     string `mapstructure:"user" json:"user"`
	Password string `mapstructure:"password" json:"password"`
	Name     string `mapstructure:"name" json:"name"`
}

type ConsulConfig struct {
	Host string `json:"host"`
	Port int    `json:"port"`
}
