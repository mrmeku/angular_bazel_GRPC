package server

import (
	"context"

	echoService "angular_bazel_example/api/echo_service"

	"github.com/golang/glog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

// Implementation of EchoServiceServer

type echoServer struct{}

func newEchoServer() echoService.EchoServiceServer {
	return new(echoServer)
}

func (s *echoServer) Echo(ctx context.Context, msg *echoService.SimpleMessage) (*echoService.SimpleMessage, error) {
	glog.Info(msg)
	return msg, nil
}

func (s *echoServer) EchoBody(ctx context.Context, msg *echoService.SimpleMessage) (*echoService.SimpleMessage, error) {
	glog.Info(msg)
	grpc.SendHeader(ctx, metadata.New(map[string]string{
		"foo": "foo1",
		"bar": "bar1",
	}))
	grpc.SetTrailer(ctx, metadata.New(map[string]string{
		"foo": "foo2",
		"bar": "bar2",
	}))
	return msg, nil
}

func (s *echoServer) EchoDelete(ctx context.Context, msg *echoService.SimpleMessage) (*echoService.SimpleMessage, error) {
	glog.Info(msg)
	return msg, nil
}
