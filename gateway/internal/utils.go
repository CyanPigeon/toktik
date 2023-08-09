package internal

import (
	"errors"
	"net/url"
	"path"
)

type OptionSetter[T any] func(t *T)

type Stack[T any] []*T

func (t *Stack[T]) Push(v *T) {
	*t = append(*t, v)
}
func (t *Stack[T]) Pop() *T {
	var v *T = nil
	if len(*t) > 0 {
		v = (*t)[len(*t)-1]
		*t = (*t)[:len(*t)-1]
	}
	return v
}

func CleanPath(p string) string {
	if p == "" {
		return "/"
	}
	if p[0] != '/' {
		p = "/" + p
	}

	condition := p[len(p)-1] == '/'
	p = path.Clean(p)

	if condition && p[len(p)-1] != '/' {
		p += "/"
	}
	return p
}

func SplitPath(p string, depth int) string {
	d := 0
	if p[0] != '/' {
		d += 1
	}
	for i, c := range p {
		if c != '/' {
			continue
		}
		d += 1
		if d > depth {
			return p[0:i]
		}
	}
	return p
}

func ParseEndpoint(endpoints []string, scheme string) (string, error) {
	for _, endpoint := range endpoints {
		u, err := url.Parse(endpoint)
		if err != nil {
			return "", err
		}
		if u.Scheme == scheme {
			return u.Host, nil
		}
	}
	return "", errors.New("specified scheme not matched")
}
