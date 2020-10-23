package odata

import (
	"net/url"
	"strings"
)

func ParseUrl(url *url.URL) Query {
	query := make(Query)
	for _, key := range keys {
		if val := strings.TrimSpace(url.Query().Get(key)); val != "" {
			query[key] = val
		}
	}
	return query
}

func ParseString(str string) Query {
	values, _ := url.ParseQuery(str)
	query := make(Query)
	for _, key := range keys {
		if val := strings.TrimSpace(values.Get(key)); val != "" {
			query[key] = val
		}
	}
	return query
}
