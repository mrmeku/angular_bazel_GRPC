/*
Command example-gateway-server is an example reverse-proxy implementation
whose HTTP handler is generated by grpc_gateway.
*/
package main

import (
	"angular_bazel_GRPC/grpc_gateway/gateway"
	"context"
	"flag"

	"github.com/golang/glog"
)

func main() {
	flag.Parse()
	defer glog.Flush()

	ctx := context.Background()
	opts := gateway.Options{
		Addr: ":8080",
		AdditionServer: gateway.Endpoint{
			Network: "tcp",
			Addr:    "localhost:9090",
		},
		MultiplicationServer: gateway.Endpoint{
			Network: "tcp",
			Addr:    "localhost:9091",
		},
		StaticData: Data,
	}
	if err := gateway.Run(ctx, opts); err != nil {
		glog.Fatal(err)
	}
}
