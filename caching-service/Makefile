build:
	make build_pb
build_pb:
	protoc -I src/ src/api.proto --go_out=plugins=grpc:pb/
clean:
	rm pb/*
