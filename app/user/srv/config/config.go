package config

import (
	"encoding/json"
	"mxshop/app/pkg/options"
	cliflag "mxshop/pkg/common/cli/flag"
	"mxshop/pkg/log"
)

type Config struct {
	MySQLOptions *options.MySQLOptions     `json:"mysql" mapstructure:"mysql"`
	Log          *log.Options              `json:"log" mapstructure:"log"`
	Server       *options.ServerOptions    `json:"server" mapstructure:"server"`
	Registry     *options.RegistryOptions  `json:"registry" mapstructure:"registry"`
	Telemetry    *options.TelemetryOptions `json:"telemetry" mapstructure:"telemetry"`
}

// 注册命令行参数至cobra，例如使得命令行支持：--log.level debug
func (c *Config) Flags() (fss cliflag.NamedFlagSets) {
	c.Log.AddFlags(fss.FlagSet("logs"))
	c.Server.AddFlags(fss.FlagSet("server"))
	c.Registry.AddFlags(fss.FlagSet("registry"))
	c.Telemetry.AddFlags(fss.FlagSet("telemetry"))
	c.MySQLOptions.AddFlags(fss.FlagSet("mysql"))
	return fss
}

// 检查命令行参数是否合法
func (c *Config) Validate() []error {
	var errors []error
	errors = append(errors, c.MySQLOptions.Validate()...)
	errors = append(errors, c.Log.Validate()...)
	errors = append(errors, c.Server.Validate()...)
	errors = append(errors, c.Registry.Validate()...)
	errors = append(errors, c.Telemetry.Validate()...)
	return errors
}

func (c *Config) String() string {
	data, _ := json.Marshal(c)
	return string(data)
}

func New() *Config {
	return &Config{
		MySQLOptions: options.NewMySQLOptions(),
		Log:          log.NewOptions(),
		Server:       options.NewServerOptions(),
		Registry:     options.NewRegistryOptions(),
		Telemetry:    options.NewTelemetryOptions(),
	}
}

func NewConfig() *Config {
	return New()
}
