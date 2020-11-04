package lib

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type HandlerFunc func(c *Context)

type H map[string]interface{}

type Context struct {
	Writer     http.ResponseWriter
	R          *http.Request
	Path       string
	Method     string
	StatusCode int
	Params     map[string]string

	index    int
	handlers []HandlerFunc
}

func NewContext(w http.ResponseWriter, r *http.Request) *Context {
	return &Context{
		Writer:   w,
		R:        r,
		Path:     r.URL.Path,
		Method:   r.Method,
		index:    -1,
		handlers: []HandlerFunc{},
	}
}
func (c *Context) Status(code int) {
	c.StatusCode = code
	c.Writer.WriteHeader(code)
}

func (c *Context) GetQuery(key string) string {
	return c.R.URL.Query().Get(key)
}
func (c *Context) GetParam(key string) string {
	value, _ := c.Params[key]
	return value
}
func (c *Context) SetHeader(key string, value string) {
	c.Writer.Header().Set(key, value)
}

func (c *Context) Fail(code int, message string) {
	if c.StatusCode == code {
		log.Printf("%d Error "+message, code)
	}
}

func (c *Context) JSON(code int, js interface{}) {
	c.SetHeader("Content-Type", "application/json")
	c.Status(code)
	encoder := json.NewEncoder(c.Writer)
	if err := encoder.Encode(js); err != nil {
		http.Error(c.Writer, err.Error(), 500)
	}
}

func (c *Context) String(code int, format string, values ...interface{}) {
	c.SetHeader("Content-Type", "text/plain")
	c.Status(code)
	c.Writer.Write([]byte(fmt.Sprintf(format, values...)))
}

func (c *Context) HTML(code int, html string) {
	c.SetHeader("Content-Type", "text/html")
	c.Status(code)
	c.Writer.Write([]byte(html))
}

func (c *Context) AppendHandler(handlers ...HandlerFunc) {
	c.handlers = append(c.handlers, handlers...)
}

func (c *Context) Next() {
	c.index++
	s := len(c.handlers)
	for ; c.index < s; c.index++ {
		c.handlers[c.index](c)
	}
}
