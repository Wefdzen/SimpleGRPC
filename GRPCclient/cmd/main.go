package main

import (
	"context"
	"fmt"
	pb "wefdzen/invoicer"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial("localhost:8190", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return
	}
	client := pb.NewAdderClient(conn)

	adderList, err := client.Add(context.Background(), &pb.HelloRequest{X: 25, Y: 13})
	if err != nil {
		return
	}
	fmt.Println(adderList)
}
