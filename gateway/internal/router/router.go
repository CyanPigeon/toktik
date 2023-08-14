package router

import (
	"github.com/CyanPigeon/toktik/gateway/internal/discovery"
	"net/http"
)

type Router interface {
	http.Handler
	Register(path string) error
	Apply(nodes []discovery.Node) error
}
