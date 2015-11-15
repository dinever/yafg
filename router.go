package Yafg

import (
	"regexp"
	"strings"
)

const (
	RouterMethodGet = 1
	RouterMethodPost = 2
	RouterMethodPut = 3
	RouterMethodDelete = 4
)

type Router struct {
	routeSlice []*Route
}

type Handler func()

type Route struct {
	method int
	regex *regexp.Regexp
	params []string
	handler Handler
}

func NewRouter() *Router {
	router := new(Router)
	router.routeSlice = make([]*Route, 0)
	return router
}

func newRoute(method int, pattern string, handler Handler) *Route {
	route := new(Route)
	route.params = make([]string, 0)
	route.regex, route.params = route.parseURL(pattern)
	route.method = method
	route.handler = handler
	return route
}

func (router *Router) Get(pattern string, handler Handler) {
	route := newRoute(RouterMethodGet, pattern, handler)
	router.registerRoute(route, handler)
}

func (router *Router) registerRoute(route *Route, handler Handler) {
	router.routeSlice = append(router.routeSlice, route)
}

func (route *Route) parseURL(pattern string) (regex *regexp.Regexp, params []string) {
	params = make([]string, 0)
	segments := strings.Split(pattern, "/")
	for i, segment := range segments {
		if strings.HasPrefix(segment, ":") {
			segments[i] = `([\w-%]+)`
			params = append(params, strings.TrimPrefix(segment, ":"))
		}
	}
	regex, _ = regexp.Compile("^" + strings.Join(segments, "/") + "$")
	return regex, params
}
