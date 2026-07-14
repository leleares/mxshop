package user

import (
	"fmt"
	"mxshop/app/user/config"
	"mxshop/pkg/app"
)

func NewApp(basename string) *app.App {
	cfg := config.NewConfig()
	// name 可以理解成服务名，basename 可以理解成cobra命令名
	appl := app.NewApp("order", basename, app.WithOptions(cfg), app.WithRunFunc(run(cfg)), app.WithNoConfig())
	return appl
}

func run(config *config.Config) app.RunFunc {
	return func(basename string) error {
		fmt.Printf("basename:%s start \n", basename)
		fmt.Printf("error level is %s", config.Log.Level)
		return nil
	}
}
