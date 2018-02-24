package errorservice

import (
	"google.golang.org/grpc/codes"
	"fmt"
	"golang.org/x/net/context"
	"github.com/golang/protobuf/proto"
	"google.golang.org/grpc/metadata"
	"encoding/base64"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

func (e *Error) Error() string {
	return e.Message
}

func Errorf(code codes.Code, userErrCode int64, temporary bool, msg string, args ...interface{}) error {
	return &Error{
		Code:          int64(code),
		Message:       fmt.Sprintf(msg, args),
		UserErrorCode: userErrCode,
		Temporary:     temporary,
	}
}

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

func UnmarshalError(err error, md metadata.MD) *Error {
	vals, ok := md["rpc_error"]
	if ! ok {
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
