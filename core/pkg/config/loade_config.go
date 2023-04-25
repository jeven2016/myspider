package config

import (
	"core/pkg/common"
	"github.com/duke-git/lancet/v2/slice"
	"log"

	"github.com/spf13/viper"
)

var configMap = make(map[common.ConfigFileType]CfgHandler)

// LoadConfig 加载配置文件
func LoadConfig(setting *GenericConfigSetting) (CfgHandler, error) {
	// 如果没有在命令行参数中指定配置文件的具体路径，则从默认路径查找
	if len(setting.Path) == 0 {
		// 支持从以下目录查找配置文件： ./config/config.yaml,  /etc/sso-proxy/config.yaml
		viper.SetConfigName(setting.Name)
		viper.SetConfigType(setting.Format)
		slice.ForEach(setting.SearchPaths, func(index int, item string) {
			viper.AddConfigPath(item)
		})
	} else {
		// 从命令行指定的目录加载配置文件
		viper.SetConfigFile(setting.Path)
	}

	// 加载配置文件
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Failed to load config.yaml: %v", err)
		return nil, err
	}

	// 序列化成struct
	target := setting.Target
	if err := viper.Unmarshal(target); err != nil {
		log.Fatalf("Failed to unmarshal config.yaml: %v", err)
		return nil, err
	}
	configMap[setting.Type] = target
	return target, nil
}

// GetConfig 返回全局配置
func GetConfig(cfgFiletype common.ConfigFileType) CfgHandler {
	return configMap[cfgFiletype]
}

func LoadAllConfigFiles(cfgPath string) error {
	configObj := new(SysConfig)
	configFileSetting := &GenericConfigSetting{
		Type:        common.SysCfg,
		Path:        cfgPath,
		Name:        common.SysCfgName,
		Format:      common.CfgFormat,
		SearchPaths: common.CfgSearchPaths,
		Target:      configObj,
	}

	siteRule := new(SiteRule)
	siteConfigFileSetting := &GenericConfigSetting{
		Type:        common.SiteRuleCfg,
		Path:        cfgPath,
		Name:        common.SiteRuleName,
		Format:      common.CfgFormat,
		SearchPaths: common.CfgSearchPaths,
		Target:      siteRule,
	}

	cfgSlice := slice.ToSlice(configFileSetting, siteConfigFileSetting)

	var err error
	var cfg CfgHandler
	for _, s := range cfgSlice {

		if cfg, err = LoadConfig(s); err != nil {
			return err
		}

		if err = cfg.Validate(); err != nil {
			return err
		}

		if err = cfg.Complete(); err != nil {
			return err
		}
	}
	return nil
}
