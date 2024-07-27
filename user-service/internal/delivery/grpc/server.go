package grpc

import (
	"user-service/internal/delivery/grpc/interceptor"
	"user-service/internal/usecase"

	pb "protobuf/user"

	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/recovery"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type GRPCServer struct {
	pb.UnimplementedUserServer
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
	pb.RegisterUserServer(s, &GRPCServer{uc: uc})
	return s
}