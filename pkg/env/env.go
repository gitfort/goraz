package env

import (
	parser "github.com/caarlos0/env"
	"github.com/joho/godotenv"
)

func LoadConfigs(configs interface{}, filenames ...string) error {
	_ = godotenv.Load(filenames...)
	if err := parser.Parse(configs); err != nil {
		return err
	}
	return nil
}
