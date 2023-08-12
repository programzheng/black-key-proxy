package server

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/programzheng/black-key-proxy/internal/service"
	pb "github.com/programzheng/black-key-proxy/pkg/proto"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedProxyServer
}

func (s *server) GetProxy(ctx context.Context, req *pb.GetProxyRequest) (*pb.GetProxyResponse, error) {
	redo := &service.RelayEventDataObject{
		Identifier: req.GetIdentifier(),
		Key:        req.GetKey(),
	}
	result := service.SendGetImageUrlByRelayEventDataObject(redo)
	return &pb.GetProxyResponse{
		StatusCode: result.StatusCode,
		Url:        result.Url,
	}, nil
}

func RunGrpcServer(port string) error {
	lis, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%s", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterProxyServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		return err
	}

	return nil
}
