package config

import (
	"github.com/spf13/viper"
)

type AgentConfig struct {
	KeeperAddress string `mapstructure:"keeper_address"`
	EncKey        string `mapstructure:"enc_key"`
}

func InitAgentConfig() (config *AgentConfig, err error) {
	config = &AgentConfig{}

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("$HOME/.keeper")
	viper.AddConfigPath("$HOME")

	if err = viper.ReadInConfig(); err != nil {
		return nil, err
	}
	if err = viper.Unmarshal(config); err != nil {
		return nil, err
	}
	return config, nil
}
