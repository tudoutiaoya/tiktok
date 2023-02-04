package config

import (
	"fmt"
)

func InitConfig() *Configuration {
	// 初始化配置文件
	cfgFile := "./config/config.yaml"
	configuration, err := NewConfiguration(cfgFile)
	if err != nil {
		panic(fmt.Sprintf("配置初始化失败 %s", err.Error()))
	}
	return configuration
}
