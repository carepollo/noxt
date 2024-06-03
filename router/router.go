package router

import (
	"fmt"
	"net/http"
	"strings"
)

type Router struct {
	// routes Tree[Handler]
}

func (router *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := router.parseUrl(r.URL.Path)

	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, strings.Join(path, "/"))
	// method := r.Method

	// req := Request{
	// 	Method:      method,
	// 	QueryParams: make(map[string]string),
	// 	UrlParams:   make(map[string]string),
	// 	Headers:     make(map[string]string),
	// }
}

func (router *Router) parseUrl(url string) []string {
	cleaned := strings.Trim(url, "/")
	res := strings.Split(cleaned, "/")
	return res
}

type Handler func(Request) Response
