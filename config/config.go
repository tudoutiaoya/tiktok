package config

import (
	"fmt"
	"github.com/spf13/viper"
)

var v *viper.Viper

type Configuration struct {
	DatabaseSettings
	JWTSettings
}

// DatabaseSettings 数据库配置
type DatabaseSettings struct {
	DatabaseURI  string
	DatabaseName string
	Username     string
	Password     string
}

// JWT配置
type JWTSettings struct {
	SecretKey string
}

func NewConfiguration(configFile string) (configuration *Configuration, err error) {
	// 读取配置文件
	v = viper.GetViper()
	v.SetConfigType("yaml")
	v.SetConfigFile(configFile)
	if err = v.ReadInConfig(); err != nil {
		fmt.Printf("配置文件读取错误: %s", err)
		return nil, err
	}
	err = v.Unmarshal(&configuration)
	if err != nil {
		fmt.Printf("解析配置文件到结构体失败: %s", err)
		return nil, err
	}
	return configuration, err
}
