package interceptor

import (
	"runtime/debug"
	"user-service/internal"

	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	logger, _ = internal.WireLogger()
)

func RecoveryHandler(p any) (err error) {
	logger.Error("recovered from panic", zap.String("trace", string(debug.Stack())))
	return status.Errorf(codes.Internal, "panic triggered: %v", p)
   }