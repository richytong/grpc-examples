build:
	make build_pb
build_pb:
	protoc src/*.proto \
		--proto_path src \
		--go_out=plugins=grpc:.
clean:
	rm -rf proto/cache/*
runServer:
	go run src/server.go
runClient:
	go run src/client.go
