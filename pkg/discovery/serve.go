package discovery

import (
	"fmt"
	"net"
	"os"
	"strings"
)

type PortType string

const (
	PortHTTP   PortType = "http"
	PortGRPC   PortType = "grpc"
	PortHealth PortType = "health"
)

var (
	defaultPorts = map[PortType]int{
		PortHTTP:   8080,
		PortGRPC:   50051,
		PortHealth: 2049,
	}
)

func Current(portType PortType) (addr string) {
	if addr := os.Getenv(strings.ToUpper(fmt.Sprint(portType, "_ADDR"))); addr != "" {
		return addr
	}
	return fmt.Sprint(":", defaultPorts[portType])
}

func Detect(service string, portType PortType) (addr string) {
	if addr := os.Getenv(strings.ToUpper(fmt.Sprint(service, "_", portType, "_ADDR"))); addr != "" {
		return addr
	}
	return fmt.Sprint(service, ":", defaultPorts[portType])
}

type ServeFunc func(listener net.Listener) error

func listenAndServe(addr string, serve ServeFunc) error {
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	defer func() {
		_ = listener.Close()
	}()
	return serve(listener)
}
