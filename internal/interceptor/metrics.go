package interceptor

import (
	"context"

	"github.com/Arkosh744/auth-service-api/internal/metric"
	"google.golang.org/grpc"
)

func MetricsInterceptor(
	ctx context.Context,
	req any,
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (any, error) {
	res, err := handler(ctx, req)
	if err != nil {
		metric.IncRequestTotal(info.FullMethod, "error")
	} else {
		metric.IncRequestTotal(info.FullMethod, "success")
	}

	return res, err
}
