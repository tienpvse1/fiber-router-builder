package route_builder

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type IRouter interface {
	Get(path string, handlers ...fiber.Handler) fiber.Router
	Post(path string, handlers ...fiber.Handler) fiber.Router
	Patch(path string, handlers ...fiber.Handler) fiber.Router
	Delete(path string, handlers ...fiber.Handler) fiber.Router
	Put(path string, handlers ...fiber.Handler) fiber.Router
	Options(path string, handlers ...fiber.Handler) fiber.Router
	Trace(path string, handlers ...fiber.Handler) fiber.Router
	Head(path string, handlers ...fiber.Handler) fiber.Router
	Connect(path string, handlers ...fiber.Handler) fiber.Router
}
type RouteBuilder struct {
	App       IRouter
	AutoBuild bool
}

type Route struct {
	path          string
	fiberInstance IRouter
	handlers      []func(c *fiber.Ctx) error
	method        string
	prefix        string
}

func (builder RouteBuilder) CreateRouteBuilder() *Route {
	return &Route{fiberInstance: builder.App}
}

func (route *Route) Get(path string) *Route {
	route.path = path
	route.method = "GET"
	return route
}

func (route *Route) Post(path string) *Route {
	route.path = path
	route.method = "POST"
	return route
}

func (route *Route) Patch(path string) *Route {
	route.path = path
	route.method = "PATCH"
	return route
}

func (route *Route) Put(path string) *Route {
	route.path = path
	route.method = "PUT"
	return route
}

func (route *Route) Delete(path string) *Route {
	route.path = path
	route.method = "DELETE"
	return route
}

func (route *Route) Options(path string) *Route {
	route.path = path
	route.method = "OPTIONS"
	return route
}

func (route *Route) Connect(path string) *Route {
	route.path = path
	route.method = "CONNECT"
	return route
}

func (route *Route) Trace(path string) *Route {
	route.path = path
	route.method = "TRACE"
	return route
}

func (route *Route) Head(path string) *Route {
	route.path = path
	route.method = "HEAD"
	return route
}

func (route *Route) SetPrefix(path string) {
	route.prefix = path
}

func (route *Route) Handler(handleFn func(c *fiber.Ctx) error) *Route {
	route.handlers = append(route.handlers, handleFn)
	fmt.Println("Handler called")
	return route.build()
}

func (route *Route) AddMiddlewares(middlewares ...fiber.Handler) *Route {
	route.handlers = middlewares
	return route
}

func (route *Route) build() *Route {
	if route.method == "GET" {
		route.fiberInstance.Get(route.prefix+route.path, route.handlers...)
	}
	if route.method == "POST" {
		fmt.Print(route.handlers)
		route.fiberInstance.Post(route.prefix+route.path, route.handlers...)
	}
	if route.method == "PUT" {
		route.fiberInstance.Put(route.prefix+route.path, route.handlers...)
	}
	if route.method == "PATCH" {
		route.fiberInstance.Patch(route.prefix+route.path, route.handlers...)
	}
	if route.method == "DELETE" {
		route.fiberInstance.Delete(route.prefix+route.path, route.handlers...)
	}
	if route.method == "OPTIONS" {
		route.fiberInstance.Options(route.prefix+route.path, route.handlers...)
	}
	if route.method == "HEAD" {
		route.fiberInstance.Head(route.prefix+route.path, route.handlers...)
	}
	if route.method == "TRACE" {
		route.fiberInstance.Trace(route.prefix+route.path, route.handlers...)
	}
	if route.method == "CONNECT" {
		route.fiberInstance.Connect(route.prefix+route.path, route.handlers...)
	}
	// reset all route value after a new route is built
	route.path = ""
	route.handlers = []func(c *fiber.Ctx) error{}
	route.method = ""
	return route
}
