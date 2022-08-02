package config

type ServiceConfig struct {
	Name      string      `mapstructure:"name"`
	MySqlInfo MySqlConfig `mapstructure:"mysql_config"`
}

type MySqlConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Name     string `mapstructure:"name"`
}

type FilePathConfig struct {
	ConfigFile string
	LogFile    string
}
