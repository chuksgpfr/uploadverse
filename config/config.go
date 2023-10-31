package config

import (
	"os"

	"github.com/spf13/viper"
)

type Configurations struct {
	PostgresDSN   string `mapstructure:"POSTGRES_DSN"`
	ServerAddress string `mapstructure:"SERVER_ADDRESS"`
	GinMode       string `mapstructure:"GIN_MODE"`
	IpfsNode      string `mapstructure:"IPFS_NODE"`
}

func LoadEnv() (config Configurations, err error) {
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	viper.AddConfigPath("./")
	viper.AddConfigPath("/app/config")
	viper.AutomaticEnv()

	if err = viper.ReadInConfig(); err != nil {
		config.PostgresDSN = os.Getenv("POSTGRES_DSN")
		config.ServerAddress = os.Getenv("SERVER_ADDRESS")
		config.GinMode = os.Getenv("GIN_MODE")
		config.IpfsNode = os.Getenv("IPFS_NODE")
		// return
	}
	err = viper.Unmarshal(&config)
	return
}
