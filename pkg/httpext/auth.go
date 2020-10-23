package httpext

import (
	"github.com/gitfort/goraz/pkg/uuidext"
	"github.com/google/uuid"
	"net/http"
	"strings"
)

func GetToken(req *http.Request) string {
	return strings.ReplaceAll(req.Header.Get("Authorization"), "Bearer ", "")
}

func GetSessionID(req *http.Request) uuid.UUID {
	id, err := uuidext.FromString(req.Header.Get("X-SESSION-ID"))
	if err != nil {
		return uuid.Nil
	}
	return id
}
