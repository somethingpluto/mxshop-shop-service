package config

// FilePathConfig
// @Description: 文件路劲配置
//
type FilePathConfig struct {
	ConfigFile string
	LogFile    string
}

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

// ServiceConfig
// @Description: 服务配置
//
type ServiceConfig struct {
	Name         string `json:"name"`
	Host         string
	MysqlInfo    MysqlConfig    `json:"mysql"`
	ConsulInfo   ConsulConfig   `json:"consul"`
	JaegerInfo   JaegerConfig   `json:"jaeger"`
	RegisterInfo RegisterConfig `json:"register"`
}

// MysqlConfig
// @Description: Mysql配置
//
type MysqlConfig struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

// ConsulConfig
// @Description: consul配置
//
type ConsulConfig struct {
	Host string `json:"host"`
	Port int    `json:"port"`
}

type JaegerConfig struct {
	Host string `json:"host"`
	Port int    `json:"port"`
}

type RegisterConfig struct {
	Tags           []string `json:"tags"`
	CheckTimeOut   string   `json:"check_time_out"`
	CheckInterval  string   `json:"check_interval"`
	DeregisterTime string   `json:"deregister_time"`
}
