package templateservice

import desc "github.com/ruko1202/balun.courses-microservice-6stream/chat-message/pkg/api/template_service"

type Implementation struct {
	desc.UnimplementedServiceServer
}

func NewImplementation() *Implementation {
	return &Implementation{}
}
