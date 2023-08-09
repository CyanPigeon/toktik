package discovery

import (
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/selector"
)

type Node interface {
	selector.Node
	PathPrefix() string
}

type node struct {
	selector.Node
	prefix string
}

func (n *node) PathPrefix() string {
	return n.prefix
}

func NewNode(scheme string, address string, service *registry.ServiceInstance) Node {
	n := &node{
		Node: selector.NewNode(scheme, address, service),
	}
	if prefix, ok := n.Metadata()["prefix"]; ok {
		n.prefix = prefix
	} else {
		n.prefix = "/"
	}
	return n
}
