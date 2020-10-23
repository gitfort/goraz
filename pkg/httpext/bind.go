package httpext

import (
	"encoding/json"
	"github.com/gitfort/goraz/pkg/causes"
	"github.com/pkg/errors"
	"net/http"
)

func BindModel(req *http.Request, model interface{}) (err error) {
	defer func() {
		_ = req.Body.Close()
		if err != nil {
			err = errors.Wrap(causes.ErrInvalidData, err.Error())
		}
	}()
	return json.NewDecoder(req.Body).Decode(model)
}
