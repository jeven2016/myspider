package config

type Config struct {
	BindAddress string      `mapstructure:"bindAddress"`
	LogSetting  *LogSetting `mapstructure:"logSetting"`
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
