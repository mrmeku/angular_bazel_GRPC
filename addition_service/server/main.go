package main

import (
	"context"
	"net"

	additionService "angular_bazel_GRPC/addition_service"

	"github.com/golang/glog"
	"google.golang.org/grpc"

	"flag"
)

var (
	addr    = flag.String("addr", ":9090", "endpoint of the gRPC service")
	network = flag.String("network", "tcp", "a valid network type which is consistent to -addr")
)

// Run starts the example gRPC service.
// "network" and "address" are passed to net.Listen.
func Run(ctx context.Context, network, address string) error {
	l, err := net.Listen(network, address)
	if err != nil {
		return err
	}
	defer func() {
		if err := l.Close(); err != nil {
			glog.Errorf("Failed to close %s %s: %v", network, address, err)
		}
	}()

	s := grpc.NewServer()
	additionService.RegisterAdditionServiceServer(s, newAdditionServer())

	go func() {
		defer s.GracefulStop()
		<-ctx.Done()
	}()
	return s.Serve(l)
}

func main() {
	flag.Parse()
	defer glog.Flush()

	ctx := context.Background()
	if err := Run(ctx, *network, *addr); err != nil {
		glog.Fatal(err)
	}
}