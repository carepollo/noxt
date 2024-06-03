package noxt

import (
	"fmt"
	"log"
	"net/http"

	"github.com/carepollo/noxt/router"
)

type Server struct {
	router *router.Router
	server http.Server
	logger log.Logger
}

func New(port int) Server {
	r := &router.Router{}
	return Server{
		server: http.Server{
			Addr:    ":" + fmt.Sprint(port),
			Handler: r,
		},
		router: r,
		logger: *log.Default(),
	}
}

func (server *Server) Run() error {
	return server.server.ListenAndServe()
}
