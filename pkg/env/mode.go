package env

import "os"

type ModeName string

const (
	ModeDevelopment ModeName = "development"
	ModeProduction  ModeName = "production"
)

func Mode() ModeName {
	switch ModeName(os.Getenv("mode")) {
	case ModeProduction:
		return ModeProduction
	default:
		return ModeDevelopment
	}
}
