package interceptor

import (
	"context"
	"github.com/biryanim/auth/internal/metric"
	"google.golang.org/grpc"
)

func MetricsInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	metric.IncRequestCounter()

	return handler(ctx, req)
}
