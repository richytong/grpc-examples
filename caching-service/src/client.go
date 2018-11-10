package main

import (
	"log"
	"context"

	"google.golang.org/grpc"
	pbCache "github.com/richytong/grpc-examples/caching-service/gen/pb/cache"
)

func runClient() error {
	conn, err := grpc.Dial("localhost:5051", grpc.WithInsecure())
	if err != nil {
		return err
	}
	cache := pbCache.NewCacheClient(conn)
	setReq := pbCache.SetReq{Key: "gopher", Val: []byte("con")}
	_, err = cache.Set(context.Background(), &setReq)
	if err != nil {
		return err
	}
	getReq := pbCache.GetReq{Key: "gopher"}
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
