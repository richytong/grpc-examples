package main

import (
	"log"
	"net"
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"github.com/richytong/grpc-examples/caching-service/pb"
)

type CacheService struct {
	Store map[string][]byte
}

func (cs *CacheService) Get(ctx context.Context, req *pb.GetReq) (*pb.GetResp, error) {
	val, ok := cs.Store[req.Key]
	if !ok {
		return nil, status.Errorf(codes.NotFound, "Key not found: %s", req.Key)
	}
	return &pb.GetResp{Val: val}, nil
}

func (cs *CacheService) Set(ctx context.Context, req *pb.SetReq) (*pb.SetResp, error) {
	cs.Store[req.Key] = req.Val
	return &pb.SetResp{}, nil
}

func runServer() error {
	srv := grpc.NewServer()
	pb.RegisterCacheServer(srv, &CacheService{Store: make(map[string][]byte)})
	listener, err := net.Listen("tcp", "localhost:5051")
	if err != nil {
		return err
	}
	log.Printf("server started at %v\n", "localhost:5051")
	return srv.Serve(listener)
}

func main() {
	if err := runServer(); err != nil {
		log.Fatalf("failed to run the cache server: %v\n", err)
	}
}
