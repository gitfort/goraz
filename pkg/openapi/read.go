package openapi

import (
	"encoding/json"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"path"
)

var (
	unmarshal = map[string]func(in []byte, out interface{}) error{
		".yml":  yaml.Unmarshal,
		".yaml": yaml.Unmarshal,
		".json": json.Unmarshal,
	}
)

func ReadFile(filename string) (*Swagger, error) {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	swagger := new(Swagger)
	if err := unmarshal[path.Ext(filename)](bytes, swagger); err != nil {
		return nil, err
	}
	return swagger, nil
}
