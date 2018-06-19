package gateway

import (
	"context"
	"net/http"
	"strings"

	"github.com/golang/glog"
	gwruntime "github.com/grpc-ecosystem/grpc-gateway/runtime"
)

// Endpoint describes a gRPC endpoint
type Endpoint struct {
	Network, Addr string
}

// Options is a set of options to be passed to Run
type Options struct {
	// Addr is the address to listen
	Addr string

	AdditionServer       Endpoint
	MultiplicationServer Endpoint

	StaticData map[string][]byte

	// Mux is a list of options to be passed to the grpc_gateway multiplexer
	Mux []gwruntime.ServeMuxOption
}

// Run starts a HTTP server and blocks while running if successful.
// The server will be shutdown when "ctx" is canceled.
func Run(ctx context.Context, opts Options) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	additionServerConn, additionServerErr := dial(ctx, opts.AdditionServer.Network, opts.AdditionServer.Addr)
	if additionServerErr != nil {
		return additionServerErr
	}

	multiplicationServerConn, multiplicationServerErr := dial(ctx, opts.MultiplicationServer.Network, opts.MultiplicationServer.Addr)
	if multiplicationServerErr != nil {
		return multiplicationServerErr
	}

	go func() {
		<-ctx.Done()
		if err := additionServerConn.Close(); err != nil {
			glog.Errorf("Failed to close a client connection to the Addition server: %v", err)
		}
		if err := multiplicationServerConn.Close(); err != nil {
			glog.Errorf("Failed to close a client connection to the Multiplication server: %v", err)
		}
	}()

	mux := http.NewServeMux()
	additionGw, additionGwErr := newGateway(ctx, additionServerConn, opts.Mux)
	if additionGwErr != nil {
		return additionGwErr
	}
	multiplicationGw, multiplicationGwErr := newGateway(ctx, multiplicationServerConn, opts.Mux)
	if multiplicationGwErr != nil {
		return multiplicationGwErr
	}

	mux.Handle("/v1/addition_service/", additionGw)
	mux.Handle("/v1/multiplication_service/", multiplicationGw)

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fileName := strings.TrimPrefix(r.URL.Path, "/")
		if fileName == "" {
			fileName = "index.html"
		}
		if val, ok := opts.StaticData[fileName]; ok {
			w.Write(val)
			return
		}
	})

	s := &http.Server{
		Addr:    opts.Addr,
		Handler: allowCORS(mux),
	}
	go func() {
		<-ctx.Done()
		glog.Infof("Shutting down the http server")
		if err := s.Shutdown(context.Background()); err != nil {
			glog.Errorf("Failed to shutdown http server: %v", err)
		}
	}()

	glog.Infof("Starting listening at %s", opts.Addr)
	if err := s.ListenAndServe(); err != http.ErrServerClosed {
		glog.Errorf("Failed to listen and serve: %v", err)
		return err
	}
	return nil
}
