package config

import (
	"fmt"
	"github.com/ppmoon/home-service/domain/entity"
	"github.com/ppmoon/home-service/infrastructure/log"
	"github.com/spf13/viper"
	"io/ioutil"
)

const YAML = "yaml"

var C entity.Config

func InitConfig(path string) {
	// init config source
	initConfig(path)
	// config log
	log.Init()

}

func initConfig(path string) {
	viper.SetConfigType(YAML)
	viper.AddConfigPath(path)
	files, err := ioutil.ReadDir(path)
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		viper.SetConfigName(file.Name())
		err = viper.MergeInConfig()
		if err != nil {
			panic(file.Name() + " " + err.Error())
		}
	}
	err = viper.ReadInConfig() // Find and read the config file
	if err != nil {            // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	err = viper.Unmarshal(&C)
	if err != nil { // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
}

// Get Config
func GetConfig() entity.Config {
	return C
}
