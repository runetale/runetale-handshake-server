package middleware

import (
	"context"

	"github.com/runetale/runetale-handshake-server/domain/entity"
)

type GrpcMiddlewareImpl interface {
	Authenticate(context.Context) (context.Context, error)
}

type GrpcMiddleware struct {
}

func NewGrpcMiddleware(env *entity.Env) GrpcMiddlewareImpl {
	return &GrpcMiddleware{}
}

func (m *GrpcMiddleware) Authenticate(ctx context.Context) (context.Context, error) {
	return ctx, nil
}
