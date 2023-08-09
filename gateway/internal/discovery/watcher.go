package discovery

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/registry"
	etcd "go.etcd.io/etcd/client/v3"
	"time"
)

// ServiceFilter 表示微服务过滤器。true = 过滤，false = 允许
type ServiceFilter func(instance *registry.ServiceInstance) bool
type watcher struct {
	key       string
	ctx       context.Context
	cancel    context.CancelFunc
	client    *etcd.Client
	watcher   etcd.Watcher
	watchChan etcd.WatchChan
	kv        etcd.KV

	first bool

	discovery Discovery
	filter    ServiceFilter
}

// DefaultServiceFilter 是一个默认的微服务过滤器。该过滤器*允许*任何数据。
var DefaultServiceFilter ServiceFilter = func(_ *registry.ServiceInstance) bool {
	return false
}

// newWatcher 用于创建一个watcher
func newWatcher(ctx context.Context, namespace string, client *etcd.Client, filter ServiceFilter) (registry.Watcher, error) {
	w := &watcher{
		key:    namespace,
		client: client,
		kv:     etcd.NewKV(client),
		first:  true,
		filter: filter,
	}

	w.ctx, w.cancel = context.WithCancel(ctx)
	if err := w.reset(); err != nil {
		return nil, err
	}
	return w, nil
}

// reset 用于初始化/重置watcher
func (w *watcher) reset() error {
	if w.watcher != nil {
		_ = w.watcher.Close()
	}
	w.watcher = etcd.NewWatcher(w.client)
	w.watchChan = w.watcher.Watch(
		w.ctx, w.key,
		etcd.WithPrefix(),
		etcd.WithRev(0),
	)
	return w.watcher.RequestProgress(w.ctx)
}

// Next 通过阻塞的方式等待微服务状态更新，然后返回当前的微服务列表
func (w *watcher) Next() ([]*registry.ServiceInstance, error) {
	//if w.first {
	//	w.first = false
	//	return fetch(w.ctx, w.kv, w.key, w.filter)
	//}

	select {
	case <-w.ctx.Done():
		return nil, w.ctx.Err()
	case response, ok := <-w.watchChan:
		if !ok || response.Err() != nil {
			log.Warnf("something was wrong. response error = %v", response.Err())
			time.Sleep(time.Second)
			if err := w.reset(); err != nil {
				return nil, err
			}
			// 和官方代码是差不多的，但是并不通过名称来过滤微服务。
		}
		return fetch(w.ctx, w.kv, w.key, w.filter)
	}
}

// Stop 用于停止watcher
func (w *watcher) Stop() error {
	w.cancel()
	return w.watcher.Close()
}
