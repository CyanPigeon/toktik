package http

import (
	"context"
	"errors"
	"github.com/CyanPigeon/toktik/gateway/internal/config"
	"io"
	"net"
	"net/http"
	"sync"
)

type Transport interface {
	http.RoundTripper
	io.Closer
}

var globalHttpClient = newClient()

func newClient() *http.Client {
	return &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			if len(via) >= 10 {
				return errors.New("too many redirect(over 10 times)")
			}
			return nil
		},
		Transport: &http.Transport{
			Proxy: http.ProxyFromEnvironment,
			DialContext: (&net.Dialer{
				Timeout:   config.Transport.Timeout,
				KeepAlive: config.Transport.KeepAlive,
			}).DialContext,
			MaxIdleConns:          config.Transport.MaxIdleConnections,
			MaxIdleConnsPerHost:   config.Transport.MaxIdleConnectionsPerHost,
			MaxConnsPerHost:       config.Transport.MaxConnectionsPerHost,
			DisableCompression:    config.Transport.DisableCompression,
			IdleConnTimeout:       config.Transport.IdleConnectionTimeout,
			TLSHandshakeTimeout:   config.Transport.TLSHandshakeTimeout,
			ExpectContinueTimeout: config.Transport.ExceptContinueTimeout,
		},
	}
}

type transport struct {
	queue map[*http.Request]context.CancelFunc
	mutex sync.Mutex
}

func NewTransport() Transport {
	return &transport{
		queue: make(map[*http.Request]context.CancelFunc),
		mutex: sync.Mutex{},
	}
}

func (t *transport) RoundTrip(request *http.Request) (response *http.Response, err error) {
	t.mutex.Lock()
	ctx := request.Context()
	if ctx == nil {
		ctx = context.Background()
	}
	ctx, cancel := context.WithCancel(ctx)
	request.WithContext(ctx)
	t.queue[request] = cancel
	t.mutex.Unlock()
	defer func() {
		t.mutex.Lock()
		delete(t.queue, request)
		t.mutex.Unlock()
	}()
	request.RequestURI = ""
	response, err = globalHttpClient.Do(request)

	return
}

func (t *transport) Close() error {
	t.mutex.Lock()
	for _, cancel := range t.queue {
		cancel()
	}
	return nil
}
