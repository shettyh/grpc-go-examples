package errorservice

import (
	"encoding/base64"
	"fmt"
	"github.com/golang/protobuf/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

// Error utils

// Error impl
func (e *Error) Error() string {
	return e.Message
}

// Errorf will return the Formatted error object
func Errorf(code codes.Code, userErrCode int64, temporary bool, msg string, args ...interface{}) error {
	return &Error{
		Code:          int64(code),
		Message:       fmt.Sprintf(msg, args...),
		UserErrorCode: userErrCode,
		Temporary:     temporary,
	}
}

// MarshalError will marshall the error and add to the context
func MarshalError(err error, ctx context.Context) error {
	rErr, ok := err.(*Error)
	if !ok {
		return err
	}

	pbErr, marshallerr := proto.Marshal(rErr)

	if marshallerr == nil {
		md := metadata.Pairs("rpc-error", base64.StdEncoding.EncodeToString(pbErr))
		_ = grpc.SetTrailer(ctx, md)
	}

	return status.Errorf(codes.Code(rErr.Code), rErr.Message)
}

// UnmarshalError will return the error object from metadata
func UnmarshalError(err error, md metadata.MD) *Error {
	vals, ok := md["rpc-error"]
	if !ok {
		return nil
	}

	buf, err := base64.StdEncoding.DecodeString(vals[0])

	if err != nil {
		return nil
	}

	var rerr Error
	if err := proto.Unmarshal(buf, &rerr); err != nil {
		return nil
	}

	return &rerr
}
