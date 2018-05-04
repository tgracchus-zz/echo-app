package config

import (
	"flag"

	"github.com/spf13/viper"
)

func NewEchoConfig() (*EchoConfig, error) {
	configPath := flag.String("configPath", "config/", "Path to the configuration directory")
	env := flag.String("env", "local", "Environment")
	flag.Parse()

	viper.AddConfigPath(*configPath)
	viper.SetConfigType("yaml")
	viper.SetConfigName("echo-" + *env)
	viper.AutomaticEnv()

	var echoConfig = EchoConfig{}
	echoConfig.Init()

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	err = viper.Unmarshal(&echoConfig)
	if err != nil {
		return nil, err
	}

	return &echoConfig, nil

}
