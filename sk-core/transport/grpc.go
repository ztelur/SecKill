package transport

import (
	"context"
	"github.com/go-kit/kit/transport/grpc"
	"SecKill/pb"
)

type grpcServer struct {
	calculate grpc.Handler
}

func (s *grpcServer) Calculate(ctx context.Context, r *pb.ArithmeticRequest) (*pb.ArithmeticResponse, error) {
	_, resp, err := s.calculate.ServeGRPC(ctx, r)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.ArithmeticResponse), nil
}

func NewGRPCServer(ctx context.Context, endpoints service.ArithmeticEndpoints) pb.ArithmeticServiceServer {
	return &grpcServer{
		calculate: grpc.NewServer(
			endpoints.CalculateEndpoint,
			service.DecodeGRPCArithmeticRequest,
			service.EncodeGRPCArithmeticResponse,
		),
	}
}
