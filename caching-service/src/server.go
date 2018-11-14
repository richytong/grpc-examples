package main

import (
	"log"
	"net"
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	pbCache "github.com/richytong/grpc-examples/caching-service/proto/cache"
)

type CacheService struct {
	Store map[string][]byte
}

func (cs *CacheService) Get(ctx context.Context, req *pbCache.GetReq) (*pbCache.GetResp, error) {
	defer log.Printf("CacheService Get call with req: %v\n", req)
	val, ok := cs.Store[req.Key]
	if !ok {
		return nil, status.Errorf(codes.NotFound, "Key not found: %s", req.Key)
	}
	return &pbCache.GetResp{Val: val}, nil
}

func (cs *CacheService) Set(ctx context.Context, req *pbCache.SetReq) (*pbCache.SetResp, error) {
	defer log.Printf("CacheService Set call with req: %v\n", req)
	cs.Store[req.Key] = req.Val
	return &pbCache.SetResp{}, nil
}

func runServer() error {
	srv := grpc.NewServer()
	pbCache.RegisterCacheServer(srv, &CacheService{Store: make(map[string][]byte)})
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
