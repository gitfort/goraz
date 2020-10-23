package discovery

import (
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"net"
	"net/http"
)

func NewServer(name string, tag string) *Server {
	srv := &Server{
		name:  name,
		tag:   tag,
		serve: make(map[PortType]ServeFunc),
	}
	srv.SetHealth(newHealthCheck())
	return srv
}

type Server struct {
	name  string
	tag   string
	serve map[PortType]ServeFunc
}

func (s *Server) SetHttp(handler http.Handler) *Server {
	s.serve[PortHTTP] = func(listener net.Listener) error {
		return http.Serve(listener, handler)
	}
	return s
}

func (s *Server) SetGrpc(server *grpc.Server) *Server {
	s.serve[PortGRPC] = func(listener net.Listener) error {
		return server.Serve(listener)
	}
	return s
}

func (s *Server) SetHealth(handler http.Handler) *Server {
	s.serve[PortHealth] = func(listener net.Listener) error {
		return http.Serve(listener, handler)
	}
	return s
}

func (s *Server) Run() {
	data := log.Fields{
		"name": s.name,
		"tag":  s.tag,
	}

	interrupt := make(chan error)
	for portType, serve := range s.serve {
		addr := Current(portType)
		data[string(portType)] = addr
		go func(addr string, serve ServeFunc) {
			interrupt <- listenAndServe(addr, serve)
		}(addr, serve)
	}

	log.WithFields(data).Info("service is running")
	if err := <-interrupt; err != nil {
		log.WithFields(data).Panic("service interrupted")
	} else {
		log.WithFields(data).WithError(err).Fatal("service interrupted")
	}
}
