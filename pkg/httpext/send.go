package httpext

import (
	"encoding/json"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net/http"
)

type Err struct {
	Message string `json:"error,omitempty" yaml:"error,omitempty"`
}

var (
	codeMap = map[codes.Code]int{
		codes.InvalidArgument:  http.StatusBadRequest,
		codes.AlreadyExists:    http.StatusConflict,
		codes.Unauthenticated:  http.StatusUnauthorized,
		codes.NotFound:         http.StatusNotFound,
		codes.PermissionDenied: http.StatusForbidden,
		codes.Unimplemented:    http.StatusNotImplemented,
		codes.Internal:         http.StatusInternalServerError,
		codes.Unknown:          http.StatusInternalServerError,
		codes.Unavailable:      http.StatusServiceUnavailable,
	}
)

func SendStatus(res http.ResponseWriter, req *http.Request, st *status.Status) {
	model := &Err{
		Message: st.Message(),
	}
	SendModel(res, req, codeMap[st.Code()], model)
}

func SendModel(res http.ResponseWriter, req *http.Request, code int, model interface{}) {
	bytes, _ := json.Marshal(model)
	SendData(res, req, code, "application/json", bytes)
}

func SendData(res http.ResponseWriter, req *http.Request, code int, mime string, data []byte) {
	res.Header().Set("Content-Type", mime)
	res.Header().Set("Accept-Charset","utf-8")
	res.WriteHeader(code)
	_, _ = res.Write(data)
}

func SendCode(res http.ResponseWriter, req *http.Request, code int) {
	res.WriteHeader(code)
}
