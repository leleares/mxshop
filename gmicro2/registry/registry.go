package registry

import "net/url"

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
