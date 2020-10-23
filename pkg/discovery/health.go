package discovery

import "net/http"

func newHealthCheck() *healthCheck {
	return &healthCheck{}
}

type healthCheck struct {
}

func (*healthCheck) ServeHTTP(res http.ResponseWriter, _ *http.Request) {
	res.WriteHeader(http.StatusOK)
}
