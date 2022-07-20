package gee

import (
	"fmt"
	"log"
	"net/http"
)

// HandlerFunc defines the request handler used by gee
type HandlerFunc func(ctx *Context)

type RouterGroup struct {
	prefix      string
	middlewares []HandlerFunc // support middleware
	parent      *RouterGroup  // support nesting
	engine      *Engine       // all groups share a Engine instance
}
type Engine struct {
	*RouterGroup
	router *router
	groups []*RouterGroup // store all groups
}

// Engine implement the interface of ServeHTTP

// New

func New() *Engine {
	engine := &Engine{router: newRouter()}
	engine.RouterGroup = &RouterGroup{engine: engine}
	engine.groups = []*RouterGroup{engine.RouterGroup}
	return engine
}

// Group 函数是用来创建一个新的 new group的
// 所有的group 共享一个engine
func (group *RouterGroup) Group(prefix string) *RouterGroup {
	engine := group.engine
	newGroup := &RouterGroup{
		engine: engine,
		prefix: group.prefix + prefix,
		parent: group,
	}
	engine.groups = append(engine.groups, newGroup)
	return newGroup
}

//
func (group *RouterGroup) addRoute(method string, comp string, handler HandlerFunc) {
	pattern := group.prefix + comp
	log.Printf("Route %4s - %s\n", method, pattern)
	group.engine.router.addRoute(method, pattern, handler)
}
func (group *RouterGroup) GET(pattern string, handler HandlerFunc) {
	group.addRoute("GET", pattern, handler)
}

func (group *RouterGroup) POST(pattern string, handler HandlerFunc) {
	group.addRoute("POST", pattern, handler)
}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	c := newContext(w, req)
	engine.router.handle(c)
	log.Printf("%v %d %v", c.Method, c.StatusCode, c.Path)
}

// Run defines the method to start a http server
func (engine *Engine) Run(addr string) (err error) {
	fmt.Println("http://localhost" + addr)
	return http.ListenAndServe(addr, engine)
}
