package config

import (
	"log"

	"github.com/spf13/viper"
)

var config = new(Config)

// LoadConfig 加载配置文件
func LoadConfig(configPath string) error {
	// 如果没有在命令行参数中指定配置文件的具体路径，则从默认路径查找
	if len(configPath) == 0 {
		// 支持从以下目录查找配置文件： ./config/config.yaml,  /etc/sso-proxy/config.yaml
		viper.SetConfigName("config")
		viper.SetConfigType("yaml")
		viper.AddConfigPath("./conf")
		viper.AddConfigPath("/etc/myspider/")
	} else {
		// 从命令行指定的目录加载配置文件
		viper.SetConfigFile(configPath)
	}

	// 加载配置文件
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Failed to load config.yaml: %v", err)
		return err
	}

	// 序列化成struct
	if err := viper.Unmarshal(config); err != nil {
		log.Fatalf("Failed to unmarshal config.yaml: %v", err)
		return err
	}
	return nil
}

// GetConfig 返回全局配置
func GetConfig() *Config {
	return config
}
