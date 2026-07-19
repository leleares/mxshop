package user

import (
	"github.com/hashicorp/consul/api"
	"mxshop/app/pkg/options"
	"mxshop/app/user/srv/config"
	gapp "mxshop/gmicro/app"
	"mxshop/gmicro/registry"
	"mxshop/gmicro/registry/consul"
	"mxshop/pkg/app"
	"mxshop/pkg/log"
)

func NewApp(basename string) *app.App {
	cfg := config.New()
	// name 可以理解成服务名，basename 可以理解成cobra命令名
	appl := app.NewApp("user", basename, app.WithOptions(cfg), app.WithRunFunc(run(cfg)))
	return appl
}

func NewRegistrar(registryOpts *options.RegistryOptions) registry.Registrar {
	c := api.DefaultConfig()
	c.Address = registryOpts.Address
	c.Scheme = registryOpts.Scheme
	cli, err := api.NewClient(c)
	if err != nil {
		panic(err)
	}
	return consul.New(cli, consul.WithHealthCheck(true))
}

func NewUserApp(cfg *config.Config) (*gapp.App, error) {
	log.Init(cfg.Log)
	defer log.Flush()

	register := NewRegistrar(cfg.Registry)

	rpcServer, err := NewUserRPCServer(cfg)
	if err != nil {
		return nil, err
	}

	return gapp.New(
		gapp.WithName(cfg.Server.Name),
		gapp.WithRPCServer(rpcServer),
		gapp.WithRegistrar(register),
	), nil
}

func run(cfg *config.Config) app.RunFunc {
	return func(baseName string) error {
		userApp, err := NewUserApp(cfg)
		if err != nil {
			return err
		}

		if err := userApp.Run(); err != nil {
			log.Errorf("run user app error: %s", err)
			return err
		}
		return nil
	}
}
