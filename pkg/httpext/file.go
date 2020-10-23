package httpext

import (
	"bytes"
	"github.com/gitfort/goraz/pkg/causes"
	"github.com/gitfort/goraz/pkg/pathext"
	"github.com/pkg/errors"
	"io"
	"net/http"
	"path"
)

type File struct {
	Name    string
	Ext     string
	Content []byte
}

func GetFile(req *http.Request, name string) (*File, error) {
	file, header, err := req.FormFile(name)
	if err != nil {
		return nil, errors.Wrap(causes.ErrInvalidData, err.Error())
	}
	defer file.Close()

	var buf bytes.Buffer
	if _, err := io.Copy(&buf, file); err != nil {
		return nil, errors.Wrap(causes.ErrInvalidData, err.Error())
	}

	return &File{
		Name:    pathext.Name(header.Filename),
		Ext:     path.Ext(header.Filename),
		Content: buf.Bytes(),
	}, nil
}
