build:
	make build_pb
build_pb:
	protoc src/api.proto \
		-I src \
		--go_out=plugins=grpc:pb
clean:
	rm -rf pb/*
