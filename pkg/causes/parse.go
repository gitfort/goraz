package causes

import (
	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"
)

func Parse(err error) (codes.Code, string) {
	switch errors.Cause(err) {
	case ErrAlreadyExists:
		return codes.AlreadyExists, err.Error()
	case ErrInvalidData:
		return codes.InvalidArgument, err.Error()
	case ErrNotFound:
		return codes.NotFound, err.Error()
	case ErrUnauthenticated:
		return codes.Unauthenticated, err.Error()
	case ErrPermissionDenied:
		return codes.PermissionDenied, err.Error()
	default:
		return codes.Unknown, err.Error()
	}
}
