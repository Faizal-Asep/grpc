package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/Faizal-Asep/grpc/pb"
	"google.golang.org/grpc"
)

var port = flag.Int("port", 50051, "The server port")

type Server struct {
	pb.UnimplementedServiceServer
}

func (s *Server) Ping(context.Context, *pb.Request) (*pb.Response, error) {
	return &pb.Response{
		Code:    200,
		Message: "Ok",
	}, nil
	// return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}

func main() {

	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// create grpc server
	grpcServer := grpc.NewServer()
	pb.RegisterServiceServer(grpcServer, &Server{})
	if err = grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
}
