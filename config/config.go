package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type EchoConfig struct {
	Db DatabaseConfig `mapstructure:"db"`
}

func (ec *EchoConfig) Init() {
	ec.Db.Init()
}

type DatabaseConfig struct {
	Name     string `mapstructure:"name"`
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
}

func (db *DatabaseConfig) Init() {
	viper.BindEnv("db.host", "DATABASE_HOST")
	viper.BindEnv("db.port", "DATABASE_PORT")
	viper.BindEnv("db.name", "DATABASE_NAME")
	viper.BindEnv("db.user", "DATABASE_USER")
	viper.BindEnv("db.password", "DATABASE_PASSWORD")

	if db.Host == "" {
		db.Host = "127.0.0.1"
	}
	if db.Name == "" {
		db.Name = "echodb"
	}
	if db.Port == "" {
		db.Port = "3306"
	}
	if db.User == "" {
		db.User = "user"
	}
	if db.Password == "" {
		db.Password = "password"
	}
}

func (db *DatabaseConfig) BuildConnectionString() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8",
		db.User, db.Password, db.Host, db.Port, db.Name)
}
