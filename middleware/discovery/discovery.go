package discovery

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/hashicorp/consul/api"
	"math/rand"
	"net"
	"net/url"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

type Discovery interface {
	GetService(ctx context.Context, filter string) ([]*registry.ServiceInstance, error)
	Watch(ctx context.Context, filter string) (registry.Watcher, error)
}

type option struct {
	*api.Config
	healthCheckInterval            int
	heartbeat                      bool
	deregisterCriticalServiceAfter int
}

type Registry struct {
	ctx           context.Context
	cancel        context.CancelFunc
	option        *option
	client        *api.Client
	mutex         *sync.RWMutex
	serviceChecks api.AgentServiceChecks
	observatories map[string]*observatory
}

func New(config *api.Config) (_ *Registry, err error) {
	r := &Registry{
		option: &option{
			Config:                         config,
			healthCheckInterval:            10,
			heartbeat:                      true,
			deregisterCriticalServiceAfter: 600,
		},
		mutex:         &sync.RWMutex{},
		observatories: make(map[string]*observatory),
	}
	r.client, err = api.NewClient(config)
	r.ctx, r.cancel = context.WithCancel(context.Background())
	if err != nil {
		return nil, err
	}
	return r, nil
}

func (t *Registry) GetService(ctx context.Context, filter string) (out []*registry.ServiceInstance, err error) {
	t.mutex.RLock()
	defer t.mutex.RUnlock()

	w, ok := t.observatories[filter]

	if !ok {
		if out = t.fetch(ctx, filter); len(out) > 0 {
			return
		}
		return nil, fmt.Errorf("filter '%s' do not match any services in registry", filter)
	}

	if out, ok = w.load(); ok && out != nil {
		return
	}
	if out = t.fetch(ctx, filter); len(out) > 0 {
		return
	}

	return nil, fmt.Errorf("filter '%s' do not match any services in registry", filter)
}

func (t *Registry) Watch(ctx context.Context, filter string) (registry.Watcher, error) {
	t.mutex.Lock()
	defer t.mutex.Unlock()

	o, ok := t.observatories[filter]
	if !ok {
		o = &observatory{
			watchers: make(map[*watcher]struct{}),
			services: &atomic.Value{},
			filter:   filter,
		}
		t.observatories[filter] = o
	}
	w := o.newWatcher(ctx)

	if !ok {
		err := o.tick(ctx, t.client)
		if err != nil {
			return nil, err
		}
	}
	return w, nil
}

