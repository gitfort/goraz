package causes

import (
	"context"
	"errors"
	"fmt"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const internalServer = "internal_server_error"

var (
	messages = map[codes.Code]string{
		codes.NotFound:         "%v_not_found",
		codes.InvalidArgument:  "%v_invalid_data",
		codes.AlreadyExists:    "%v_already_exists",
		codes.Unauthenticated:  "%v_unauthenticated",
		codes.PermissionDenied: "%v_permission_denied",
	}
)

type ToStatusFunc func(ctx context.Context, err error, module string) *status.Status

func ToStatus(translator Translator) ToStatusFunc {
	return func(ctx context.Context, err error, module string) *status.Status {
		st, ok := status.FromError(err)
		if ok {
			return st
		}
		switch code, msg := Parse(err); code {
		case codes.NotFound,
			codes.AlreadyExists,
			codes.InvalidArgument,
			codes.Unauthenticated,
			codes.PermissionDenied:
			log.WithError(errors.New(msg)).WithField("module", module).Debug("handled error")
			return status.New(code,
				translator.ByContext(ctx, fmt.Sprintf(messages[code], module)))
		default:
			log.WithError(errors.New(msg)).WithField("module", module).Error("unhandled error")
			return status.New(code, translator.ByContext(ctx, internalServer))
		}
	}
}
