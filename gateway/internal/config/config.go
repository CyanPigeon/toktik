package config

import (
	"fmt"
	"math"
	"time"
)

type Protocol int32

var Discovery = struct {
	Endpoints []string
	Namespace string
	TTL       time.Duration
	MaxRetry  int
}{
	Endpoints: []string{"localhost:2379"},
	Namespace: "/cyan-pigeon/toktik/api",
	TTL:       time.Second * 15,
	MaxRetry:  5,
}
var Endpoint = struct {
	Timeout  time.Duration
	MaxRetry int
}{
	Timeout:  time.Second * 15,
	MaxRetry: 5,
}
var Router = struct {
	StopApplyWhenError bool
	FollowRedirect     bool
}{
	StopApplyWhenError: false,
	FollowRedirect:     true,
}
var Transport = struct {
	Timeout                   time.Duration
	KeepAlive                 time.Duration
	MaxIdleConnections        int
	MaxIdleConnectionsPerHost int
	MaxConnectionsPerHost     int
	DisableCompression        bool
	IdleConnectionTimeout     time.Duration
	TLSHandshakeTimeout       time.Duration
	ExceptContinueTimeout     time.Duration
}{
	Timeout:                   time.Second * 15,
	KeepAlive:                 time.Second * 30,
	MaxIdleConnections:        10000,
	MaxIdleConnectionsPerHost: 1000,
	MaxConnectionsPerHost:     1000,
	DisableCompression:        true,
	IdleConnectionTimeout:     time.Second * 90,
	TLSHandshakeTimeout:       time.Second * 10,
	ExceptContinueTimeout:     time.Second * 1,
}
var Server = struct {
	Address           string
	Port              uint32
	MaxCurrentStream  uint32
	ReadTimeout       time.Duration
	ReadHeaderTimeout time.Duration
	WriteTimeout      time.Duration
	IdleTimeout       time.Duration
}{
	Address:           "0.0.0.0",
	Port:              8081,
	MaxCurrentStream:  math.MaxInt32,
	ReadTimeout:       time.Second * 15,
	ReadHeaderTimeout: time.Second * 10,
	WriteTimeout:      time.Second * 15,
	IdleTimeout:       time.Second * 120,
}

func ServerListenAddress() string {
	return fmt.Sprintf("%s:%d", Server.Address, Server.Port)
}

func init() {
}
