package config

import (
	"mxshop/pkg/common/cli/flag"
	"mxshop/pkg/log"
)

type Config struct {
	Log *log.Options `json:"log" mapstructure:"log"`
}

// 注册命令行参数至cobra，例如使得命令行支持：--log.level debug
func (c *Config) Flags() (fss flag.NamedFlagSets) {
	c.Log.AddFlags(fss.FlagSet("log"))
	return fss
}

// 检查命令行参数是否合法
func (c *Config) Validate() []error {
	var errors []error
	errors = append(errors, c.Log.Validate()...)
	return errors
}

func NewConfig() *Config {
	cfg := &Config{
		Log: log.NewOptions(),
	}
	return cfg
}
