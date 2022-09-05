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
	Name             string                 `json:"name"`
	Mode             string                 `json:"mode"`
	Mysql            MysqlConfig            `json:"mysql"`
	Consul           ConsulConfig           `json:"consul"`
	Redis            RedisConfig            `json:"redis"`
	GoodsService     GoodsServiceConfig     `json:"goods_service"`
	InventoryService InventoryServiceConfig `json:"inventory_service"`
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

// RedisConfig
// @Description: Redis配置
//
type RedisConfig struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Password string `json:"password"`
	PoolSize int    `json:"poolSize"`
}

type GoodsServiceConfig struct {
	Name string `json:"name"`
}

type InventoryServiceConfig struct {
	Name string `json:"name"`
}
