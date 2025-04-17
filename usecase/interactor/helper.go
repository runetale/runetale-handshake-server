package interactor

import (
	"context"

	"github.com/runetale/runetale-handshake-server/domain/entity"
	"github.com/runetale/runetale-handshake-server/utility"
	"google.golang.org/grpc/metadata"
)

func getNodeKeyFromCtx(ctx context.Context) (string, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", entity.ErrInvalidHeader
	}

	if _, found := md[utility.NodeKey]; !found {
		return "", entity.ErrInvalidHeader
	}

	return md[utility.NodeKey][0], nil
}

func getWgPubKeyFromCtx(ctx context.Context) (string, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", entity.ErrInvalidHeader
	}

	if _, found := md[utility.WgPubKey]; !found {
		return "", entity.ErrInvalidHeader
	}

	return md[utility.WgPubKey][0], nil
}

func getHostFromCtx(ctx context.Context) (string, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", entity.ErrInvalidHeader
	}

	if _, found := md[utility.HostName]; !found {
		return "", entity.ErrInvalidHeader
	}

	return md[utility.HostName][0], nil
}

func getOSFromCtx(ctx context.Context) (string, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", entity.ErrInvalidHeader
	}

	if _, found := md[utility.OS]; !found {
		return "", entity.ErrInvalidHeader
	}

	return md[utility.OS][0], nil
}

func getDistroFromCtx(ctx context.Context) (string, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", entity.ErrInvalidHeader
	}

	if _, found := md[utility.Distro]; !found {
		return "", entity.ErrInvalidHeader
	}

	return md[utility.Distro][0], nil
}
