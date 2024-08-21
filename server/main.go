package main

import (
	"context"
	"fmt"
	"net"

	"github.com/HrithikMJ/go-grpc-server/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	pb.UnimplementedCalculatorServer
}

func (s *server) Add(ctx context.Context, req *pb.Req) (*pb.Resp, error) {
	return &pb.Resp{
		Result: req.A + req.B,
	}, nil

}
func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("failed to create listener:", err)
	}
	fmt.Println("Started listener at port 8080")
	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)
	pb.RegisterCalculatorServer(grpcServer, &server{})
	if err := grpcServer.Serve(listener); err != nil {
		fmt.Println("failed to serve: ", err)
	}
}
