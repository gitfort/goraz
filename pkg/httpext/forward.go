package httpext

import (
	"net/http"
	"net/http/httputil"
)

func Forward(res http.ResponseWriter, req *http.Request, host string) (err error) {
	reverseProxy := &httputil.ReverseProxy{Director: func(r *http.Request) {
		r.Header.Add("X-Forwarded-Host", req.Host)
		r.Header.Add("X-Origin-Host", host)
		r.URL.Scheme = "http"
		r.URL.Host = host
	}, ErrorHandler: func(_ http.ResponseWriter, _ *http.Request, e error) {
		err = e
	}}
	reverseProxy.ServeHTTP(res, req)
	return
}
