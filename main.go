package main

import "net/http"

type Router struct {
	Middleware []Handler
}

func (p *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c := &Context{Response: w, Request: r, Stack: p.Middleware}

	if r.URL.Path == "/hello" {
		p.HandleRequest(c, HelloHandler)
	}

	if r.URL.Path == "/hello_world" {
		p.HandleRequest(c, HelloHandler, WorldHandler)
	}
}

func (r *Router) HandleRequest(c *Context, handlers ...Handler) {
	s := c.Stack

	for _, h := range handlers {
		s = append(s, h)
	}

	c.Stack = s

	c.ProcessStack()
}

type Handler func(c *Context)

type Context struct {
	Response http.ResponseWriter
	Request  *http.Request
	Stack    []Handler
	Current  int
}

func (c *Context) ProcessStack() {
	if c.Current < len(c.Stack) {
		c.Stack[c.Current](c)
	}
}

func (c *Context) Continue() {
	c.Current++
	c.ProcessStack()
}

func HelloHandler(c *Context) {
	c.Response.Write([]byte("Hello"))
	c.Continue()
}

func WorldHandler(c *Context) {
	c.Response.Write([]byte(" World\n"))
	c.Continue()
}

func OtherHandler(c *Context) {
	c.Response.Write([]byte("I'm middleware\n"))
	c.Continue()
}

func InitializeRouter() *Router {
	return &Router{Middleware: []Handler{OtherHandler}}
}

func main() {
	r := InitializeRouter()

	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}
