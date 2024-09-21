package main

import (
	"fmt"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func main() {
	initViper()

	server := InitWebServer()
	err := server.Run(fmt.Sprintf("%s:%d", viper.GetString("web.host"), viper.GetInt("web.port")))
	if err != nil {
		panic(err)
	}
}

func initViper() {
	file := pflag.String("config", "./config/dev.yaml", "指定配置文件路径")
	pflag.Parse()
	viper.SetConfigFile(*file)

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}
