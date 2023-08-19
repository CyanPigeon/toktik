package registrar

import (
	"github.com/CyanPigeon/toktik/middleware/discovery"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/google/wire"
	"github.com/hashicorp/consul/api"
)

var ProviderSet = wire.NewSet(NewRegistry)

type RegistryConfig struct {
	Metadata  map[string]string
	Registrar registry.Registrar
}

func NewRegistry() (config *RegistryConfig, err error) {
	config = new(RegistryConfig)
	config.Metadata = map[string]string{
		// 当前微服务注册的终结点前缀。
		// 例如：对于user，这里应该设置为/douyin/user
		"prefix": "/douyin/user",
	}
	config.Registrar, err = discovery.New(api.DefaultConfig())
	return
}
