package openapi

import (
	"encoding/json"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"path"
)

var (
	marshal = map[string]func(in interface{}) ([]byte, error){
		".yml":  yaml.Marshal,
		".yaml": yaml.Marshal,
		".json": func(in interface{}) ([]byte, error) {
			return json.MarshalIndent(in, "", "  ")
		},
	}
)

func WriteFile(in interface{}, filename string) error {
	bytes, err := marshal[path.Ext(filename)](in)
	if err != nil {
		return err
	}
	if err := ioutil.WriteFile(filename, bytes, os.ModePerm); err != nil {
		return err
	}
	return nil
}
