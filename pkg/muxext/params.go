package muxext

import (
	"github.com/gitfort/goraz/api"
	"github.com/gitfort/goraz/pkg/causes"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"net/http"
)

func GetOnce(req *http.Request, name string) *api.Once {
	return &api.Once{
		Key: mux.Vars(req)[name],
	}
}

func GetMany(req *http.Request) *api.Many {
	return &api.Many{
		Query: req.URL.Query().Encode(),
	}
}

func GetID(req *http.Request, name string) (uuid.UUID, error) {
	id, err := uuid.Parse(mux.Vars(req)[name])
	if err != nil {
		return uuid.Nil, errors.Wrap(causes.ErrInvalidData, err.Error())
	}
	return id, nil
}
