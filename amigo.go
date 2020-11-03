package amigo

import (
	"net/http"
)

type HandlerFunc func(c *Context)

type Engine struct {
	*RouterGroup
	router *router
	group  []*RouterGroup
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
	engine.group = []*RouterGroup{engine.RouterGroup}
	return engine
}

func (g *RouterGroup) Group(prefix string) *RouterGroup {
	engine := g.engine
	result := &RouterGroup{
		prefix: g.prefix + prefix,
		parent: g,
		engine: engine,
	}
	engine.group = append(engine.group, result)
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

// implement handler interface
func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	e.router.handle(NewContext(w, r))
}

func (e *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, e)
}
