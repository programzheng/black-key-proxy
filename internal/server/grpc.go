package server

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

func RunGrpcServer(port string) error {
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%s", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	gs := grpc.NewServer()
	if err := gs.Serve(lis); err != nil {
		return err
	}

	return nil
}
