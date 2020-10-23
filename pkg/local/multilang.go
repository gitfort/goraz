package local

import (
	"database/sql/driver"
	"encoding/json"
)

type MultiLang map[string]string

func (ml MultiLang) Value() (driver.Value, error) {
	valueString, err := json.Marshal(ml)
	return string(valueString), err
}

func (ml *MultiLang) Scan(value interface{}) error {
	if err := json.Unmarshal(value.([]byte), &ml); err != nil {
		return err
	}
	return nil
}

func (ml *MultiLang) Get(lang string) string {
	return (*ml)[lang]
}

func (ml *MultiLang) ToMap() map[string]string {
	return (*ml)
}

func ToMultiLang(input map[string]string) *MultiLang {
	result := MultiLang{}
	for k, v := range input {
		result[k] = v
	}
	return &result
}
