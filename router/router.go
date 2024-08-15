package router

import (
	"github.com/carepollo/noxt/ui"
)

type Router struct {
	routes *Tree[Handler]
}

type Route struct {
	Path      string
	Function  Handler
	Component Renderer
	Children  []Route
	RunBefore []Middleware
	RunAfter  []Middleware
}

type Handler func(Request) Response

type Renderer func(Request) ui.Component

type Middleware func(Request) Middleware

type Response struct {
	Status  int
	Headers map[string]string
	Body    []byte
}

type Request struct {
	Method      string
	QueryParams map[string]string
	PathParams  map[string]string
	Headers     map[string]string
	Body        []byte
}
