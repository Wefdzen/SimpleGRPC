package main

import (
	"context"
	"log"
	"net"
	"wefdzen/invoicer"

	"google.golang.org/grpc"
)

type myAdderServer struct {
	invoicer.UnimplementedAdderServer
}

func (s myAdderServer) Add(ctx context.Context, per *invoicer.HelloRequest) (*invoicer.HelloResponse, error) {

	return &invoicer.HelloResponse{
		Z: per.X + per.Y, // по приколу
	}, nil
}
func main() {
	lis, err := net.Listen("tcp", ":8190")
	if err != nil {
		log.Fatal("Cannot create listener: ", err)
	}
	serverAdder := grpc.NewServer()
	service := &myAdderServer{}
	invoicer.RegisterAdderServer(serverAdder, service)

	err = serverAdder.Serve(lis)
	if err != nil {
		log.Fatal("Can't Serve server ", err)
	}
}
