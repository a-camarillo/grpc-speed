package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/a-camarillo/grpc-speed/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
  var (
    grpcAddress = ":4000"
    ctx = context.Background()
  )

  mockPosition := [...]int32{0, 10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
  
  conn, err := grpc.Dial(grpcAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
  if err != nil {
    log.Fatal("Dial error: ", err)
  }

  defer conn.Close()

  grpcClient := proto.NewSpeedFetcherClient(conn)

  stream, err := grpcClient.FetchSpeed(ctx)
  if err != nil {
    log.Fatal("Opening stream ", err)
  }
  
  startTime := time.Now()
  time.Sleep(2*time.Second) 

  for _, i := range mockPosition {
    currentTime := time.Since(startTime).Seconds()
    err := stream.Send(&proto.TravelInfo{
      Position: i,
      Time: int32(currentTime),
    })
    fmt.Println("Sending")

    if err != nil {
      log.Fatalln(err)
    }
  }

  if err := stream.CloseSend(); err != nil {
      log.Fatal("Close Send ", err)
  }


  for {
    res, err := stream.Recv()
    if err == io.EOF {
      break
    }
    
      if err != nil {
        log.Fatalln("Closing", err)
      }

      fmt.Println("Speed", res.AverageSpeed)
    }
}

