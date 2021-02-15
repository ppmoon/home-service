package main

import (
	"fmt"
	"github.com/ppmoon/home-service/app/cli/cmd"
	"github.com/ppmoon/home-service/boot"
	"github.com/ppmoon/home-service/infrastructure/config"
	"github.com/spf13/viper"
)

func main() {
	boot.Boot()
	conf := config.GetConfig()
	fmt.Println(conf.SourceList)
	fmt.Println(viper.Get("source_list"))
	cmd.Execute()
}
