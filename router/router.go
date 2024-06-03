package router

type Router struct {
	routes Tree[Handler]
}

type Handler func(Request) Response
