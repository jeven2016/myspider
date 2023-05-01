package config

import "core/pkg/model"

type SysConfig struct {
	BindAddress string      `mapstructure:"bindAddress"`
	Execution   Execution   `mapstructure:"execution"`
	LogSetting  *LogSetting `mapstructure:"logSetting"`
	Redis       *Redis      `mapstructure:"redis"`
	MongoDb     *MongoDb    `mapstructure:"mongodb"`
	Spider      *Spider     `mapstructure:"spider"`
}

type Execution map[string]*model.Site

type Action struct {
	Load    bool `mapstructure:"load"`
	Save    bool `mapstructure:"save"`
	Analyse bool `mapstructure:"analyse"`
}

// LogSetting 日志相关配置
type LogSetting struct {
	LogLevel      string `mapstructure:"logLevel"`
	LogPath       string `mapstructure:"logPath"`
	OutputConsole bool   `mapstructure:"outputToConsole"`
	FileName      string `mapstructure:"fileName"`
	MaxSizeInMB   int    `mapstructure:"maxSizeInMB"`
	MaxAgeInDay   int    `mapstructure:"maxAgeInDay"`
	MaxBackups    int    `mapstructure:"maxBackups"`
	Compress      bool   `mapstructure:"compress"`
}
type Redis struct {
	Address      string `mapstructure:"address,omitempty"`
	Password     string `mapstructure:"password,omitempty"`
	Db           int    `mapstructure:"db,omitempty"`
	PoolSize     int    `mapstructure:"poolSize,omitempty"`
	PoolTimeout  int    `mapstructure:"poolTimeoutSeconds"`
	ReadTimeout  int    `mapstructure:"readTimeoutSeconds"`
	WriteTimeout int    `mapstructure:"writeTimeoutSeconds"`
}

type MongoDb struct {
	Uri      string `mapstructure:"uri"`
	Database string `mapstructure:"database"`
}

type Spider struct {
	HttpProxy string `mapstructure:"httpProxy"`
}

func (c *SysConfig) Validate() error {
	return nil
}

func (c *SysConfig) Complete() error {
	executionMap := c.Execution
	if executionMap != nil {
		for k, v := range executionMap {
			//set the name with the key if name isn't present
			if len(v.Name) == 0 {
				v.Name = k
				executionMap[k] = v
			}
		}
	}
	return nil
}
