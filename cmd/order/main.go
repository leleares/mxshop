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

func (c *Config) Flags() (fss flag.NamedFlagSets) {
	c.Log.AddFlags(fss.FlagSet("log"))
	return fss
}

func (c *Config) Validate() []error {
	var errors []error
	errors = append(errors, c.Log.Validate()...)
	return errors
}

func main() {
	cfg := &Config{
		Log: log.NewOptions(),
	}
	appl := app.NewApp("order", "mxshop", app.WithOptions(cfg), app.WithRunFunc(run))
	appl.Run()
}

func run(basename string) error {
	fmt.Printf("basename:%s start \n", basename)
	return nil
}
