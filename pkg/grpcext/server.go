package grpcext

import (
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"google.golang.org/grpc"
)

func NewServer(opt ...grpc.ServerOption) *Server {
	s := new(Server)
	s.opts = append(s.opts, opt...)
	return s
}

type Server struct {
	opts []grpc.ServerOption
}

func (s *Server) WithMiddleware(mdl ...grpc.UnaryServerInterceptor) *Server {
	s.opts = append(s.opts, grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
		mdl...,
	)))
	return s
}

func (s *Server) Grpc() *grpc.Server {
	return grpc.NewServer(s.opts...)
}
