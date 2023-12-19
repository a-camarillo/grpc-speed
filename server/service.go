package main

import (
	"fmt"
	"io"

	"github.com/a-camarillo/grpc-speed/proto"
)

//type SpeedFetcher interface {
//  FetchSpeed(context.Context, int32, int32) (int32, error)
//}

type speedFetcherService struct {
  proto.UnimplementedSpeedFetcherServer
}

func (s *speedFetcherService) FetchSpeed(stream proto.SpeedFetcher_FetchSpeedServer) error {
  
  for {
    req, err := stream.Recv()
    if err == io.EOF {
      return nil
    }

    if err != nil {
      return err
    }
    
    fmt.Printf("position: %d", req.Position)

    avg := (req.Position/req.Time)
    
    fmt.Println("Average: ", avg)

    if err := stream.Send(&proto.TravelResponse{
      AverageSpeed: avg,
    }); err != nil {
      return err
    }
  }
}

