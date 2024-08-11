package grpc

import (
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

func RegisterGrpc(port int, server func(*grpc.Server)) error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return err
	}

	s := grpc.NewServer()
	server(s)
	reflection.Register(s)

	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		return err
	}
	return nil
}
