package configs

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

type DataSource struct {
	Type string
	Dsn  string
}

type Config struct {
	DataSource DataSource
	Server     Server
	Jwt        Jwt
}

var EnvConfig *Config

func LoadConfig() *Config {

	path, err := os.Getwd() // get curent path
	if err != nil {
		panic(err)
	}

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(path + "/configs") // path to look for the config file in

	if err := viper.ReadInConfig(); err != nil { // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	config := &Config{}
	if err := viper.Unmarshal(config); err != nil {
		panic(err)
	}

	return config
}

func Init() {
	EnvConfig = LoadConfig()
}
