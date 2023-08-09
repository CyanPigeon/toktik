package discovery

import (
	"context"
	"encoding/json"
	"github.com/CyanPigeon/toktik/gateway/internal/config"
	"github.com/go-kratos/kratos/v2/registry"
	etcd "go.etcd.io/etcd/client/v3"
)

type Discovery interface {
	GetService(ctx context.Context, filter ServiceFilter) ([]*registry.ServiceInstance, error)
	Watch(ctx context.Context, filter ServiceFilter) (registry.Watcher, error)
}

type discovery struct {
	namespace string
	client    *etcd.Client
	kv        etcd.KV
}

// New 用于创建一个服务发现实例
func New(client *etcd.Client) Discovery {
	return &discovery{
		namespace: config.Discovery.Namespace,
		client:    client,
		kv:        etcd.NewKV(client),
	}
}

// fetch 用于获取当前微服务列表
func fetch(ctx context.Context, kv etcd.KV, namespace string, filter ServiceFilter) (items []*registry.ServiceInstance, e error) {
	var response *etcd.GetResponse

	response, e = kv.Get(ctx, namespace, etcd.WithPrefix())
	if e != nil {
		return nil, e
	}

	items = make([]*registry.ServiceInstance, 0, len(response.Kvs))
	for _, elem := range response.Kvs {
		var item registry.ServiceInstance
		if e = json.Unmarshal(elem.Value, &item); e != nil {
			return nil, e
		}
		if !filter(&item) {
			items = append(items, &item)
		}
	}

	return
}

// GetService 用于根据指定的过滤器获取微服务列表
func (d *discovery) GetService(ctx context.Context, filter ServiceFilter) (items []*registry.ServiceInstance, e error) {
	return fetch(ctx, d.kv, d.namespace, filter)
}

// Watch 用于创建一个新的监听器
func (d *discovery) Watch(ctx context.Context, filter ServiceFilter) (registry.Watcher, error) {
	return newWatcher(ctx, d.namespace, d.client, filter)
}
