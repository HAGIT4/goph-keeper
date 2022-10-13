package config

import (
	"github.com/spf13/viper"
)

type KeeperConfig struct {
	KeeperAddress string `mapstructure:"server_adddress"`
	DbString      string `mapstructure:"db_string"`
}

func InitKeeperConfig() (config *KeeperConfig, err error) {
	config = &KeeperConfig{}

	viper.SetConfigName("keeper")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	if err = viper.ReadInConfig(); err != nil {
		return nil, err
	}
	if err = viper.Unmarshal(config); err != nil {
		return nil, err
	}
	return config, nil
}
