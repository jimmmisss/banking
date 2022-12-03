package util

import "github.com/spf13/viper"

type DBConfig struct {
	User     string `mapstructure:"db_user"`
	Password string `mapstructure:"db_passwd"`
	Address  string `mapstructure:"db_addr"`
	Port     string `mapstructure:"db_port"`
	Name     string `mapstructure:"db_name"`
}

type Server struct {
	Address string `mapstructure:"server_address"`
	Port    string `mapstructure:"server_port"`
}

type Config struct {
	DB     DBConfig `mapstructure:"db"`
	Server Server   `mapstructure:"server"`
}

func LoadConfig() (Config, error) {
	var config Config

	viper.AddConfigPath("util")
	viper.SetConfigName("env")

	err := viper.ReadInConfig()
	if err != nil {
		return Config{}, nil
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		return Config{}, nil
	}

	return config, nil
}
