package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

var config Config

func init() {
	vi := viper.New()
	vi.AddConfigPath(".")
	vi.SetConfigName("config")
	vi.SetConfigType("yaml")

	if err := vi.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Println("[Warn] Config Not Found.")
		} else {
			fmt.Println("[Error] Config Read Error: ", err)
			os.Exit(-1)
		}
	}
	err := vi.Unmarshal(&config)
	if err != nil {
		fmt.Println("[Error] config unmarshal failed. error: ", err)
		os.Exit(-1)
	}
}

func GetConfig() Config {
	return config
}
