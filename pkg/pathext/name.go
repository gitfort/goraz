package pathext

import (
	"path"
	"strings"
)

func Name(fileName string) string {
	return strings.TrimSuffix(fileName, path.Ext(fileName))
}
