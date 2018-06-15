package main

import (
	"context"

	additionService "angular_bazel_GRPC/addition_service"

	"github.com/golang/glog"
)

// Implementation of AdditionServiceServer

type additionServer struct{}

func newAdditionServer() additionService.AdditionServiceServer {
	return new(additionServer)
}

func (s *additionServer) Sum(ctx context.Context, msg *additionService.SumRequest) (*additionService.SumResponse, error) {
	glog.Info(msg)
	sum := int32(0)

	for i := range msg.Integers {
		sum += msg.Integers[i]
	}

	response := &additionService.SumResponse{
		Sum: sum,
	}
	return response, nil
}
