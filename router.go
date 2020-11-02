package amigo

import "net/http"

type router struct {
	handler map[string]HandlerFunc
}

func newRouter() *router {
	return &router{handler: make(map[string]HandlerFunc)}
}

func (r *router) addRoute(method string, pattern string, handler HandlerFunc) {
	key := method + "-" + pattern
	r.handler[key] = handler
}
func (r *router) handle(c *Context) {
	key := c.Method + "-" + c.Path
	if handler, ok := r.handler[key]; ok {
		handler(c)
	} else {
		c.JSON(http.StatusNotFound, H{
			"message": "404 Not Found",
		})
	}
}
