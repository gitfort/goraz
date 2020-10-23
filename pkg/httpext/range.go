package httpext

import (
	"net/http"
	"strconv"
	"strings"
)

func GetRange(req *http.Request) []int {
	value := req.Header.Get("Range")
	if value == "" {
		return nil
	}
	var rng []int
	if strings.HasPrefix(value, "bytes") {
		value = strings.ReplaceAll(value, "bytes=", "")
		for _, item := range strings.Split(value, "-") {
			if item == "" {
				continue
			}
			val, _ := strconv.Atoi(item)
			rng = append(rng, val)
		}
	}
	return rng
}
