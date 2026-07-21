package app

import (
	"context"
	registry "mxshop/gmicro2/registry"
	"mxshop/pkg/log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/google/uuid"
)

type App struct {
	options  Options
	lk       sync.Mutex
	instance *registry.ServiceInstance
}

func NewApp(opts ...Option) *App {
	// 配置默认值
	optss := &Options{
		Signals: []os.Signal{
			syscall.SIGTERM,
			syscall.SIGQUIT,
			syscall.SIGINT,
		},
		RegisterTimeout: time.Second * 10,
		StopTimeout:     time.Second * 10,
	}

	uuid, err := uuid.NewUUID()
	if err != nil {
		optss.ID = uuid.String()
	}

	for _, f := range opts {
		f(optss)
	}

	return &App{
		options: *optss,
	}
}

// 微服务框架启动函数
func (a *App) Run() error {
	instance, err := a.buildInstance()
	if err != nil {
		return err
	}

	a.lk.Lock()
	a.instance = instance
	a.lk.Unlock()

	// 服务注册
	if a.options.Register != nil {
		// 需要进行超时约束
		ctx, cancel := context.WithTimeout(context.Background(), a.options.RegisterTimeout)
		defer cancel()
		err := a.options.Register.Register(ctx, instance)
		if err != nil {
			log.Errorf("register service error: %s", err)
			return err
		}
	}

	// 监听退出信号
	ch := make(chan os.Signal)
	signal.Notify(ch, a.options.Signals...)
	<-ch
	return nil
}

// 微服务框架退出函数
func (a *App) Stop() error {
	a.lk.Lock()
	instance := a.instance
	a.lk.Unlock()

	// 服务注销
	if a.options.Register != nil && instance != nil {
		// 需要进行超时约束
		ctx, cancel := context.WithTimeout(context.Background(), a.options.RegisterTimeout)
		defer cancel()
		err := a.options.Register.DeRegister(ctx, instance)
		if err != nil {
			log.Errorf("deregister service error: %s", err)
			return err
		}
	}

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
