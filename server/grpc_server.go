package main

import (
	"fmt"
	"net"

	"github.com/a-camarillo/grpc-speed/proto"
	"google.golang.org/grpc"
)

func makeGRPCServerAndRun(listenAddress string, service *speedFetcherService) error {
  listener, err := net.Listen("tcp", listenAddress)
  fmt.Printf("now listening on port %s", listenAddress)
  if err != nil {
    return err
  }

  opts := []grpc.ServerOption{}
  server := grpc.NewServer(opts...)

  proto.RegisterSpeedFetcherServer(server, service)
  return server.Serve(listener)
}


