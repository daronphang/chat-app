package grpc

import (
	"notification-service/internal/delivery/grpc/interceptor"
	"notification-service/internal/usecase"

	pb "protobuf/proto/notification"

	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/recovery"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type GRPCServer struct {
	pb.UnimplementedNotificationServer
	uc *usecase.UseCaseService
}

func NewServer(logger *zap.Logger, uc *usecase.UseCaseService) *grpc.Server {
	// Configure interceptors.
	opts := []logging.Option{
		logging.WithLogOnEvents(logging.FinishCall),
		logging.WithTimestampFormat("2006/01/02 - 15:04:05"),
		logging.WithFieldsFromContextAndCallMeta(interceptor.LogAdditionalFields),
	}
	
	s := grpc.NewServer(grpc.ChainUnaryInterceptor(
		logging.UnaryServerInterceptor(interceptor.InterceptorLogger(logger), opts...),
		recovery.UnaryServerInterceptor(recovery.WithRecoveryHandler(interceptor.RecoveryHandler)),
	))
	pb.RegisterNotificationServer(s, &GRPCServer{uc: uc})
	return s
}