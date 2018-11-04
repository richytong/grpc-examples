package main

import (
	"log"
	"context"

	"google.golang.org/grpc"
	"github.com/richytong/grpc-examples/caching-service/pb"
)

func runClient() error {
	conn, err := grpc.Dial("localhost:5051", grpc.WithInsecure())
	if err != nil {
		return err
	}
	cache := pb.NewCacheClient(conn)
	setReq := pb.SetReq{Key: "gopher", Val: []byte("con")}
	_, err = cache.Set(context.Background(), &setReq)
	if err != nil {
		return err
	}
	getReq := pb.GetReq{Key: "gopher"}
	resp, err := cache.Get(context.Background(), &getReq)
	if err != nil {
		return err
	}
	log.Printf("For key: %v, got cached value: %v, bytes: %v", getReq.Key, string(resp.Val), resp.Val)
	return nil
}

func main() {
	if err := runClient(); err != nil {
		log.Fatalf("failed to run client: %v", err)
	}
}
