package interceptor

import (
	"context"
	"fmt"

	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"go.uber.org/zap"
	"google.golang.org/protobuf/proto"
)

// a utility designed to bridge the gap between gRPC's logging system
// and the underlying logger typically provided by your application.
// It takes log.Logger instance as an argument, representing the
// existing logger, and returns a gRPC-compatible logging.Logger.
func InterceptorLogger(l *zap.Logger) logging.Logger {
	return logging.LoggerFunc(func(ctx context.Context, lvl logging.Level, msg string, fields ...any) {
		f := make([]zap.Field, 0, len(fields)/2)

		for i := 0; i < len(fields); i += 2 {
			key := fields[i]
			value := fields[i+1]

			switch v := value.(type) {
			case string:
				f = append(f, zap.String(key.(string), v))
			case int:
				f = append(f, zap.Int(key.(string), v))
			case bool:
				f = append(f, zap.Bool(key.(string), v))
			default:
				f = append(f, zap.Any(key.(string), v))
			}
		}

		logger := l.WithOptions(zap.AddCallerSkip(1)).With(f...)

		switch lvl {
		case logging.LevelDebug:
			logger.Debug(msg)
		case logging.LevelInfo:
			logger.Info(msg)
		case logging.LevelWarn:
			logger.Warn(msg)
		case logging.LevelError:
			logger.Error(msg)
		default:
			panic(fmt.Sprintf("unknown level %v", lvl))
		}
	})
}

func LogAdditionalFields(ctx context.Context, c interceptors.CallMeta) logging.Fields {
	var fields logging.Fields

	// Add payload and body size.
	if c.ReqOrNil != nil {
		p, ok := c.ReqOrNil.(proto.Message)
		if !ok {
			return fields
		}
		b, err := proto.Marshal(p) 
		if err != nil {
			return fields
		}
		fields = logging.Fields.AppendUnique(fields, logging.Fields{"payload", string(b)})
		fields = logging.Fields.AppendUnique(fields, logging.Fields{"bodySize", byteCountIEC(len(b))})
	}

	return fields
}

func byteCountIEC(b int) string {
	const unit = 1024
	if b < unit {
		return fmt.Sprintf("%d B", b)
	}
	div, exp := int(unit), 0
	for n := b / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %ciB",
		float64(b)/float64(div), "KMGTPE"[exp])
}