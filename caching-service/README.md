# caching-service
Barebones gRPC implementation of an in-memory cache

source: https://www.youtube.com/watch?v=7FZ6ZyzGex0

## Getting started
Install `protoc` - `brew install protobuf` (for mac), see http://google.github.io/proto-lens/installing-protoc.html

Install `protoc-gen-go`, a dependency of `protoc` - `go get -u github.com/golang/protobuf/protoc-gen-go`. This will install the `protoc-gen-go` binary in your `$GOBIN`. Make sure your `$GOBIN` directory is added to your `$PATH`.

## Generate gRPC Code
```bash
	protoc src/*.proto \
		--proto_path src \
		--go_out=plugins=grpc:.
```
  - `src/*.proto` - generate for all .proto files in `src`
  - `--proto_path src` - look in src for .proto files and trim `src` from final name
  - `--go_out=plugins=grpc:.` - use the grpc plugin and output into the `proto` directory at project root, configured in `cache.proto`

## Run Server
```bash
go run src/server.go
```

## Run Client
```bash
go run src/client.go
```
