package main

import (
	"context"
	"fmt"
	"log"

	"github.com/Faizal-Asep/grpc/pb"
	"google.golang.org/grpc"
)

type Client struct {
	conn    *grpc.ClientConn
	service pb.ServiceClient
}

func (c *Client) Close() {
	c.conn.Close()
}

func main() {
	connProfile, err := grpc.Dial(":50051", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
		return
	}

	serviceProfile := pb.NewServiceClient(connProfile)
	Client := &Client{connProfile, serviceProfile}

	res, err := Client.service.Ping(context.Background(), &pb.Request{})
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println(res)
}
