package registry

import (
	"context"
	"net/url"
)

// 微服务注册和注销接口
type RegisterInterface interface {
	Register(ctx context.Context, instance *ServiceInstance) error
	DeRegister(ctx context.Context, instance *ServiceInstance) error
}

// 服务发现
type DiscoveryInterface interface {
	// 获取具体微服务
	GetService(ctx context.Context, serviceName string) (*ServiceInstance, error)
	// 服务监听器
	Watch(ctx context.Context, serviceName string) (*WatcherInterface, error)
}

type WatcherInterface interface {
	// 	获取服务实例，next 在下面的情況下会返回服务
	// 1. 第一次监听时，如果服务实例列表不为空，则这回服务实例列表
	// 2. 如果服务实例发生变化，则返回服务实例列表
	// 3. 如果上面两种情况都不满足，则会阻塞到 context deadline 或者 cancel
	Next() ([]*ServiceInstance, error)
	// 主动放弃监听
	Stop() error
}

// 微服务基本信息
type ServiceInstance struct {
	// 注册到服务中心的服务id
	ID string `json:"id"`

	// 服务名称
	Name string `json:"name"`

	// 服务版本
	Version string `json:"version"`

	// 服务元信息
	Metadata map[string]string `json:"metadata"`

	// 服务访问 url
	Endpoints []url.URL `json:"endpoints"`
}
