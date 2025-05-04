package template_service

import (
	"context"

	"github.com/go-ozzo/ozzo-validation/v4"
	desc "github.com/ruko1202/balun.courses-microservice-6stream/pkg/api/template_service"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (i *Implementation) Ping(ctx context.Context, req *desc.PingRequest) (*desc.PingResponse, error) {
	_ = ctx
	
	err := validation.ValidateStruct(req,
		validation.Field(&req.Ping, validation.Required),
	)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	name := req.GetPing()

	return &desc.PingResponse{
		Pong:      name,
		Timestamp: timestamppb.Now(),
	}, nil
}
