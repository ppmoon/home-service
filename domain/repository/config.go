package repository

import (
	"bytes"
	"errors"
	"github.com/BurntSushi/toml"
	"github.com/m1/go-generate-password/generator"
	"github.com/ppmoon/home-service/domain/aggregates"
	"github.com/ppmoon/home-service/domain/entity"
	"io/ioutil"
	"os"
	"sync"
)

// TODO abstract common method to infrastructure
const configFilePath = "./config.toml"

var (
	config *aggregates.Config
	once   sync.Once
)

type ConfigRepo struct {
}

func NewConfigRepo() Config {
	return &ConfigRepo{}
}

func (c *ConfigRepo) Get() (*aggregates.Config, error) {
	var err error
	if config != nil {
		return config, err
	}
	once.Do(func() {
		var isExist bool
		isExist, err = c.checkFileExist()
		if err != nil {
			return
		}
		if isExist {
			var data string
			data, err = c.loadConfigFile()
			if err != nil {
				return
			}
			// parse config file
			config, err = c.decodeConfig(data)
			if err != nil {
				return
			}
		} else {
			config, err = c.genConfig()
			if err != nil {
				return
			}
			// write to file
			var configBytes []byte
			configBytes, err = c.encodeConfig(config)
			if err != nil {
				return
			}
			err = c.createFile(configBytes)
			if err != nil {
				return
			}
		}
	})
	return config, err
}
func (c *ConfigRepo) genConfig() (conf *aggregates.Config, err error) {
	userCenterConfig, err := c.genUserCenterConfig()
	if err != nil {
		return nil, err
	}
	conf.UserCenterConfig = userCenterConfig
	return conf, err
}
func (c *ConfigRepo) genUserCenterConfig() (userCenterConfig *entity.UserCenterConfig, err error) {
	loginSalt, err := c.genSalt()
	if err != nil {
		return nil, err
	}
	userCenterConfig.LoginSalt = loginSalt
	return userCenterConfig, nil
}
func (c *ConfigRepo) genSalt() (string, error) {
	saltConfig := generator.Config{
		Length:                     generator.LengthVeryStrong,
		IncludeSymbols:             false,
		IncludeNumbers:             true,
		IncludeLowercaseLetters:    true,
		IncludeUppercaseLetters:    true,
		ExcludeSimilarCharacters:   true,
		ExcludeAmbiguousCharacters: true,
	}
	g, err := generator.New(&saltConfig)
	if err != nil {
		return "", err
	}
	pwd, err := g.Generate()
	if err != nil {
		return "", err
	}
	return *pwd, nil
}
func (c *ConfigRepo) decodeConfig(data string) (*aggregates.Config, error) {
	var conf aggregates.Config
	_, err := toml.Decode(data, &conf)
	if err != nil {
		return nil, err
	}
	return &conf, nil
}
func (c *ConfigRepo) encodeConfig(conf *aggregates.Config) ([]byte, error) {
	var configBuffer bytes.Buffer
	e := toml.NewEncoder(&configBuffer)
	err := e.Encode(conf)
	if err != nil {
		return []byte(""), err
	}
	return configBuffer.Bytes(), nil
}
func (c *ConfigRepo) checkFileExist() (isExist bool, err error) {
	fileInfo, err := os.Stat(configFilePath)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, err
	}
	if !fileInfo.IsDir() {
		return true, nil
	}
	// TODO notice absolute path
	err = errors.New("path is not a config file please check")
	return false, err
}
func (c *ConfigRepo) loadConfigFile() (data string, err error) {
	content, err := os.ReadFile(configFilePath)
	if err != nil {
		return "", err
	}
	return string(content), nil
}
func (c *ConfigRepo) createFile(data []byte) error {
	return ioutil.WriteFile(configFilePath, data, 0666)
}
