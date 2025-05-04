package server

import (
	"context"
	"fmt"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpcrecovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/ruko1202/balun.courses-microservice-6stream/friend/internal/app/templateservice"
	"github.com/ruko1202/balun.courses-microservice-6stream/friend/internal/config"
	desc "github.com/ruko1202/balun.courses-microservice-6stream/friend/pkg/api/template_service"
)

func initGRPCServer(ctx context.Context, mux *runtime.ServeMux, grpcCfg config.Grpc) (*grpc.Server, error) {
	conn, err := grpc.NewClient(grpcCfg.BuildHostPort(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to dial grpc %s: %w", grpcCfg.BuildHostPort(), err)
	}

	grpcServer := grpc.NewServer(
		grpc.KeepaliveParams(grpcCfg.ServerParameters),
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_ctxtags.UnaryServerInterceptor(),
			grpcrecovery.UnaryServerInterceptor(),
		)),
	)

	// backend API
	desc.RegisterServiceServer(grpcServer, templateservice.NewImplementation())
	if err = desc.RegisterServiceHandler(ctx, mux, conn); err != nil {
		return nil, err
	}

	return grpcServer, nil
}
