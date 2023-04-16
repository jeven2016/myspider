package config

type Config struct {
	BindAddress string      `mapstructure:"bindAddress"`
	Execution   Execution   `mapstructure:"execution"`
	LogSetting  *LogSetting `mapstructure:"logSetting"`
}

type Execution map[string]WebSite

type WebSite struct {
	Enabled  bool        `mapstructure:"enabled"`
	Parallel bool        `mapstructure:"parallel"`
	Jobs     []JobParams `mapstructure:"jobs"`
}

type JobParams struct {
	Name        string `mapstructure:"name"`
	Url         string `mapstructure:"url"`
	Source      string `mapstructure:"source"`
	Destination string `mapstructure:"destination"`
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

func (c Config) Validate() error {
	return nil
}

func (c Config) Complete() error {
	return nil
}
