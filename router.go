package amigo

import (
	"github.com/odysa/amigo/lib"
	"strings"
)

type router struct {
	roots   map[string]*lib.Node
	handler map[string]HandlerFunc
}

func newRouter() *router {
	return &router{handler: make(map[string]HandlerFunc),
		roots: make(map[string]*lib.Node),
	}
}

func parsePattern(pattern string) []string {
	s := strings.Split(pattern, "/")
	parts := make([]string, 0)
	for _, item := range s {
		if item != "" {
			parts = append(parts, item)
			if item[0] == '*' {
				break
			}
		}
	}
	return parts
}

func (r *router) addRoute(method string, pattern string, handler HandlerFunc) {
	parts := parsePattern(pattern)
	key := method + "-" + pattern
	if _, ok := r.roots[method]; !ok {
		r.roots[method] = &lib.Node{}
	}
	r.roots[method].Insert(pattern, parts, 0)
	r.handler[key] = handler
}

func (r *router) getRoute(method string, path string) (*lib.Node, map[string]string) {
	searchParts := parsePattern(path)
	params := make(map[string]string)
	root, ok := r.roots[method]
	if !ok {
		return nil, nil
	}
	n := root.Find(searchParts, 0)
	if n != nil {
		parts := parsePattern(n.Pattern())
		for index, part := range parts {
			if part[0] == ':' {
				// find a param
				params[part[1:]] = searchParts[index]
			}
			if part[0] == '*' && len(part) > 1 {
				params[part[1:]] = strings.Join(searchParts[index:], "/")
				break
			}
		}
		return n, params
	}
	return nil, nil
}

func (r *router) handle(c *Context) {
	n, params := r.getRoute(c.Method, c.Path)
	if n != nil {
		c.Params = params
		key := c.Method + "-" + n.Pattern()
		if handler, ok := r.handler[key]; ok {
			c.AppendHandler(handler)
		}
	} else {
		c.AppendHandler(func(c *Context) {
			c.Fail(404, "Page Not Found")
		})
	}
	//call next handler
	c.Next()
}
