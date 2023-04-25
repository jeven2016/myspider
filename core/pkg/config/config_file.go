package config

import "core/pkg/common"

type CfgHandler interface {
	Validate() error
	Complete() error
}

type GenericConfigSetting struct {
	Type        common.ConfigFileType
	Path        string
	Name        string
	Format      string
	SearchPaths []string
	Target      CfgHandler
}
