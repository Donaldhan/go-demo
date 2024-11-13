package evm

import (
	"github.com/spf13/viper"
	"log"
)

var config Config

func init() {
	loadConfig()
}

type Config struct {
	RpcUrl     string `mapstructure:"RpcUrl"`
	PrivateKey string `mapstructure:"PrivateKey"`
}

func loadConfig() {
	viper.SetConfigName("config") // name of config file (without extension)
	viper.SetConfigType("yaml")   // REQUIRED if the config file does not have the extension in the name
	//viper.AddConfigPath("/etc/appname/")  // path to look for the config file in
	viper.AddConfigPath(".") // optionally look for config in the working directory
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore error if desired
			log.Println("config file not found")
		} else {
			// Config file was found but another error was produced
			log.Fatalln("error reading config file", err)
		}
	}
	// 解析到结构体
	if err := viper.Unmarshal(&config); err != nil {
		log.Fatalf("解析配置到结构体失败: %v", err)
	}
	log.Println("config:", config)
}
