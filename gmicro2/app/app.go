package app

import (
	registry "mxshop/gmicro2/registry"
)

type App struct {
	options Options
}

func NewApp(opts ...Option) *App {
	optss := &Options{}

	for _, f := range opts {
		f(optss)
	}

	return &App{
		options: *optss,
	}
}

// 微服务框架启动函数
func (a *App) Run() error {
	return nil
}

// 微服务框架退出函数
func (a *App) Stop() error {
	return nil
}

// 创建服务注册的结构体，仅本文件使用，不暴露
func (a *App) buildInstance() (*registry.ServiceInstance, error) {
	endpoints := make([]string, 0)
	for _, e := range a.options.Endpoints {
		endpoints = append(endpoints, e.String())
	}

	return &registry.ServiceInstance{
		ID:        a.options.ID,
		Name:      a.options.Name,
		Endpoints: a.options.Endpoints,
	}, nil
}