func (t *Registry) Register(_ context.Context, service *registry.ServiceInstance) error {
	addresses := make(map[string]api.ServiceAddress, len(service.Endpoints))
	checkAddresses := make([]string, 0, len(service.Endpoints))
	for _, endpoint := range service.Endpoints {
		raw, err := url.Parse(endpoint)
		if err != nil {
			return err
		}
		addr := raw.Hostname()
		port, _ := strconv.ParseUint(raw.Port(), 10, 16)

		checkAddresses = append(checkAddresses, net.JoinHostPort(addr, strconv.FormatUint(port, 10)))
		addresses[raw.Scheme] = api.ServiceAddress{Address: endpoint, Port: int(port)}
	}

	asr := &api.AgentServiceRegistration{
		ID:   service.ID,
		Name: service.Name,
		Meta: service.Metadata,
		Tags: []string{
			fmt.Sprintf("version=%s", service.Version),
			"team=CyanPigeon", // <-- 加了个tag
		},
		TaggedAddresses: addresses,
	}
	if len(checkAddresses) > 0 {
		host, portRaw, _ := net.SplitHostPort(checkAddresses[0])
		port, _ := strconv.ParseInt(portRaw, 10, 32)
		asr.Address = host
		asr.Port = int(port)
	}

	for _, address := range checkAddresses {
		asr.Checks = append(asr.Checks, &api.AgentServiceCheck{
			TCP:                            address,
			Interval:                       fmt.Sprintf("%ds", t.option.healthCheckInterval),
			DeregisterCriticalServiceAfter: fmt.Sprintf("%ds", t.option.deregisterCriticalServiceAfter),
			Timeout:                        "5s",
		})
	}
	// custom checks
	asr.Checks = append(asr.Checks, t.serviceChecks...)

	if t.option.heartbeat {
		asr.Checks = append(asr.Checks, &api.AgentServiceCheck{
			CheckID:                        "service:" + service.ID,
			TTL:                            fmt.Sprintf("%ds", t.option.healthCheckInterval*2),
			DeregisterCriticalServiceAfter: fmt.Sprintf("%ds", t.option.deregisterCriticalServiceAfter),
		})
	}

	err := t.client.Agent().ServiceRegister(asr)
	if err != nil {
		return err
	}
	if t.option.heartbeat {
		go func() {
			time.Sleep(time.Second)
			err = t.client.Agent().UpdateTTL("service:"+service.ID, "pass", "pass")
			if err != nil {
				log.Errorf("[Consul]update ttl heartbeat to consul failed!err:=%v", err)
			}
			ticker := time.NewTicker(time.Second * time.Duration(t.option.healthCheckInterval))
			defer ticker.Stop()
			for {
				select {
				case <-t.ctx.Done():
					_ = t.client.Agent().ServiceDeregister(service.ID)
					return
				default:
				}
				select {
				case <-t.ctx.Done():
					_ = t.client.Agent().ServiceDeregister(service.ID)
					return
				case <-ticker.C:
					// ensure that unregistered services will not be re-registered by mistake
					if errors.Is(t.ctx.Err(), context.Canceled) || errors.Is(t.ctx.Err(), context.DeadlineExceeded) {
						_ = t.client.Agent().ServiceDeregister(service.ID)
						return
					}
					err = t.client.Agent().UpdateTTL("service:"+service.ID, "pass", "pass")
					if err != nil {
						log.Errorf("[Consul] update ttl heartbeat to consul failed! err=%v", err)
						// when the previous report fails, try to re register the service
						time.Sleep(time.Duration(rand.Intn(5)) * time.Second)
						if err := t.client.Agent().ServiceRegister(asr); err != nil {
							log.Errorf("[Consul] re registry service failed!, err=%v", err)
						} else {
							log.Warn("[Consul] re registry of service occurred success")
						}
					}
				}
			}
		}()
	}
	return nil
}

func (t *Registry) Deregister(_ context.Context, service *registry.ServiceInstance) error {
	defer t.cancel()
	return t.client.Agent().ServiceDeregister(service.ID)
}

func (t *Registry) fetch(ctx context.Context, filter string) []*registry.ServiceInstance {
	s, e := t.client.Agent().ServicesWithFilter(filter)
	if e == nil && len(s) > 0 {
		return resolver(ctx, s)
	}
	return nil
}

func resolver(_ context.Context, services map[string]*api.AgentService) (out []*registry.ServiceInstance) {
	out = make([]*registry.ServiceInstance, 0, len(services))
	for _, service := range services {
		var version string
		for _, tag := range service.Tags {
			s := strings.SplitN(tag, "=", 2)
			if len(s) == 2 && s[0] == "version" {
				version = s[1]
			}
		}

		endpoints := make([]string, 0, 2)
		for scheme, address := range service.TaggedAddresses {
			if strings.HasSuffix(scheme, "ipv4") || strings.HasSuffix(scheme, "ipv6") {
				continue
			}
			endpoints = append(endpoints, address.Address)
		}

		if len(endpoints) > 0 && service.Address != "" && service.Port != 0 {
			endpoints = append(endpoints, fmt.Sprintf("http://%s:%d", service.Address, service.Port))
		}

		out = append(out, &registry.ServiceInstance{
			ID:        service.ID,
			Name:      service.Service,
			Metadata:  service.Meta,
			Version:   version,
			Endpoints: endpoints,
		})
	}
	return
}
