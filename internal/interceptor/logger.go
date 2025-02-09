package interceptor

import (
	"context"
	"github.com/biryanim/platform_common/pkg/logger"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"time"
)

func LogInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	now := time.Now()

	res, err := handler(ctx, req)
	if err != nil {
		logger.Error(err.Error(), zap.String("method", info.FullMethod), zap.Any("request", req))
		return nil, err
	}

	logger.Info("request success", zap.String("method", info.FullMethod), zap.Any("response", res), zap.Duration("duration", time.Since(now)))

	return res, err
}
