package main

import (
	"log"

	"github.com/a-camarillo/grpc-speed/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type grpcClient struct {
  Client proto.SpeedFetcherClient
  Conn *grpc.ClientConn
}

func NewGRPCClient(remoteAddr string) *grpcClient {
  conn, err := grpc.Dial(remoteAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
  if err != nil {
    log.Fatal(err)
  }
  client := proto.NewSpeedFetcherClient(conn)
  return &grpcClient{
    Client: client, 
    Conn: conn,
  }
}

