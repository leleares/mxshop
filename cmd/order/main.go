package main

import (
	"fmt"
	"mxshop/pkg/app"
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

func main() {
	cfg := &Config{
		Log: log.NewOptions(),
	}
	// name 可以理解成服务名，basename 可以理解成cobra命令名
	appl := app.NewApp("order", "mxshop-order", app.WithOptions(cfg), app.WithRunFunc(run))
	appl.Run()
}

func run(basename string) error {
	fmt.Printf("basename:%s start \n", basename)
	return nil
}
