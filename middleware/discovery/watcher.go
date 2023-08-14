package discovery

import (
	"context"
	"errors"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/hashicorp/consul/api"
	"sync"
	"sync/atomic"
	"time"
)

type watcher struct {
	event   chan struct{}
	station *observatory
	ctx     context.Context
	cancel  context.CancelFunc
}

func (t *watcher) Next() (out []*registry.ServiceInstance, err error) {
	select {
	case <-t.ctx.Done():
		err = t.ctx.Err()
		return
	case <-t.event:
	}
	var ok bool
	if out, ok = t.station.load(); ok {
		return
	}
	return nil, errors.New("unexpected error: cannot get services list")
}

func (t *watcher) Stop() error {
	t.cancel()
	t.station.mutex.Lock()
	defer t.station.mutex.Unlock()
	delete(t.station.watchers, t)
	return nil
}

type observatory struct {
	watchers map[*watcher]struct{}
	services *atomic.Value
	filter   string
	mutex    sync.Mutex
	timeout  time.Duration
	count    int
}

func (t *observatory) newWatcher(ctx context.Context) *watcher {
	t.mutex.Lock()
	defer t.mutex.Unlock()
	p := &watcher{event: make(chan struct{}, 1), station: t}
	p.ctx, p.cancel = context.WithCancel(ctx)
	t.watchers[p] = struct{}{}
	if s, ok := t.load(); ok && len(s) > 0 {
		p.event <- struct{}{}
	}
	return p
}

func (t *observatory) load() (out []*registry.ServiceInstance, ok bool) {
	out, ok = t.services.Load().([]*registry.ServiceInstance)
	return
}

func (t *observatory) store(services []*registry.ServiceInstance) {
	t.mutex.Lock()
	defer t.mutex.Unlock()
	t.services.Store(services)
	t.count = len(services)
	for w := range t.watchers {
		w.event <- struct{}{}
	}
}

func (t *observatory) tick(ctx context.Context, client *api.Client) error {
	if t.timeout == 0 {
		t.timeout = time.Second * 5
	}

	services, err := client.Agent().ServicesWithFilter(t.filter)
	if err != nil {
		return err
	}

	if len(services) > 0 {
		t.store(resolver(ctx, services))
	}

	go func() {
		ticker := time.NewTicker(time.Second)
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				_service, _err := client.Agent().ServicesWithFilter(t.filter)
				if _err != nil {
					time.Sleep(time.Second)
					continue
				}
				if len(_service) != t.count {
					t.store(resolver(context.Background(), _service))
				}
			case <-ctx.Done():
				return
			}
		}
	}()
	return nil
}
