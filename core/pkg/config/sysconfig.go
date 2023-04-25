package config

type SysConfig struct {
	BindAddress string      `mapstructure:"bindAddress"`
	Execution   Execution   `mapstructure:"execution"`
	LogSetting  *LogSetting `mapstructure:"logSetting"`
	Redis       *Redis      `mapstructure:"redis"`
	MongoDb     *MongoDb    `mapstructure:"mongodb"`
}

type Execution map[string]WebSite

type WebSite struct {
	Enabled  bool        `mapstructure:"enabled"`
	Parallel bool        `mapstructure:"parallel"`
	Jobs     []JobParams `mapstructure:"jobs"`
}

type JobParams struct {
	Enabled     bool    `mapstructure:"enabled"`
	Name        string  `mapstructure:"name"`
	Parser      string  `mapstructure:"parser"`
	Url         string  `mapstructure:"url"`
	Source      string  `mapstructure:"source"`
	Destination string  `mapstructure:"destination"`
	Action      *Action `mapstructure:"action"`
}

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
	Uri      string `bson:"uri"`
	Database string `bson:"database"`
}

func (c *SysConfig) Validate() error {
	return nil
}

func (c *SysConfig) Complete() error {
	return nil
}
