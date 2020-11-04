package amigo

import (
	"net/http"
	"strings"
)

type HandlerFunc func(c *Context)

type Engine struct {
	*RouterGroup
	router *router
	groups []*RouterGroup
}
type RouterGroup struct {
	prefix      string
	middlewares []HandlerFunc
	parent      *RouterGroup
	engine      *Engine
}

func New() *Engine {
	engine := &Engine{router: newRouter()}
	engine.RouterGroup = &RouterGroup{engine: engine}
	engine.groups = []*RouterGroup{engine.RouterGroup}
	return engine
}

func (g *RouterGroup) Group(prefix string) *RouterGroup {
	engine := g.engine
	result := &RouterGroup{
		prefix: g.prefix + prefix,
		parent: g,
		engine: engine,
	}
	engine.groups = append(engine.groups, result)
	return result
}

func (g *RouterGroup) addRoute(method string, part string, handler HandlerFunc) {
	g.engine.router.addRoute(method, g.prefix+part, handler)
}

func (g *RouterGroup) Get(pattern string, handler HandlerFunc) {
	g.addRoute("GET", pattern, handler)
}

func (g *RouterGroup) Post(pattern string, handler HandlerFunc) {
	g.addRoute("POST", pattern, handler)
}

// Append middleware to group
func (g *RouterGroup) Add(wares ...HandlerFunc) {
	g.middlewares = append(g.middlewares, wares...)
}

// implement handler interface
func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c := NewContext(w, r)

	for _, group := range e.groups {
		if strings.HasPrefix(r.URL.Path, group.prefix) {
			c.handlers = append(c.handlers, group.middlewares...)
		}
	}

	e.router.handle(c)
}

func (e *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, e)
}
