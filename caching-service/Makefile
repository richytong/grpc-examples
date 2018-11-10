build:
	make build_pb
build_pb:
	protoc src/*.proto \
		--proto_path src \
		--go_out=plugins=grpc:gen
clean:
	rm -rf gen/*
